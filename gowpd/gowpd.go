package gowpd

/*
#cgo windows CFLAGS: -I "${SRCDIR}/../libgowpd/libgowpd"
#cgo windows LDFLAGS: -L "${SRCDIR}/../libgowpd/x64/Debug" -llibgowpd -lOle32
// -lPortableDeviceGuids -luuid

#include "libgowpd.h"
 */
import "C"
import (
	"unsafe"
	"log"
	"fmt"
)

const (
	S_OK HRESULT = C.S_OK & 0xffffffff// 0x00000000

	E_ABORT HRESULT = C.E_ABORT & 0xffffffff// 0x80004004
	E_ACCESSDENIED HRESULT = C.E_ACCESSDENIED & 0xffffffff// 0x80070005
	E_FAIL HRESULT = C.E_FAIL & 0xffffffff// 0x80004005
	E_HANDLE HRESULT = C.E_HANDLE & 0xffffffff// 0x80070006
	E_INVALIDARG HRESULT = C.E_INVALIDARG & 0xffffffff// 0x80070057
	E_NOINTERFACE HRESULT = C.E_NOINTERFACE & 0xffffffff// 0x80004002
	E_NOTIMPL HRESULT = C.E_NOTIMPL & 0xffffffff// 0x80004001
	E_OUTOFMEMORY HRESULT = C.E_OUTOFMEMORY & 0xffffffff// 0x8007000E
	E_POINTER HRESULT = C.E_POINTER & 0xffffffff// 0x80004003
	E_UNEXPECTED HRESULT = C.E_UNEXPECTED & 0xffffffff// 0x8000FFFF

	CO_E_NOTINITIALIZED HRESULT = C.CO_E_NOTINITIALIZED & 0xffffffff// 0x800401f0
)

const (
	WPD_CLIENT_NAME PropertyKey = iota
	WPD_CLIENT_MAJOR_VERSION
	WPD_CLIENT_MINOR_VERSION
	WPD_CLIENT_REVISION
	WPD_CLIENT_SECURITY_QUALITY_OF_SERVICE
	WPD_CLIENT_DESIRED_ACCESS
)

const (
	WPD_DEVICE_OBJECT_ID = "DEVICE"
)

const (
	GENERIC_READ = 0x80000000
	GENERIC_WRITE = 0x40000000
	GENERIC_EXECUTE = 0x20000000
	GENERIC_ALL = 0x10000000
)

// C.HRESULT
type HRESULT uint32
// C.DWORD
type DWORD uint32
// unsigned long
type ULONG uint32
// *WCHAR
type PnPDeviceID C.PnPDeviceID
// C.PROPERTYKEY
type PropertyKey int

type IPortableDevice C.IPortableDevice
type IPortableDeviceValues C.IPortableDeviceValues
type IPortableDeviceManager C.IPortableDeviceManager
type IPortableDeviceContent C.IPortableDeviceContent

type IEnumPortableDeviceObjectIDs C.IEnumPortableDeviceObjectIDs

func (hr HRESULT) Error() string {
	return fmt.Sprintf("error code: %s", hr.String())
}

func (hr HRESULT) String() string {
	switch hr {
	case S_OK:
		return "S_OK"
	case E_ABORT:
		return "E_ABORT"
	case E_ACCESSDENIED:
		return "E_ACCESSDENIED"
	case E_FAIL:
		return "E_FAIL"
	case E_HANDLE:
		return "E_HANDLE"
	case E_INVALIDARG:
		return "E_INVALIDARG"
	case E_NOINTERFACE:
		return "E_NOINTERFACE"
	case E_NOTIMPL:
		return "E_NOTIMPL"
	case E_OUTOFMEMORY:
		return "E_OUTOFMEMORY"
	case E_POINTER:
		return "E_POINTER"
	case E_UNEXPECTED:
		return "E_UNEXPECTED"
	case CO_E_NOTINITIALIZED:
		return "CO_E_NOTINITIALIZED"
	default:
		return fmt.Sprintf("%#x", uint32(hr))
	}
}

func (propertyKey PropertyKey) toCPropertyKey() *C.PROPERTYKEY {
	switch propertyKey {
	case WPD_CLIENT_NAME:
		return &C.WPD_CLIENT_NAME
	case WPD_CLIENT_MAJOR_VERSION:
		return &C.WPD_CLIENT_MAJOR_VERSION
	case WPD_CLIENT_MINOR_VERSION:
		return &C.WPD_CLIENT_MINOR_VERSION
	case WPD_CLIENT_REVISION:
		return &C.WPD_CLIENT_REVISION
	case WPD_CLIENT_SECURITY_QUALITY_OF_SERVICE:
		return &C.WPD_CLIENT_SECURITY_QUALITY_OF_SERVICE
	case WPD_CLIENT_DESIRED_ACCESS:
		return &C.WPD_CLIENT_DESIRED_ACCESS
	default:
		return nil// TODO error
	}
}

func Initialize() error {
	hr := C.CoInitializeEx(nil, C.COINIT_MULTITHREADED)
	if (hr < 0) {
		return HRESULT(hr)
	}

	return nil
}

func Uninitialize() {
	C.CoUninitialize()
}

func FreeDeviceID(pnpDeviceID PnPDeviceID) {
	C.CoTaskMemFree(C.LPVOID(pnpDeviceID))
}

func CreatePortableDevice() (*IPortableDevice, error) {
	var (
		pPortableDevice *C.IPortableDevice
	)

	log.Println("CreatePortableDevice(): Ready")
	hr := C.createPortableDevice(&pPortableDevice)

	if hr < 0 {
		return nil, HRESULT(hr)
	}

	log.Println("CreatePortableDevice(): Create portable device instance.")

	return (*IPortableDevice)(pPortableDevice), nil
}

func CreatePortableDeviceValues() (*IPortableDeviceValues, error) {
	var (
		pPortableDeviceValues *C.IPortableDeviceValues
	)

	log.Println("CreatePortableDeviceValues(): Ready")
	hr := C.createPortableDeviceValues(&pPortableDeviceValues)

	if hr < 0 {
		return nil, HRESULT(hr)
	}

	return (*IPortableDeviceValues)(pPortableDeviceValues), nil
}

func CreatePortableDeviceManager() (*IPortableDeviceManager, error) {
	var (
		pPortableDeviceManager *C.IPortableDeviceManager
	)

	log.Println("CreatePortableDeviceManager(): Ready")
	hr := C.createPortableDeviceManager(&pPortableDeviceManager)

	if hr < 0 {
		return nil, HRESULT(hr)
	}
	if pPortableDeviceManager == nil {
		return nil, E_POINTER
	}

	log.Println("CreatePortableDeviceManager(): Create portable device manager instance.")

	return (*IPortableDeviceManager)(pPortableDeviceManager), nil
}

func (pPortableDevice *IPortableDevice) Open(pnpDeviceID PnPDeviceID, pClientInfo *IPortableDeviceValues) error {
	log.Println("Open(): Ready")

	hr := C.portableDevice_Open((*C.IPortableDevice)(pPortableDevice), pnpDeviceID, (*C.IPortableDeviceValues)(pClientInfo))

	if hr < 0 {
		return HRESULT(hr)
	}

	return nil
}

func (pPortableDevice *IPortableDevice) Content() (*IPortableDeviceContent, error) {
	var (
		pPortableDeviceContent *C.IPortableDeviceContent
	)

	log.Println("Content(): Ready")

	hr := C.portableDevice_Content((*C.IPortableDevice)(pPortableDevice), &pPortableDeviceContent)

	if hr < 0 {
		return nil, HRESULT(hr)
	}

	return (*IPortableDeviceContent)(pPortableDeviceContent), nil
}

func (pPortableDevice *IPortableDevice) Release() error {
	hr := C.portableDevice_Release((*C.IPortableDevice)(pPortableDevice))

	if hr < 0 {
		return HRESULT(hr)
	}

	return nil
}

func (pPortableDeviceValues *IPortableDeviceValues) GetStringValue(key PropertyKey) (string, error) {
	var (
		pwstr C.PWSTR
		cPwstr C.DWORD
	)

	log.Println("GetStringValue(): Ready")

	hr := C.portableDeviceValues_GetStringValue((*C.IPortableDeviceValues)(pPortableDeviceValues), key.toCPropertyKey(), &pwstr, &cPwstr)
	defer C.CoTaskMemFree(C.LPVOID(pwstr))

	if hr < 0 {
		return "", HRESULT(hr)
	}

	raw := (*[1 << 30]C.WCHAR)(unsafe.Pointer(pwstr))[:cPwstr:cPwstr]
	str := make([]byte, DWORD(cPwstr))
	for i, wchar := range raw {
		str[i] = byte(wchar)
	}

	log.Printf("GetStringValue(): Result: %s\n", string(str))

	return string(str), nil
}

func (pPortableDeviceValues *IPortableDeviceValues) GetUnsignedIntegerValue(key PropertyKey) (uint32, error) {
	var (
		value C.ULONG
	)

	log.Println("GetUnsignedIntegerValue(): Ready")

	hr := C.portableDeviceValues_GetUnsignedIntegerValue((*C.IPortableDeviceValues)(pPortableDeviceValues), key.toCPropertyKey(), &value)

	if hr < 0 {
		return 0, HRESULT(hr)
	}

	log.Printf("GetUnsignedIntegerValue(): Result: %d\n", value)

	return uint32(value), nil
}

func (pPortableDeviceValues *IPortableDeviceValues) SetStringValue(key PropertyKey, value string) error {
	var (
		pwstr C.PWSTR
	)

	log.Println("SetStringValue(): Ready")

	pwstr = C.PWSTR(C.malloc(C.size_t(C.sizeof_WCHAR * (len(value) + 1))))
	defer C.free(unsafe.Pointer(pwstr))

	raw := (*[1 << 30]C.WCHAR)(unsafe.Pointer(pwstr))[:len(value) + 1:len(value) + 1]
	for i, char := range []byte(value) {
		raw[i] = C.WCHAR(char)
	}
	raw[len(value)] = 0

	hr := C.portableDeviceValues_SetStringValue((*C.IPortableDeviceValues)(pPortableDeviceValues), key.toCPropertyKey(), pwstr)

	if hr < 0 {
		return HRESULT(hr)
	}

	return nil
}

func (pPortableDeviceValues *IPortableDeviceValues) SetUnsignedIntegerValue(key PropertyKey, value uint32) error {
	log.Println("SetUnsignedIntegerValue(): Ready")

	hr := C.portableDeviceValues_SetUnsignedIntegerValue((*C.IPortableDeviceValues)(pPortableDeviceValues), key.toCPropertyKey(), C.ULONG(value))

	if hr < 0 {
		return HRESULT(hr)
	}

	return nil
}

func (pPortableDeviceManager *IPortableDeviceManager) GetDevices() ([]PnPDeviceID, error) {
	var (
		pPnPDeviceIDs *C.PnPDeviceID = nil
		cPnPDeviceIDs C.DWORD = 0
	)

	log.Println("GetDevices(): Ready")

	hr := C.portableDeviceManager_GetDevices((*C.IPortableDeviceManager)(pPortableDeviceManager), &pPnPDeviceIDs, &cPnPDeviceIDs)
	defer C.free(unsafe.Pointer(pPnPDeviceIDs))

	if hr < 0 {
		return nil, HRESULT(hr)
	}

	log.Printf("GetDevices(): %d devices has been found.\n", cPnPDeviceIDs)

	raw := (*[1 << 30]C.PnPDeviceID)(unsafe.Pointer(pPnPDeviceIDs))[:cPnPDeviceIDs:cPnPDeviceIDs]
	ids := make([]PnPDeviceID, uint32(cPnPDeviceIDs))
	for i, id := range raw {
		ids[i] = PnPDeviceID(id)
	}

	return ids, nil
}

func (pPortableDeviceManager *IPortableDeviceManager) GetDeviceFriendlyName(pnpDeviceID PnPDeviceID) (string, error) {
	var (
		pFriendlyName C.PWSTR
		cFriendlyName C.DWORD
	)

	hr := C.portableDeviceManager_GetDeviceFriendlyName((*C.struct_IPortableDeviceManager)(pPortableDeviceManager), pnpDeviceID, &pFriendlyName, &cFriendlyName)
	defer C.free(unsafe.Pointer(pFriendlyName))

	if (hr < 0) {
		return "", HRESULT(hr)
	}

	str := toGoString(pFriendlyName, cFriendlyName)

	log.Printf("GetDeviceFriendlyName(): %s\n", string(str))

	return string(str), nil
}

func (pPortableDeviceManager *IPortableDeviceManager) GetDeviceManufacturer(pnpDeviceID PnPDeviceID) (string, error) {
	var (
		pManufacturer C.PWSTR
		cManufacturer C.DWORD
	)

	hr := C.portableDeviceManager_GetDeviceManufacturer((*C.IPortableDeviceManager)(pPortableDeviceManager), pnpDeviceID, &pManufacturer, &cManufacturer)
	defer C.free(unsafe.Pointer(pManufacturer))

	if (hr < 0) {
		return "", HRESULT(hr)
	}

	raw := (*[1 << 30]C.WCHAR)(unsafe.Pointer(pManufacturer))[:cManufacturer:cManufacturer]
	str := make([]byte, uint32(cManufacturer))
	for i, wchar := range raw {
		str[i] = byte(wchar)
	}

	log.Printf("GetDeviceManufacturer(): %s\n", string(str))

	return string(str), nil
}

func (pPortableDeviceManager *IPortableDeviceManager) GetDeviceDescription(pnpDeviceID PnPDeviceID) (string, error) {
	var (
		pDescription C.PWSTR
		cDescription C.DWORD
	)

	hr := C.portableDeviceManager_GetDeviceDescription((*C.IPortableDeviceManager)(pPortableDeviceManager), pnpDeviceID, &pDescription, &cDescription)
	defer C.free(unsafe.Pointer(pDescription))

	if (hr < 0) {
		return "", HRESULT(hr)
	}

	raw := (*[1 << 30]C.WCHAR)(unsafe.Pointer(pDescription))[:cDescription:cDescription]
	str := make([]byte, uint32(cDescription))
	for i, wchar := range raw {
		str[i] = byte(wchar)
	}

	log.Printf("GetDeviceDescription(): %s\n", string(str))

	return string(str), nil
}

func (pPortableDeviceManager *IPortableDeviceManager) Release() {

}

// TODO not finished
func (pPortableDeviceContent *IPortableDeviceContent) EnumObjects(parentObjectID string) (*IEnumPortableDeviceObjectIDs, error) {
	var (
		flags C.DWORD = 0// ignored
		pwstrParentObjectID C.PWSTR// "DEVICE", empty string is valid but not should not be nullptr.
		pFilter *C.IPortableDeviceValues = nil// ignored
		pEnum *C.IEnumPortableDeviceObjectIDs
	)

	pwstrParentObjectID = allocatePWSTR(parentObjectID)
	defer C.free(unsafe.Pointer(pwstrParentObjectID))

	hr := C.portableDeviceContent_EnumObjects((*C.IPortableDeviceContent)(pPortableDeviceContent), flags, pwstrParentObjectID, pFilter, &pEnum)

	if hr < 0 {
		return nil, HRESULT(hr)
	}

	return (*IEnumPortableDeviceObjectIDs)(pEnum), nil
}

// TODO not implemented
func (pEnumObjectIDs *IEnumPortableDeviceObjectIDs) Next() {

}

// PWSTR to GoString.
func toGoString(str C.PWSTR, cStr C.DWORD) string {
	raw := (*[1 << 30]C.WCHAR)(unsafe.Pointer(str))[:cStr:cStr]
	goString := make([]byte, DWORD(cStr))
	for i, wchar := range raw {
		goString[i] = byte(wchar)
	}

	return string(goString)
}

// Allocate memory.
// Must be unallocated after being used.
func allocatePWSTR(value string) C.PWSTR {
	pwstr := C.malloc(C.size_t(C.sizeof_WCHAR * (len(value) + 1)))
	raw := (*[1 << 30]C.WCHAR)(pwstr)[:len(value) + 1:len(value) + 1]
	for i, b := range []byte(value) {
		raw[i] = C.WCHAR(b)
	}
	raw[len(value)] = 0

	return C.PWSTR(pwstr)
}

/*
This package provides WindowsPortableDevice API.


Example

IPortableDeviceManager

	gowpd.Initialize()

	pPortableDeviceManager, err := gowpd.CreatePortableDeviceManager()
	if err != nil {
		panic(err)
	}

	deviceIDs, err := pPortableDeviceManager.GetDevices()


Warnings not solved

	Warning: .drectve `/DEFAULTLIB:"uuid.lib" /DEFAULTLIB:"uuid.lib" ' unrecognized

	Warning: .drectve `/FAILIFMISMATCH:"_CRT_STDIO_ISO_WIDE_SPECIFIERS=0" /DEFAULTLIB:"uuid.lib" /DEFAULTLIB:"uuid.lib" ' unrecognized

	Warning: corrupt .drectve at end of def file

비주얼 스튜디오 에서 빌드시 추가되는 링크 옵션이라고 함. 말그대로 경고 이므로 무시해도 되는 듯 함.

*/
package gowpd

//go:generate go run gen/genconst.go

/*
#cgo windows amd64 CFLAGS: -I "${SRCDIR}/libgowpd/libgowpd"
#cgo windows amd64 LDFLAGS: -L "${SRCDIR}/libgowpd/x64/Debug" -llibgowpd -lOle32
//-lPortableDeviceGuids -luuid

#include "libgowpd.h"
 */
import "C"
import (
	"unsafe"
	"log"
	"fmt"
	"unicode/utf16"
	"unicode/utf8"
	"encoding/binary"
)

const (
	CLSCTX_INPROC_SERVER CLSCTX = 1 << iota
	CLSCTX_INPROC_HANDLER
	CLSCTX_LOCAL_SERVER
	CLSCTX_INPROC_SERVER16
	CLSCTX_REMOTE_SERVER
	CLSCTX_INPROC_HANDLER16
	CLSCTX_RESERVED1
	CLSCTX_RESERVED2
	CLSCTX_RESERVED3
	CLSCTX_RESERVED4
	CLSCTX_NO_CODE_DOWNLOAD
	CLSCTX_RESERVED5
	CLSCTX_NO_CUSTOM_MARSHAL
	CLSCTX_ENABLE_CODE_DOWNLOAD
	CLSCTX_NO_FAILURE_LOG
	CLSCTX_DISABLE_AAA
	CLSCTX_ENABLE_AAA
	CLSCTX_FROM_DEFAULT_CONTEXT
	CLSCTX_ACTIVATE_32_BIT_SERVER
	CLSCTX_ACTIVATE_64_BIT_SERVER
	CLSCTX_ENABLE_CLOAKING
	CLSCTX_APPCONTAINER
	CLSCTX_ACTIVATE_AAA_AS_IU
	CLSCTX_PS_DLL
)

const (
	S_OK HRESULT = C.S_OK & 0xffffffff// 0x00000000
	S_FALSE HRESULT = C.S_FALSE & 0xffffffff// 0x00000001

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

	//E_FILE_ALREADY_EXISTS HRESULT = 0x80070050
	//E_FILE_IS_BEING_USED_BY_ANOTHER_PROCESS HRESULT = 0x80070020
)

const (
	CLSID_PortableDevice CLSID = iota
	CLSID_PortableDeviceFTM
	CLSID_PortableDeviceManager
	CLSID_PortableDeviceKeyCollection
	CLSID_PortableDeviceValues
	CLSID_PortableDevicePropVariantCollection
)

const (
	IID_IPortableDevice IID = iota
	IID_IPortableDeviceManager
	IID_IPortableDeviceKeyCollection
	IID_IPortableDeviceContent
	IID_IPortableDeviceProperties
	IID_IPortableDeviceValues
	IID_IPortableDeviceDataStream
	IID_IPortableDevicePropVariantCollection
)

const (
	WPD_CLIENT_NAME PropertyKey = iota
	WPD_CLIENT_MAJOR_VERSION
	WPD_CLIENT_MINOR_VERSION
	WPD_CLIENT_REVISION
	WPD_CLIENT_SECURITY_QUALITY_OF_SERVICE
	WPD_CLIENT_DESIRED_ACCESS

	WPD_OBJECT_PARENT_ID
	WPD_OBJECT_NAME
	WPD_OBJECT_PERSISTENT_UNIQUE_ID
	WPD_OBJECT_FORMAT
	WPD_OBJECT_CONTENT_TYPE
	WPD_OBJECT_SIZE
	WPD_OBJECT_ORIGINAL_FILE_NAME

	WPD_PROPERTY_ATTRIBUTE_FORM
	WPD_PROPERTY_ATTRIBUTE_CAN_READ
	WPD_PROPERTY_ATTRIBUTE_CAN_WRITE
	WPD_PROPERTY_ATTRIBUTE_CAN_DELETE
	WPD_PROPERTY_ATTRIBUTE_DEFAULT_VALUE
	WPD_PROPERTY_ATTRIBUTE_FAST_PROPERTY
	WPD_PROPERTY_ATTRIBUTE_RANGE_MIN
	WPD_PROPERTY_ATTRIBUTE_RANGE_MAX
	WPD_PROPERTY_ATTRIBUTE_RANGE_STEP
	WPD_PROPERTY_ATTRIBUTE_ENUMERATION_ELEMENTS
	WPD_PROPERTY_ATTRIBUTE_REGULAR_EXPRESSION
	WPD_PROPERTY_ATTRIBUTE_MAX_SIZE

	WPD_RESOURCE_DEFAULT
)

const (
	WPD_CONTENT_TYPE_FUNCTIONAL_OBJECT GUID = iota
	WPD_CONTENT_TYPE_FOLDER
	WPD_CONTENT_TYPE_IMAGE
	WPD_CONTENT_TYPE_DOCUMENT
	WPD_CONTENT_TYPE_CONTACT
	WPD_CONTENT_TYPE_CONTACT_GROUP
	WPD_CONTENT_TYPE_AUDIO
	WPD_CONTENT_TYPE_VIDEO
	WPD_CONTENT_TYPE_TELEVISION
	WPD_CONTENT_TYPE_PLAYLIST
	WPD_CONTENT_TYPE_MIXED_CONTENT_ALBUM
	WPD_CONTENT_TYPE_AUDIO_ALBUM
	WPD_CONTENT_TYPE_IMAGE_ALBUM
	WPD_CONTENT_TYPE_VIDEO_ALBUM
	WPD_CONTENT_TYPE_MEMO
	WPD_CONTENT_TYPE_EMAIL
	WPD_CONTENT_TYPE_APPOINTMENT
	WPD_CONTENT_TYPE_TASK
	WPD_CONTENT_TYPE_PROGRAM
	WPD_CONTENT_TYPE_GENERIC_FILE
	WPD_CONTENT_TYPE_CALENDAR
	WPD_CONTENT_TYPE_GENERIC_MESSAGE
	WPD_CONTENT_TYPE_NETWORK_ASSOCIATION
	WPD_CONTENT_TYPE_CERTIFICATE
	WPD_CONTENT_TYPE_WIRELESS_PROFILE
	WPD_CONTENT_TYPE_MEDIA_CAST
	WPD_CONTENT_TYPE_SECTION
	WPD_CONTENT_TYPE_UNSPECIFIED
	WPD_CONTENT_TYPE_ALL

	WPD_OBJECT_FORMAT_EXIF
	WPD_OBJECT_FORMAT_WMA
	WPD_OBJECT_FORMAT_VCARD2
)

const (
	WPD_DEVICE_OBJECT_ID string = "DEVICE"
)

const (
	GENERIC_READ = C.GENERIC_READ & 0xffffffff// 0x80000000
	GENERIC_WRITE = C.GENERIC_WRITE & 0xffffffff// 0x40000000
	GENERIC_EXECUTE = C.GENERIC_EXECUTE & 0xffffffff// 0x20000000
	GENERIC_ALL = C.GENERIC_ALL & 0xffffffff// 0x10000000
)

const (
	STATFLAG_DEFAULT = C.STATFLAG_DEFAULT
	STATFLAG_NONAME = C.STATFLAG_NONAME
	STATFLAG_NOOPEN = C.STATFLAG_NOOPEN
)

const (
	STGM_READ = C.STGM_READ
	STGM_CREATE = C.STGM_CREATE
	STGM_WRITE = C.STGM_WRITE
)

const (
	VT_EMPTY VARTYPE = 0
	VT_NULL VARTYPE = 1
	VT_I2 VARTYPE = 2
	VT_I4 VARTYPE = 3
	VT_R4 VARTYPE = 4
	VT_R8 VARTYPE = 5
	VT_CY VARTYPE = 6
	VT_DATE VARTYPE = 7
	VT_BSTR VARTYPE = 8
	VT_DISPATCH VARTYPE = 9
	VT_ERROR VARTYPE = 10
	VT_BOOL VARTYPE = 11
	VT_VARIANT VARTYPE = 12
	VT_UNKNOWN VARTYPE = 13
	VT_DECIMAL VARTYPE = 14
	VT_I1 VARTYPE = 16
	VT_UI1 VARTYPE = 17
	VT_UI2 VARTYPE = 18
	VT_UI4 VARTYPE = 19
	VT_I8 VARTYPE = 20
	VT_UI8 VARTYPE = 21
	VT_INT VARTYPE = 22
	VT_UINT VARTYPE = 23
	VT_VOID VARTYPE = 24
	VT_HRESULT VARTYPE = 25
	VT_PTR VARTYPE = 26
	VT_SAFEARRAY VARTYPE = 27
	VT_CARRAY VARTYPE = 28
	VT_USERDEFINED VARTYPE = 29
	VT_LPSTR VARTYPE = 30
	VT_LPWSTR VARTYPE = 31
	VT_RECORD VARTYPE = 36
	VT_INT_PTR VARTYPE = 37
	VT_UINT_PTR VARTYPE = 38
	VT_FILETIME VARTYPE = 64
	VT_BLOB VARTYPE = 65
	VT_STREAM VARTYPE = 66
	VT_STORAGE VARTYPE = 67
	VT_STREAMED_OBJECT VARTYPE = 68
	VT_STORED_OBJECT VARTYPE = 69
	VT_BLOB_OBJECT VARTYPE = 70
	VT_CF VARTYPE = 71
	VT_CLSID VARTYPE = 72
	VT_VERSIONED_STREAM VARTYPE = 73
	VT_BSTR_BLOB VARTYPE = 0xfff
	VT_VECTOR VARTYPE = 0x1000
	VT_ARRAY VARTYPE = 0x2000
	VT_BYREF VARTYPE = 0x4000
	VT_RESERVED VARTYPE = 0x8000
	VT_ILLEGAL VARTYPE = 0xffff
	VT_ILLEGALMASKED VARTYPE = 0xfff
	VT_TYPEMASK VARTYPE = 0xfff
)

const (
	PORTABLE_DEVICE_DELETE_NO_RECURSION uint32 = iota
	PORTABLE_DEVICE_DELETE_WITH_RECURSION
)

// C.WCHAR
// 16bit-encoded
type WCHAR uint16;
// C.HRESULT
type HRESULT uint32
// C.DWORD
type DWORD uint32
// unsigned long
type ULONG uint32
// *WCHAR
type PnPDeviceID C.PnPDeviceID
// C.CLSCTX
type CLSCTX int
// C.CLSID
type CLSID int
// C.IID
type IID int
// C.PROPERTYKEY
type PropertyKey int
// C.GUID
type GUID int
// C.VARENUM
type VARTYPE uint16

type IPortableDevice C.IPortableDevice
type IPortableDeviceValues C.IPortableDeviceValues
type IPortableDeviceManager C.IPortableDeviceManager
type IPortableDeviceContent C.IPortableDeviceContent
type IPortableDeviceKeyCollection C.IPortableDeviceKeyCollection
type IPortableDeviceProperties C.IPortableDeviceProperties
type IPortableDeviceDataStream C.IPortableDeviceDataStream
type IPortableDeviceCapabilities C.IPortableDeviceCapabilities
type IPortableDevicePropVariantCollection C.IPortableDevicePropVariantCollection
type IPortableDeviceEventCallback C.IPortableDeviceEventCallback
type IPortableDeviceResources C.IPortableDeviceResources
type IStream C.IStream
type ISequentialStream C.ISequentialStream
type IPropertyStore C.IPropertyStore
type IUnknown C.IUnknown

type IEnumPortableDeviceObjectIDs C.IEnumPortableDeviceObjectIDs

type StatStg struct {
	pwcsName string
	_type uint32
	cbSize uint64
	mtime uint64
	ctime uint64
	atime uint64
	grfMode uint32
	grfLocksSupported uint32
	clsid CLSID
	grfStateBits uint32
	reserved uint32
}

func (hr HRESULT) Error() string {
	return fmt.Sprintf("error code: %s", hr.String())
}

func Initialize() error {
	log.Println("Initialize():")

	hr := C.CoInitializeEx(nil, C.COINIT_MULTITHREADED)
	if (hr < 0) {
		return HRESULT(hr)
	}

	return nil
}

func Uninitialize() {
	log.Println("Uninitialize():")

	C.CoUninitialize()
}

func FreeDeviceID(pnpDeviceID PnPDeviceID) {
	C.CoTaskMemFree(C.LPVOID(pnpDeviceID))
}

func CoCreateInstance(clsid CLSID, iid IID) (unsafe.Pointer, error) {
	var (
		pInstance C.LPVOID
	)

	hr := C.CoCreateInstance(clsid.toCCLSID(), nil, C.DWORD(CLSCTX_INPROC_SERVER), iid.toCIID(), &pInstance)

	if hr < 0 {
		return nil, HRESULT(hr)
	}

	return unsafe.Pointer(pInstance), nil
}

func CreatePortableDevice() (*IPortableDevice, error) {
	log.Println("CreatePortableDevice(): Ready")

	ptr, err := CoCreateInstance(CLSID_PortableDeviceFTM, IID_IPortableDevice)

	return (*IPortableDevice)(ptr), err
}

func CreatePortableDeviceValues() (*IPortableDeviceValues, error) {
	log.Println("CreatePortableDeviceValues(): Ready")

	ptr, err := CoCreateInstance(CLSID_PortableDeviceValues, IID_IPortableDeviceValues)

	return (*IPortableDeviceValues)(ptr), err
}

func CreatePortableDeviceManager() (*IPortableDeviceManager, error) {
	log.Println("CreatePortableDeviceManager(): Create portable device manager instance.")

	ptr, err := CoCreateInstance(CLSID_PortableDeviceManager, IID_IPortableDeviceManager)

	return (*IPortableDeviceManager)(ptr), err
}

func CreatePortableDeviceKeyCollection() (*IPortableDeviceKeyCollection, error) {
	log.Println("CreatePortableDeviceKeyCollection(): Ready")

	ptr, err := CoCreateInstance(CLSID_PortableDeviceKeyCollection, IID_IPortableDeviceKeyCollection)

	return (*IPortableDeviceKeyCollection)(ptr), err
}

func CreatePortableDevicePropVariantCollection() (*IPortableDevicePropVariantCollection, error) {
	log.Println("CreatePortableDevicePropVariantCollection(): Ready")

	ptr, err := CoCreateInstance(CLSID_PortableDevicePropVariantCollection, IID_IPortableDevicePropVariantCollection)

	return (*IPortableDevicePropVariantCollection)(ptr), err
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

func (pPortableDevice *IPortableDevice) Open(pnpDeviceID PnPDeviceID, pClientInfo *IPortableDeviceValues) error {
	log.Println("Open(): Ready")

	hr := C.portableDevice_Open((*C.IPortableDevice)(pPortableDevice), pnpDeviceID, (*C.IPortableDeviceValues)(pClientInfo))

	if hr < 0 {
		return HRESULT(hr)
	}

	return nil
}

func (pPortableDevice *IPortableDevice) Release() error {
	hr := C.portableDevice_Release((*C.IPortableDevice)(pPortableDevice))

	if hr < 0 {
		return HRESULT(hr)
	}

	return nil
}

func (pPortableDeviceValues *IPortableDeviceValues) GetBoolValue(key PropertyKey) (bool, error) {
	var (
		value C.BOOL// non-zero is TRUE, zero is FALSE. Windows Type.
	)

	hr := C.portableDeviceValues_GetBoolValue((*C.IPortableDeviceValues)(pPortableDeviceValues), key.toCPropertyKey(), &value)

	if hr < 0 {
		return false, HRESULT(hr)
	}

	return value != 0, nil
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

func (pPortableDeviceValues *IPortableDeviceValues) SetGuidValue(key PropertyKey, value GUID) error {
	hr := C.portableDeviceValues_SetGuidValue((*C.IPortableDeviceValues)(pPortableDeviceValues), key.toCPropertyKey(), value.toCGUID())

	if hr < 0 {
		return HRESULT(hr)
	}

	return nil
}

func (pPortableDeviceValues *IPortableDeviceValues) SetStringValue(key PropertyKey, value string) error {
	var (
		pwstr C.PWSTR
	)

	log.Println("SetStringValue(): Ready")

	pwstr = C.PWSTR(C.malloc(C.size_t(C.sizeof_WCHAR * (len(value) + 1))))
	if pwstr == nil {
		return E_POINTER
	}
	defer C.free(unsafe.Pointer(pwstr))

	log.Println("SetStringValue(): memory allocated")

	raw := (*[1 << 30]C.WCHAR)(unsafe.Pointer(pwstr))[:len(value) + 1:len(value) + 1]
	for i, char := range []byte(value) {
		raw[i] = C.WCHAR(char)
	}
	raw[len(value)] = 0

	hr := C.portableDeviceValues_SetStringValue((*C.IPortableDeviceValues)(pPortableDeviceValues), key.toCPropertyKey(), pwstr)

	if hr < 0 {
		return HRESULT(hr)
	}

	log.Printf("SetStringValue(): {hresult: %#x}\n", hr)

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

func (pPortableDeviceValeus *IPortableDeviceValues) SetUnsignedLargeIntegerValue(key PropertyKey, value uint64) error {
	hr := C.portableDeviceValues_SetUnsignedLargeIntegerValue((*C.IPortableDeviceValues)(pPortableDeviceValeus), key.toCPropertyKey(), C.ULONGLONG(value))

	if hr < 0 {
		return HRESULT(hr)
	}

	return nil
}

func (pPortableDeviceValues *IPortableDeviceValues) QueryInterface(iid IID) (unsafe.Pointer, error) {
	pUnknown := (*IUnknown)(unsafe.Pointer(pPortableDeviceValues))

	return pUnknown.QueryInterface(iid)
}

func (pPortableDeviceValues *IPortableDeviceValues) Release() error {
	hr := C.portableDeviceValues_Release((*C.IPortableDeviceValues)(pPortableDeviceValues))

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
		cFriendlyName C.DWORD = 0
	)

	log.Println("GetDeviceFriendlyName(): Ready")

	hr := C.portableDeviceManager_GetDeviceFriendlyName((*C.struct_IPortableDeviceManager)(pPortableDeviceManager), pnpDeviceID, &pFriendlyName, &cFriendlyName)
	defer C.free(unsafe.Pointer(pFriendlyName))

	if (hr < 0) {
		return "", HRESULT(hr)
	}

	str := toGoString(pFriendlyName, uint32(cFriendlyName))

	log.Printf("GetDeviceFriendlyName(): %s\n", str)

	return str, nil
}

func (pPortableDeviceManager *IPortableDeviceManager) GetDeviceManufacturer(pnpDeviceID PnPDeviceID) (string, error) {
	var (
		pManufacturer C.PWSTR
		cManufacturer C.DWORD = 0
	)

	hr := C.portableDeviceManager_GetDeviceManufacturer((*C.IPortableDeviceManager)(pPortableDeviceManager), pnpDeviceID, &pManufacturer, &cManufacturer)
	defer C.free(unsafe.Pointer(pManufacturer))

	if (hr < 0) {
		return "", HRESULT(hr)
	}

	str := toGoString(pManufacturer, uint32(cManufacturer))

	log.Printf("GetDeviceManufacturer(): %s\n", str)

	return str, nil
}

func (pPortableDeviceManager *IPortableDeviceManager) GetDeviceDescription(pnpDeviceID PnPDeviceID) (string, error) {
	var (
		pDescription C.PWSTR
		cDescription C.DWORD = 0
	)

	hr := C.portableDeviceManager_GetDeviceDescription((*C.IPortableDeviceManager)(pPortableDeviceManager), pnpDeviceID, &pDescription, &cDescription)
	defer C.free(unsafe.Pointer(pDescription))

	if (hr < 0) {
		return "", HRESULT(hr)
	}

	str := toGoString(pDescription, uint32(cDescription))

	log.Printf("GetDeviceDescription(): %s\n", str)

	return str, nil
}

func (pPortableDeviceManager *IPortableDeviceManager) Release() {

}

// TODO not finished
func (pPortableDeviceContent *IPortableDeviceContent) CreateObjectWithPropertiesAndData(pValues *IPortableDeviceValues) (*IStream, uint32, error) {
	var (
		pData *C.IStream
		optimalWriteBufferSize C.DWORD = 0// TRUE for ignoring
		pCookie C.PWSTR = nil// Optional.
	)

	hr := C.portableDeviceContent_CreateObjectWithPropertiesAndData((*C.IPortableDeviceContent)(pPortableDeviceContent), (*C.IPortableDeviceValues)(pValues), &pData, &optimalWriteBufferSize, &pCookie)

	if hr < 0 {
		return nil, 0, HRESULT(hr)
	}

	log.Printf("CreateObjectWithPropertiesAndData(): {optimalWriteBufferSize: %d}\n", optimalWriteBufferSize)

	return (*IStream)(pData), uint32(optimalWriteBufferSize), nil
}

// TODO not finished
// parentObjectID: start from it.
func (pPortableDeviceContent *IPortableDeviceContent) EnumObjects(parentObjectID string) (*IEnumPortableDeviceObjectIDs, error) {
	var (
		flags C.DWORD = 0// ignored
		pwstrParentObjectID C.PWSTR// "DEVICE", empty string is valid but not should not be nullptr.
		pFilter *C.IPortableDeviceValues = nil// ignored
		pEnum *C.IEnumPortableDeviceObjectIDs
	)

	pwstrParentObjectID, _ = allocatePWSTR(parentObjectID)
	defer C.free(unsafe.Pointer(pwstrParentObjectID))

	log.Println("EnumObjects(): Ready")

	hr := C.portableDeviceContent_EnumObjects((*C.IPortableDeviceContent)(pPortableDeviceContent), flags, pwstrParentObjectID, pFilter, &pEnum)

	if hr < 0 {
		return nil, HRESULT(hr)
	}

	return (*IEnumPortableDeviceObjectIDs)(pEnum), nil
}

func (pPortableDeviceContent *IPortableDeviceContent) Properties() (*IPortableDeviceProperties, error) {
	var (
		pPortableDeviceProperties *C.IPortableDeviceProperties
	)

	log.Println("Properties(): Ready")

	hr := C.portableDeviceContent_Properties((*C.IPortableDeviceContent)(pPortableDeviceContent), &pPortableDeviceProperties)

	if hr < 0 {
		return nil, HRESULT(hr)
	}

	return (*IPortableDeviceProperties)(pPortableDeviceProperties), nil
}

func (pPortableDeviceContent *IPortableDeviceContent) Transfer() (*IPortableDeviceResources, error) {
	var (
		pPortableDeviceResources *C.IPortableDeviceResources
	)

	hr := C.portableDeviceContent_Transfer((*C.IPortableDeviceContent)(pPortableDeviceContent), &pPortableDeviceResources)

	if hr < 0 {
		return nil, HRESULT(hr)
	}

	return (*IPortableDeviceResources)(pPortableDeviceResources), nil
}

func (pPortableDeviceContent *IPortableDeviceContent) Delete(options uint32, objectIDs *IPortableDevicePropVariantCollection) (*IPortableDevicePropVariantCollection, error) {
	var (
		results *C.IPortableDevicePropVariantCollection
	)

	hr := C.portableDeviceContent_Delete((*C.IPortableDeviceContent)(pPortableDeviceContent), C.DWORD(options), (*C.IPortableDevicePropVariantCollection)(objectIDs), &results)

	if hr < 0 {
		return nil, HRESULT(hr)
	}
	if HRESULT(hr) == S_FALSE {
		return (*IPortableDevicePropVariantCollection)(results), HRESULT(hr)
	}

	log.Printf("Delete(): Result %#x\n", uint32(hr))

	return (*IPortableDevicePropVariantCollection)(results), nil
}

func (pPortableDeviceKeyCollection *IPortableDeviceKeyCollection) Add(key PropertyKey) error {
	log.Println("Add(): Ready")

	hr := C.portableDeviceKeyCollection_Add((*C.IPortableDeviceKeyCollection)(pPortableDeviceKeyCollection), key.toCPropertyKey())

	if hr < 0 {
		return HRESULT(hr)
	}

	return nil
}

func (pPortableDeviceProperties *IPortableDeviceProperties) GetValues(objectID string, keys *IPortableDeviceKeyCollection) (*IPortableDeviceValues, error) {
	var (
		pPortableDeviceValues *C.IPortableDeviceValues
	)

	pObjectID, err := allocatePWSTR(objectID)
	if err != nil {
		panic(err)
	}
	defer C.free(unsafe.Pointer(pObjectID))

	hr := C.portableDeviceProperties_GetValues((*C.IPortableDeviceProperties)(pPortableDeviceProperties), pObjectID, (*C.IPortableDeviceKeyCollection)(keys), &pPortableDeviceValues)

	if hr < 0 {
		return nil, HRESULT(hr)
	}

	return (*IPortableDeviceValues)(pPortableDeviceValues), nil
}

func (pPortableDeviceProperties *IPortableDeviceProperties) GetPropertyAttributes(objectID string, key PropertyKey) (*IPortableDeviceValues, error) {
	var (
		pPortableDeviceValues *C.IPortableDeviceValues
	)

	pObjectID, err := allocatePWSTR(objectID)
	if err != nil {
		panic(err)
	}
	defer C.free(unsafe.Pointer(pObjectID))

	hr := C.portableDeviceProperties_GetPropertyAttributes((*C.IPortableDeviceProperties)(pPortableDeviceProperties), pObjectID, key.toCPropertyKey(), &pPortableDeviceValues)

	if hr < 0 {
		return nil, HRESULT(hr)
	}

	return (*IPortableDeviceValues)(pPortableDeviceValues), nil
}

// TODO not finished
func (pPortableDeviceProperties *IPortableDeviceProperties) SetValues(objectID string, pValues *IPortableDeviceValues) error {
	var (
		pResults *C.IPortableDeviceValues
	)

	pObjectID, err := allocatePWSTR(objectID)
	if err != nil {
		panic(err)
	}
	defer C.free(unsafe.Pointer(pObjectID))

	hr := C.portableDeviceProperties_SetValues((*C.IPortableDeviceProperties)(pPortableDeviceProperties), pObjectID, (*C.IPortableDeviceValues)(pValues), &pResults)
	if hr < 0 {
		return HRESULT(hr)
	}

	// TODO do something with pResults


	err = (*IPortableDeviceValues)(pResults).Release()
	if err != nil {
		panic(err)
	}

	return nil
}

func (pPortableDeviceDataStream *IPortableDeviceDataStream) Commit(dataFlags uint32) error {
	//hr := C.portableDeviceDataStream_Commit((*C.IPortableDeviceDataStream)(pPortableDeviceDataStream), C.DWORD(dataFlags))
	//
	//if hr < 0 {
	//	return HRESULT(hr)
	//}
	//
	//return nil

	pStream := (*IStream)(unsafe.Pointer(pPortableDeviceDataStream))

	return pStream.Commit(dataFlags)
}

func (pPortableDeviceDataStream *IPortableDeviceDataStream) GetObjectID() (string, error) {
	var (
		pObjectID C.PWSTR
	)

	hr := C.portableDeviceDataStream_GetObjectID((*C.IPortableDeviceDataStream)(pPortableDeviceDataStream), &pObjectID)
	defer C.CoTaskMemFree(C.LPVOID(pObjectID))

	if hr < 0 {
		return "", HRESULT(hr)
	}

	objectID := toGoString(pObjectID, wcslen(pObjectID))

	return objectID, nil
}

// cObjects: Number of objects to request on each NEXT
//
func (pEnumObjectIDs *IEnumPortableDeviceObjectIDs) Next(cObjects uint32) ([]string, error) {
	var (
		pObjIDs *C.PWSTR// Array of PWSTR. Not a PWSTR. Must have a size of cObjects. ObjectIDs will be here.
		cObjIDs *C.DWORD// Array of size of PWSTR.
		cPetched C.ULONG// amounts of ObjectID placed in pObjIDs.
	)

	pObjIDs = (*C.PWSTR)(C.malloc(C.size_t(C.sizeof_PWSTR * cObjects)))
	cObjIDs = (*C.DWORD)(C.malloc(C.size_t(C.sizeof_DWORD * cObjects)))
	defer C.free(unsafe.Pointer(pObjIDs))
	defer C.free(unsafe.Pointer(cObjIDs))

	log.Println("Next(): Ready")

	hr := C.enumPortableDeviceObjectIDs_Next((*C.IEnumPortableDeviceObjectIDs)(pEnumObjectIDs), C.ULONG(cObjects), pObjIDs, cObjIDs, &cPetched)

	log.Printf("Next(): {result: %s, cPetched: %d}\n", HRESULT(hr), cPetched)

	if hr < 0 {
		return nil, HRESULT(hr)
	}

	rawPWSTR := (*[1 << 30]C.PWSTR)(unsafe.Pointer(pObjIDs))[:cPetched:cPetched]
	//rawDWORD := (*[1 << 30]C.DWORD)(unsafe.Pointer(cObjIDs))[:cPetched:cPetched]// TODO delete

	objects := make([]string, ULONG(cPetched))
	for i, pwstr := range rawPWSTR {
		str := toGoString(pwstr, wcslen(pwstr))
		objects[i] = str

		log.Printf("Next(): {object: %s}\n", str)

		C.CoTaskMemFree(C.LPVOID(pwstr))
	}

	return objects, nil
}

func (pPortableDevicePropVariantCollection *IPortableDevicePropVariantCollection) Add(value *PropVariant) error {
	hr := C.portableDevicePropVariantCollection_Add((*C.IPortableDevicePropVariantCollection)(pPortableDevicePropVariantCollection), (*C.PROPVARIANT)(value))

	if hr < 0 {
		return HRESULT(hr)
	}

	return nil
}

func (pPortableDevicePropVariantCollection *IPortableDevicePropVariantCollection) GetAt(index uint32) (*PropVariant, error) {
	var (
		value *C.PROPVARIANT = new(C.PROPVARIANT)
	)

	if pPortableDevicePropVariantCollection == nil {
		return nil, E_POINTER
	}

	hr := C.portableDevicePropVariantCollection_GetAt((*C.IPortableDevicePropVariantCollection)(pPortableDevicePropVariantCollection), C.DWORD(index), value)

	if hr < 0 {
		return nil, HRESULT(hr)
	}

	return (*PropVariant)(value), nil
}

func (pPortableDevicePropVariantCollection *IPortableDevicePropVariantCollection) GetCount() (uint32, error) {
	var (
		count C.DWORD = 0
	)

	if pPortableDevicePropVariantCollection == nil {
		return 0, E_POINTER
	}

	hr := C.portableDevicePropVariantCollection_GetCount((*C.IPortableDevicePropVariantCollection)(pPortableDevicePropVariantCollection), &count)

	if hr < 0 {
		return 0, HRESULT(hr)
	}

	return uint32(count), nil
}

func (pPortableDeviceResources *IPortableDeviceResources) GetStream(objectID string, key PropertyKey, mode uint32) (*IStream, uint32, error) {
	var (
		pStream *C.IStream
		optimalBufferSize C.DWORD
	)

	pObjectID, err := allocatePWSTR(objectID)
	if err != nil {
		return nil, 0, err
	}
	defer C.free(unsafe.Pointer(pObjectID))

	hr := C.portableDeviceResources_GetStream((*C.IPortableDeviceResources)(pPortableDeviceResources), pObjectID, key.toCPropertyKey(), C.DWORD(mode), &optimalBufferSize, &pStream)

	if hr < 0 {
		return nil, 0, HRESULT(hr)
	}

	return (*IStream)(pStream), uint32(optimalBufferSize), nil
}

func (pStream *IStream) Commit(dataFlag uint32) error {
	hr := C.stream_Commit((*C.IStream)(pStream), C.DWORD(dataFlag))

	if hr < 0 {
		return HRESULT(hr)
	}

	return nil
}

func (pStream *IStream) Stat(statFlags uint32) (*StatStg, error) {
	var (
		statstg C.STATSTG
	)

	hr := C.stream_Stat((*C.IStream)(pStream), &statstg, C.DWORD(statFlags))

	if hr < 0 {
		return nil, HRESULT(hr)
	}

	result := new(StatStg)

	if statFlags & STATFLAG_NONAME == 0 {
		pwcsName := C.PWSTR(unsafe.Pointer(statstg.pwcsName))// nil if noname flags is set. if pwcsName is not nil, must call CoTaskMemFree method.
		result.pwcsName = toGoString(pwcsName, wcslen(pwcsName))
	}
	result._type = uint32(statstg._type)
	cbSize := [8]byte(statstg.cbSize)
	result.cbSize = binary.BigEndian.Uint64(cbSize[:])// C.ULARGE_INTEGER 64bit(8byte, DWORD times 2) union
	result.mtime = uint64(statstg.mtime.dwLowDateTime << 32 | statstg.mtime.dwHighDateTime)// C.FILETIME, struct contains two DWORD fields.
	result.ctime = uint64(statstg.ctime.dwLowDateTime << 32 | statstg.ctime.dwHighDateTime)// C.FILETIME
	result.atime = uint64(statstg.atime.dwLowDateTime << 32 | statstg.atime.dwHighDateTime)// C.FILETIME
	result.grfMode = uint32(statstg.grfMode)
	result.grfLocksSupported = uint32(statstg.grfLocksSupported)
	result.clsid = 0//C.CLSID(statstg.clsid)
	result.grfStateBits = uint32(statstg.grfStateBits)
	result.reserved = uint32(statstg.reserved)

	return result, nil
}

func (pStream *IStream) QueryInterface(iid IID) (unsafe.Pointer, error) {
	pUnknown := (*IUnknown)(unsafe.Pointer(pStream))

	return pUnknown.QueryInterface(iid)
}

func (pStream *IStream) Read(buffer []byte) (uint32, error) {
	pSequentialStream := (*ISequentialStream)(unsafe.Pointer(pStream))

	return pSequentialStream.Read(buffer)
}

func (pStream *IStream) Write(buffer []byte) (uint32, error) {
	pSequentialStream := (*ISequentialStream)(unsafe.Pointer(pStream))

	return pSequentialStream.Write(buffer)
}

func (pSequentialStream *ISequentialStream) Read(buffer []byte) (uint32, error) {
	var (
		pBuffer C.LPVOID
		cb C.ULONG = C.ULONG(len(buffer))
		cbRead C.ULONG
	)

	pBuffer = C.LPVOID(C.malloc(C.size_t(len(buffer))))
	defer C.free(unsafe.Pointer(pBuffer))

	hr := C.sequentialStream_Read((*C.ISequentialStream)(pSequentialStream), pBuffer, cb, &cbRead)

	if hr < 0 {
		return 0, HRESULT(hr)
	}

	raw := (*[1 << 30]C.BYTE)(unsafe.Pointer(pBuffer))[:cbRead:cbRead]
	for i := uint64(0); i < uint64(cbRead); i++ {
		buffer[i] = byte(raw[i])
	}

	return uint32(cbRead), nil
}

func (pSequentialStream *ISequentialStream) Write(buffer []byte) (uint32, error) {
	var (
		pBuffer C.LPVOID
		cb C.ULONG = C.ULONG(len(buffer))
		cbWritten C.ULONG
	)

	pBuffer = C.LPVOID(C.malloc(C.size_t(len(buffer))))
	defer C.free(unsafe.Pointer(pBuffer))

	raw := (*[1 << 30]C.BYTE)(unsafe.Pointer(pBuffer))[:len(buffer):len(buffer)]
	for i := uint64(0); i < uint64(len(buffer)); i++ {
		raw[i] = C.BYTE(buffer[i])
	}

	hr := C.sequentialStream_Write((*C.ISequentialStream)(pSequentialStream), pBuffer, cb, &cbWritten)

	if hr < 0 {
		return 0, HRESULT(hr)
	}

	return uint32(cbWritten), nil
}

func (pUnknown *IUnknown) QueryInterface(iid IID) (unsafe.Pointer, error) {
	var (
		pObject C.LPVOID
	)

	hr := C.unknown_QueryInterface((*C.IUnknown)(pUnknown), iid.toCIID(), &pObject)

	if hr < 0 {
		return nil, HRESULT(hr)
	}

	return unsafe.Pointer(pObject), nil
}

// Convert PWSTR to GoString.
// PWSTR is null-terminated string. So maybe cStr has bigger size than actual length of str by 1 byte.
// PWSTR is utf-16 formatted. It converts utf-16 format into utf-8 format which Go-lang originally supports.
func toGoString(str C.PWSTR, cStr uint32) string {
	raw := (*[1 << 30]C.WCHAR)(unsafe.Pointer(str))[:cStr:cStr]
	utf16Str := make([]uint16, DWORD(cStr))
	for i, wchar := range raw {
		utf16Str[i] = uint16(wchar)
	}

	decodedStr := utf16.Decode(utf16Str)
	goString := make([]byte, 0)
	for _, decodedChar := range decodedStr {
		buffer := make([]byte, 4)
		bytes := utf8.EncodeRune(buffer, decodedChar)

		goString = append(goString, buffer[:bytes]...)
	}

	return string(goString)
}

func allocatePWSTRCoTask(value string) (C.PWSTR, error) {
	decodedStr := make([]rune, 0, len(value))
	utf8Str := []byte(value)
	for len(utf8Str) > 0 {
		r, s := utf8.DecodeRune(utf8Str)
		decodedStr = append(decodedStr, r)
		utf8Str = utf8Str[s:]
	}
	utf16Str := utf16.Encode(decodedStr)

	pwstr := C.CoTaskMemAlloc(C.SIZE_T(C.sizeof_WCHAR * (len(value) + 1)))
	if pwstr == nil {
		return nil, E_POINTER
	}
	raw := (*[1 << 30]C.WCHAR)(pwstr)[:len(value) + 1:len(value) + 1]

	for i, r := range utf16Str {
		raw[i] = C.WCHAR(r)
	}
	raw[len(value)] = 0

	return C.PWSTR(pwstr), nil
}

// Allocate memory.
// Must be unallocated after being used.
// PWSTR is null-terminated string so it allocates memory size of len(value) + 1
func allocatePWSTR(value string) (C.PWSTR, error) {
	decodedStr := make([]rune, 0, len(value))
	utf8Str := []byte(value)
	for len(utf8Str) > 0 {
		r, s := utf8.DecodeRune(utf8Str)
		decodedStr = append(decodedStr, r)
		utf8Str = utf8Str[s:]
	}
	utf16Str := utf16.Encode(decodedStr)

	pwstr := C.malloc(C.size_t(C.sizeof_WCHAR * (len(value) + 1)))
	if pwstr == nil {
		return nil, E_POINTER
	}
	raw := (*[1 << 30]C.WCHAR)(pwstr)[:len(value) + 1:len(value) + 1]

	for i, r := range utf16Str {
		raw[i] = C.WCHAR(r)
	}
	raw[len(value)] = 0

	return C.PWSTR(pwstr), nil
}

func wcslen(pwstr C.PWSTR) uint32 {
	return uint32(C.wcslen((*C.wchar_t)(unsafe.Pointer(pwstr))))
}

/*
static int wcslen(WCHAR* wstr)
{
  char str[MB_CUR_MAX];
  char msg[80];
  WCHAR* wptr;
  int nbytes, i = 0;

  wptr = wstr;

  if ( wptr == NULL )
     return 0;

  while ( *wptr != L'\0' ) {
     nbytes = wctomb ( str, *wptr );
     i += nbytes; (char*)wptr += nbytes;
  }
  return i;
}
 */
// Implementation of wcslen function.
//func wcslen(pwstr C.PWSTR) int {
//	s := unsafe.Pointer(pwstr)
//	p := s
//	for *(*C.WCHAR)(p) != 0 {
//		p = unsafe.Pointer(uintptr(p) + uintptr(1))
//	}
//
//	return int(uintptr(p) - uintptr(s))
//}

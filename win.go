package gowpd

// #include <propidl.h>
// //#include <libgowpd.h>
import "C"
import (
	"unsafe"
	"log"
)

type PropVariant C.PROPVARIANT

func (prop *PropVariant) GetType() VARTYPE {
	return VARTYPE(*(*uint16)(unsafe.Pointer(prop)))
}

func (prop *PropVariant) GetError() uint32 {
	return uint32(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(prop)) + uintptr(8))))
}

func (prop *PropVariant) Set(value interface{}) error {
	switch value.(type) {
	case string:
		*(*uint16)(unsafe.Pointer(prop)) = uint16(VT_LPWSTR)

		strValue := value.(string)
		pStrValue, err := allocatePWSTRCoTask(strValue)
		if err != nil {
			return err
		}

		*(*C.PWSTR)(unsafe.Pointer(uintptr(unsafe.Pointer(prop)) + uintptr(8))) = pStrValue
		return nil
	default:
		return E_UNEXPECTED
	}
}

func (prop *PropVariant) Init() {
	C.memset(unsafe.Pointer(prop), 0, C.sizeof_PROPVARIANT)
}

func (prop *PropVariant) Clear() {
	C.PropVariantClear((*C.PROPVARIANT)(prop))
}

func PropTest() {
	var prop PropVariant
	prop.Init()

	prop.Set("Test")

	log.Printf("size: %d\n", C.sizeof_VARTYPE)
	log.Printf("size: %d\n", C.sizeof_WORD)
	log.Printf("size: %d\n", C.sizeof_PROPVARIANT)
	log.Printf("size: %d\n", C.sizeof_ULARGE_INTEGER)
	log.Printf("size: %d\n", C.sizeof_DECIMAL)
	log.Printf("size: %d\n", C.sizeof_SAFEARRAY)
	log.Printf("size: %d\n", C.sizeof_LPVOID)

	prop.Clear()

	pwstr := C.PWSTR(C.CoTaskMemAlloc(C.SIZE_T(C.sizeof_WCHAR * 5)))

	log.Printf("CoTaskMemAlloc: %p\n", pwstr)

	C.CoTaskMemFree(C.LPVOID(pwstr))
}

//func HRESULTErrorMessage(hr HRESULT) string {
//	allocatePWSTR("")
//
//	str, err := FormatMessage(
//		C.FORMAT_MESSAGE_FROM_SYSTEM,
//		nil,
//		uint32(hr),
//		0,
//		100,
//		nil)
//
//	if err != nil {
//		panic(err)
//	}
//
//	return str
//}

//func FormatMessage(flags uint32, lpSource unsafe.Pointer, messageID uint32, languageID uint32, nSize uint32, arguments interface{}) (string, error) {
//	var (
//		lpBuffer C.PWSTR
//	)
//
//	result := C.FormatMessage(
//		C.DWORD(flags),
//		lpSource,
//		C.DWORD(messageID),
//		C.DWORD(languageID), // Default language
//		lpBuffer,
//		C.DWORD(nSize),
//		nil)
//
//	// if succeeded, result is length of buffer.
//	// if failed, result is zero. To get more information, call GetLastError().
//	if result == 0 {
//		return "", S_FALSE
//	}
//
//	str := toGoString(lpBuffer, uint32(result))
//
//	return str, nil
//}

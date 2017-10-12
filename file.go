package gowpd

/*
#cgo windows amd64 LDFLAGS: -lshlwapi

#include <libgowpd.h>
#include <shlwapi.h>
*/
import "C"
import "unsafe"

func SHCreateStreamOnFile(fileName string, mode uint32) (*IStream, error) {
	var (
		pFileName C.PWSTR
		grfMode C.DWORD = C.DWORD(mode)
		pStream *C.IStream
	)

	pFileName, err := allocatePWSTR(fileName)
	if err != nil {
		return nil, err
	}
	defer C.free(unsafe.Pointer(pFileName))

	hr := C.SHCreateStreamOnFileW(pFileName, grfMode, &pStream)

	if hr < 0 {
		return nil, HRESULT(hr)
	}

	return (*IStream)(pStream), nil
}

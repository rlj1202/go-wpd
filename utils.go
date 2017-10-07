package gowpd

import (
	"path/filepath"
	"log"
)

func GetRequiredPropertiesForAllContentTypes(pObjectProperties *IPortableDeviceValues, parentObjectID, filePath string, pFileStream *IStream) error {
	log.Println("GetRequiredPropertiesForAllContentTypes(): Ready")

	err := pObjectProperties.SetStringValue(WPD_OBJECT_PARENT_ID, parentObjectID)
	if err != nil {
		return err
	}

	if pFileStream == nil {
		return E_POINTER
	}
	statstg, err := pFileStream.Stat(STATFLAG_NONAME)
	if err != nil {
		return err
	}

	err = pObjectProperties.SetUnsignedLargeIntegerValue(WPD_OBJECT_SIZE, statstg.cbSize)
	if err != nil {
		return err
	}

	originalFileName := filepath.Base(filePath)
	ext := filepath.Ext(filePath)
	pObjectProperties.SetStringValue(WPD_OBJECT_ORIGINAL_FILE_NAME, originalFileName)
	pObjectProperties.SetStringValue(WPD_OBJECT_NAME, originalFileName[:len(originalFileName) - len(ext)])

	return nil
}

func GetRequiredPropertiesForContentType(contentType GUID, parentObjectID, filePath string, pFileStream *IStream) (*IPortableDeviceValues, error) {
	log.Println("GetRequiredPropertiesForContentType(): Ready")

	pObjectProperties, err := CreatePortableDeviceValues()
	if pObjectProperties == nil {
		return nil, E_UNEXPECTED
	}

	if err != nil {
		return nil, err
	}

	err = GetRequiredPropertiesForAllContentTypes(pObjectProperties, parentObjectID, filePath, pFileStream)
	if err != nil {
		return nil, err
	}

	switch contentType {
	case WPD_CONTENT_TYPE_IMAGE:
		err = GetRequiredPropertiesForImageContentTypes(pObjectProperties)
	case WPD_CONTENT_TYPE_AUDIO:
		err = GetRequiredPropertiesForMusicContentTypes(pObjectProperties)
	case WPD_CONTENT_TYPE_CONTACT:
		err = GetRequiredPropertiesForContactContentTypes(pObjectProperties)
	}
	if err != nil {
		return nil, err
	}

	result, err := pObjectProperties.QueryInterface(IID_IPortableDeviceValues)
	if err != nil {
		return nil, err
	}

	return (*IPortableDeviceValues)(result), nil
}

func GetRequiredPropertiesForImageContentTypes(pObjectProperties *IPortableDeviceValues) error {
	err := pObjectProperties.SetGuidValue(WPD_OBJECT_CONTENT_TYPE, WPD_CONTENT_TYPE_IMAGE)
	if err != nil {
		return err
	}

	err = pObjectProperties.SetGuidValue(WPD_OBJECT_FORMAT, WPD_OBJECT_FORMAT_EXIF)
	if err != nil {
		return err
	}

	return nil
}

func GetRequiredPropertiesForMusicContentTypes(pObjectProperties *IPortableDeviceValues) error {
	err := pObjectProperties.SetGuidValue(WPD_OBJECT_CONTENT_TYPE, WPD_CONTENT_TYPE_AUDIO)
	if err != nil {
		return err
	}

	err = pObjectProperties.SetGuidValue(WPD_OBJECT_FORMAT, WPD_OBJECT_FORMAT_WMA)
	if err != nil {
		return err
	}

	return nil
}

func GetRequiredPropertiesForContactContentTypes(pObjectProperties *IPortableDeviceValues) error {
	err := pObjectProperties.SetGuidValue(WPD_OBJECT_CONTENT_TYPE, WPD_CONTENT_TYPE_CONTACT)
	if err != nil {
		return err
	}

	err = pObjectProperties.SetGuidValue(WPD_OBJECT_FORMAT, WPD_OBJECT_FORMAT_VCARD2)
	if err != nil {
		return err
	}

	return nil
}

func StreamCopy(pDestStream *IStream, pSourceStream *IStream, cbTransferSize uint32) (uint32, error) {
	pObjectData := make([]byte, cbTransferSize)

	cbTotalBytesRead := uint32(0)
	cbTotalBytesWritten := uint32(0)

	for {
		cbBytesRead, err := pSourceStream.Read(pObjectData)
		if err != nil {
			panic(err)
		}
		log.Printf("StreamCopy(): Read %d bytes.\n", cbBytesRead)

		if cbBytesRead <= 0 {
			break
		}

		cbTotalBytesRead += cbBytesRead
		cbBytesWritten, err := pDestStream.Write(pObjectData[:cbBytesRead])
		if err != nil {
			panic(err)
		}
		cbTotalBytesWritten += cbBytesWritten
	}

	return cbTotalBytesWritten, nil
}

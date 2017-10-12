package gowpd

import (
	"path/filepath"
	"log"
)

// Set necessary properties for all content type.
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

// Return required properties for content type to transfer data.
// Properties which will be set necessarily are WPD_OBJECT_PARENT_ID, WPD_OBJECT_SIZE,
// WPD_OBJECT_ORIGINAL_FILE_NAME, WPD_OBJECT_NAME and properties which will be set optionally are
// WPD_OBJECT_CONTENT_TYPE, WPD_OBJECT_FORMAT.
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

// Set properties for image type.
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

// Set properties for music type.
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

// Set properties for contact type.
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

// Copy the data from pSourceStream to pDestStream. cbTransferSize is buffer size temporarily to store data.
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

// Reads a string property from the IPortableDeviceProperties
// interface and returns it as string.
func GetStringValue(pProperties *IPortableDeviceProperties, objectID string, key PropertyKey) (string, error) {
	// 1) Create an IPortableDeviceKeyCollection interface to hold
	// the property key we wish to read.
	pPropertiesToRead, err := CreatePortableDeviceKeyCollection()
	if err != nil {
		return "", err
	}

	// 2) Populate the IPortableDeviceKeyCollection with the keys
	// we wish to read.
	// NOTE: We are not handling any special error cases here so
	// we can proceed with adding as many of the target properties
	// as we can.
	err = pPropertiesToRead.Add(key)

	// 3) Call GetValues() passing the collection of specified
	// PROPERTYKEYs.
	pObjectProperties, err := pProperties.GetValues(objectID, pPropertiesToRead)
	if err != nil {
		return "", err
	}

	// 4) Extract the string value from the returned property
	// collections
	stringValue, err := pObjectProperties.GetStringValue(key)
	if err != nil {
		return "", err
	}

	return stringValue, nil
}

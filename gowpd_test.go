package gowpd_test

import (
	"fmt"
	"github.com/rlj1202/go-wpd"
	"log"
	"testing"
)

func TestAll(t *testing.T) {
	gowpd.Initialize()

	pPortableDeviceManager, err := gowpd.CreatePortableDeviceManager()
	if err != nil {
		panic(err)
	}

	deviceIDs, err := pPortableDeviceManager.GetDevices()
	if err != nil {
		panic(err)
	}

	pClientInfo, err := gowpd.CreatePortableDeviceValues()
	if err != nil {
		panic(err)
	}
	pClientInfo.SetStringValue(gowpd.WPD_CLIENT_NAME, "libgowpd")
	pClientInfo.SetUnsignedIntegerValue(gowpd.WPD_CLIENT_MAJOR_VERSION, 1)
	pClientInfo.SetUnsignedIntegerValue(gowpd.WPD_CLIENT_MINOR_VERSION, 0)
	pClientInfo.SetUnsignedIntegerValue(gowpd.WPD_CLIENT_REVISION, 2)

	for i, id := range deviceIDs {
		friendlyName, err := pPortableDeviceManager.GetDeviceFriendlyName(id)
		if err != nil {
			panic(err)
		}
		manufacturer, err := pPortableDeviceManager.GetDeviceManufacturer(id)
		if err != nil {
			panic(err)
		}
		description, err := pPortableDeviceManager.GetDeviceDescription(id)
		if err != nil {
			panic(err)
		}

		fmt.Printf("[%d]\n", i)
		fmt.Printf("\tdeviceName:   %s\n", friendlyName)
		fmt.Printf("\tmanufacturer: %s\n", manufacturer)
		fmt.Printf("\tdescription:  %s\n", description)

		if friendlyName != "SANDISK" {
			gowpd.FreeDeviceID(id)
			continue
		}

		pPortableDevice, err := gowpd.CreatePortableDevice()
		if err != nil {
			panic(err)
		}

		// Establish a connection
		err = pPortableDevice.Open(id, pClientInfo)
		if err != nil {
			panic(err)
		}

		// objectId where the file will be transferred under.
		objectID := "F:\\"

		// path to selected file to transfer to device.
		filePath := "E:\\RedLaboratory\\Media\\Picture\\result.png"
		filePath = "E:\\test.md"

		// open file as IStream.
		pFileStream, err := gowpd.SHCreateStreamOnFile(filePath, 0)
		if err != nil {
			panic(err)
		}

		// acquire properties needed to transfer file to device
		pObjectProperties, err := gowpd.GetRequiredPropertiesForContentType(gowpd.WPD_CONTENT_TYPE_IMAGE, objectID, filePath, pFileStream)
		if err != nil {
			panic(err)
		}

		// get stream to device
		content, err := pPortableDevice.Content()
		if err != nil {
			panic(err)
		}
		pTempStream, cbTransferSize, err := content.CreateObjectWithPropertiesAndData(pObjectProperties)
		if err != nil {
			panic(err)
		}

		// convert pTempStream to PortableDeviceDataStream to use more method e.g newly created object id.
		_pFinalObjectDataStream, err := pTempStream.QueryInterface(gowpd.IID_IPortableDeviceDataStream)
		if err != nil {
			panic(err)
		}
		pFinalObjectDataStream := (*gowpd.IPortableDeviceDataStream)(_pFinalObjectDataStream)

		// copy data from pFileStream to pFinalObjectDataStream
		cbBytesWritten, err := gowpd.StreamCopy((*gowpd.IStream)(_pFinalObjectDataStream), pFileStream, cbTransferSize)
		if err != nil {
			panic(err)
		}
		// call commit method to notice device that transferring data is finished.
		err = pFinalObjectDataStream.Commit(0)
		if err != nil {
			panic(err)
		}

		newlyCreatedObjectID, err := pFinalObjectDataStream.GetObjectID()
		if err != nil {
			panic(err)
		}
		log.Printf("\"%s\" has been transferred to device successfully: %d\n", newlyCreatedObjectID, cbBytesWritten)

		// transferring is finished. release the deviceID.
		gowpd.FreeDeviceID(id)
		// release device interface too.
		pPortableDevice.Release()
	}

	pPortableDeviceManager.Release()
	gowpd.Uninitialize()
}

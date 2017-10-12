package gowpd_test

import (
	"github.com/rlj1202/go-wpd"
	"log"
)

func ExampleTransferToPC() {
	gowpd.Initialize()

	mng, err := gowpd.CreatePortableDeviceManager()
	if err != nil {
		panic(err)
	}

	deviceIDs, err := mng.GetDevices()
	if err != nil {
		panic(err)
	}

	clientInfo, err := gowpd.CreatePortableDeviceValues()
	if err != nil {
		panic(err)
	}
	clientInfo.SetStringValue(gowpd.WPD_CLIENT_NAME, "libgowpd")
	clientInfo.SetUnsignedIntegerValue(gowpd.WPD_CLIENT_MAJOR_VERSION, 1)
	clientInfo.SetUnsignedIntegerValue(gowpd.WPD_CLIENT_MINOR_VERSION, 0)
	clientInfo.SetUnsignedIntegerValue(gowpd.WPD_CLIENT_REVISION, 2)

	// object ID which will be transferred to PC.
	targetObjectID := "F:\\test.txt"
	// location where file will be transferred into.
	targetDestination := "E:\\test.txt"

	for _, id := range deviceIDs {
		portableDevice, err := gowpd.CreatePortableDevice()
		if err != nil {
			panic(err)
		}

		portableDevice.Open(id, clientInfo)

		content, err := portableDevice.Content()
		if err != nil {
			panic(err)
		}
		resources, err := content.Transfer()
		if err != nil {
			panic(err)
		}

		objectDataStream, optimalTransferSize, err := resources.GetStream(targetObjectID, gowpd.WPD_RESOURCE_DEFAULT, gowpd.STGM_READ)
		if err != nil {
			panic(err)
		}

		pFinalFileStream, err := gowpd.SHCreateStreamOnFile(targetDestination, gowpd.STGM_CREATE | gowpd.STGM_WRITE)
		if err != nil {
			panic(err)
		}

		totalBytesWritten, err := gowpd.StreamCopy(pFinalFileStream, objectDataStream, optimalTransferSize)
		if err != nil {
			panic(err)
		}

		err = pFinalFileStream.Commit(0)
		if err != nil {
			panic(err)
		}

		log.Printf("Total bytes written: %d\n", totalBytesWritten)

		gowpd.FreeDeviceID(id)
		portableDevice.Release()
	}

	mng.Release()
	gowpd.Uninitialize()

	// Output:
}
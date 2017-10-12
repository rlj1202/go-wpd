package gowpd_test

import (
	"github.com/rlj1202/go-wpd"
	"log"
)

func ExampleDeleteFromDevice() {
	gowpd.Initialize()

	mng, err := gowpd.CreatePortableDeviceManager()
	if err != nil {
		panic(err)
	}

	devices, err := mng.GetDevices()
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

	// objectID which will be deleted from the device.
	targetObjectID := "F:\\test.txt"

	for _, id := range devices {
		portableDevice, err := gowpd.CreatePortableDevice()
		if err != nil {
			panic(err)
		}

		err = portableDevice.Open(id, clientInfo)
		if err != nil {
			panic(err)
		}

		content, err := portableDevice.Content()
		if err != nil {
			panic(err)
		}

		pObjectsToDelete, err := gowpd.CreatePortableDevicePropVariantCollection()

		pv := new(gowpd.PropVariant)
		pv.Init()
		pv.Set(targetObjectID)
		err = pObjectsToDelete.Add(pv)
		if err != nil {
			panic(err)
		}
		results, err := content.Delete(gowpd.PORTABLE_DEVICE_DELETE_NO_RECURSION, pObjectsToDelete)
		if err != nil {
			count, err := results.GetCount()
			if err != nil {
				panic(err)
			}
			log.Printf("Count: %d\n", count)
			result, err := results.GetAt(0)
			if err != nil {
				panic(err)
			}
			log.Printf("Type: %d\n", result.GetType())
			if result.GetType() == gowpd.VT_ERROR {
				log.Printf("error: %#x\n", result.GetError())
			}

			panic(err)
		}

		pv.Clear()
		gowpd.FreeDeviceID(id)
	}

	mng.Release()
	gowpd.Uninitialize()

	// Output:
}
package gowpd_test

import (
	"fmt"
	"github.com/rlj1202/go-wpd"
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

	for _, id := range deviceIDs {
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

		fmt.Println(friendlyName)
		fmt.Println(manufacturer)
		fmt.Println(description)

		pPortableDevice, err := gowpd.CreatePortableDevice()
		if err != nil {
			panic(err)
		}

		err = pPortableDevice.Open(id, pClientInfo)
		if err != nil {
			panic(err)
		}

		//objectID := "F:\"
		_, err = pPortableDevice.Content()
		if err != nil {
			panic(err)
		}

		// select file to transfer to device.

		// open file as IStream

		// acquire properties needed to transfer file to device

		// transfer file to device

		//var pObjectProperties *gowpd.IPortableDeviceValues
		//
		//_, err = content.CreateObjectWithPropertiesAndData(pObjectProperties)
		//if err != nil {
		//	panic(err)
		//}

		gowpd.FreeDeviceID(id)
		pPortableDevice.Release()
	}

	pPortableDeviceManager.Release()
	gowpd.Uninitialize()
}
package main

import (
	"fmt"
	"github.com/rlj1202/go-wpd/gowpd"
)

func main () {
	gowpd.Initialize()

	pPortableDeviceManager, err := gowpd.CreatePortableDeviceManager()
	if err != nil {
		panic(err)
	}

	pPortableDevice, err := gowpd.CreatePortableDevice()
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
	pClientInfo.GetStringValue(gowpd.WPD_CLIENT_NAME)
	pClientInfo.GetUnsignedIntegerValue(gowpd.WPD_CLIENT_MAJOR_VERSION)
	pClientInfo.GetUnsignedIntegerValue(gowpd.WPD_CLIENT_MINOR_VERSION)
	pClientInfo.GetUnsignedIntegerValue(gowpd.WPD_CLIENT_REVISION)

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

		err = pPortableDevice.Open(id, pClientInfo)
		if err != nil {
			panic(err)
		}

		content, err := pPortableDevice.Content()
		if err != nil {
			panic(err)
		}
		_, err = content.EnumObjects(gowpd.WPD_DEVICE_OBJECT_ID)
		if err != nil {
			panic(err)
		}

		gowpd.FreeDeviceID(id)
	}

	pPortableDevice.Release()
	pPortableDeviceManager.Release()

	gowpd.Uninitialize()
}
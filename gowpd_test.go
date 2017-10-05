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

		pPortableDevice, err := gowpd.CreatePortableDevice()
		if err != nil {
			panic(err)
		}

		err = pPortableDevice.Open(id, pClientInfo)
		if err != nil {
			panic(err)
		}

		content, err := pPortableDevice.Content()
		if err != nil {
			panic(err)
		}
		pEnumObjects, err := content.EnumObjects(gowpd.WPD_DEVICE_OBJECT_ID)
		if err != nil {
			panic(err)
		}
		objects, err := pEnumObjects.Next(10)
		if err != nil {
			panic(err)
		}
		for _, obj := range objects {
			fmt.Println(obj)

			test, _ := content.EnumObjects(obj)
			tests, _ := test.Next(30)
			for _, t := range tests {
				fmt.Println(t)
			}
		}

		gowpd.FreeDeviceID(id)
		pPortableDevice.Release()
	}

	pPortableDeviceManager.Release()

	gowpd.Uninitialize()
}
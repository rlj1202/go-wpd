package gowpd_test

import (
	"github.com/rlj1202/go-wpd"
	"log"
)

func RecursiveEnumerate(parentObjectID string, content *gowpd.IPortableDeviceContent) {
	enum, err := content.EnumObjects(parentObjectID)
	if err != nil {
		panic(err)
	}

	objectIDs := make([]string, 0)
	for {
		tmp, err := enum.Next(10)
		if err != nil {
			panic(err)
		}
		if len(tmp) == 0 {
			break
		}
		objectIDs = append(objectIDs, tmp...)
	}

	for _, objectID := range objectIDs {
		log.Println(objectID)
	}

	for _, objectID := range objectIDs {
		RecursiveEnumerate(objectID, content)
	}
}

func Example_contentEnumerate() {
	gowpd.Initialize()

	mng, err := gowpd.CreatePortableDeviceManager()
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

	deviceIDs, err := mng.GetDevices()
	if err != nil {
		panic(err)
	}

	for _, deviceID := range deviceIDs {
		device, err := gowpd.CreatePortableDevice()
		if err != nil {
			panic(err)
		}

		err = device.Open(deviceID, pClientInfo)
		if err != nil {
			panic(err)
		}

		content, err := device.Content()
		if err != nil {
			panic(err)
		}

		RecursiveEnumerate(gowpd.WPD_DEVICE_OBJECT_ID, content)

		gowpd.FreeDeviceID(deviceID)
	}

	gowpd.Uninitialize()
}

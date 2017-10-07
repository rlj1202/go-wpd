package gowpd_test

import (
	"github.com/rlj1202/go-wpd"
	"log"
)

func Example_deviceEnumerate() {
	gowpd.Initialize()

	mng, err := gowpd.CreatePortableDeviceManager()
	if err != nil {
		panic(err)
	}

	deviceIDs, err := mng.GetDevices()
	if err != nil {
		panic(err)
	}

	for i, deviceID := range deviceIDs {
		friendlyName, err := mng.GetDeviceFriendlyName(deviceID)
		if err != nil {
			panic(err)
		}
		manufacturer, err := mng.GetDeviceManufacturer(deviceID)
		if err != nil {
			panic(err)
		}
		description, err := mng.GetDeviceDescription(deviceID)
		if err != nil {
			panic(err)
		}

		log.Printf("[%d]:\n", i)
		log.Printf("\tName:         %s\n", friendlyName)
		log.Printf("\tManufacturer: %s\n", manufacturer)
		log.Printf("\tDescription:  %s\n", description)

		gowpd.FreeDeviceID(deviceID)
	}

	gowpd.Uninitialize()
}
package gowpd_test

import "github.com/rlj1202/go-wpd"

func Example_transferToDevice() {
	gowpd.Initialize()

	mng, err := gowpd.CreatePortableDeviceManager()
	if err != nil {
		panic(err)
	}
	deviceIDs, err := mng.GetDevices()
	if err != nil {
		panic(err)
	}


	for _, id := range deviceIDs {
		gowpd.FreeDeviceID(id)
	}

	gowpd.Uninitialize()
}
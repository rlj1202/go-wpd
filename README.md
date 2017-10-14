# go-wpd

[![GoDoc](https://godoc.org/github.com/rlj1202/go-wpd?status.svg)](https://godoc.org/github.com/rlj1202/go-wpd)
[![Go Report Card](https://goreportcard.com/badge/github.com/rlj1202/go-wpd)](https://goreportcard.com/report/github.com/rlj1202/go-wpd)
[![Build status](https://ci.appveyor.com/api/projects/status/yei4t5h2rq1cmsao?svg=true)](https://ci.appveyor.com/project/rlj1202/go-wpd)

Window Portable Device binding for Go language.

## Examples

Enumerate devices

```go
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
```

package handler

import (
	"log"
	"time"
)

var _devices map[string]STDevice

//STDevice ..
type STDevice struct {
	Name string `json:"name"`
	Host string `json:"host"`
	Time string `json:"time"`
}

func init() {
	_devices = make(map[string]STDevice)
}

//SaveDeviceMessage ..
func SaveDeviceMessage(device string, host string) bool {
	if len(device) < 2 {
		return false
	}

	timeStr := time.Now().Format("2006-01-02 15:04:05")

	log.Println("device:" + device + " - " + host)

	var d STDevice
	d.Name = device
	d.Host = host
	d.Time = timeStr

	_devices[device] = d
	return true
}

//GetDevices ..
func GetDevices() string {

	str := ""
	for _, keyB := range _devices {
		str += keyB.Name + " - " + keyB.Time + " - " + keyB.Host + "; "
	}

	return str
}

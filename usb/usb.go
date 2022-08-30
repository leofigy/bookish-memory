package usb

import "time"

type Info struct {
	LastTimeUpdated time.Time
	FriendlyName    string
}

type USBInspektor interface {
	GetDevicesHistory() ([]Info, error)
}

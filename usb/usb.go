package usb

import "time"

type Info struct {
	LastTimeUpdated time.Time
	FriendlyName    string
}

type USB interface {
	GetHistory() Info
}

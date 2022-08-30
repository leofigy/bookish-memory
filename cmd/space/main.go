package main

import (
	"fmt"

	fs "gihub.com/leofigy/bookish-memory/fs"
	"gihub.com/leofigy/bookish-memory/usb"
)

func main() {
	inf, err := fs.CalcSpace("")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(inf.GetSpace())

	usbdev := usb.NewInspektor()

	items, err := usbdev.GetDevicesHistory()
	if err != nil {
		fmt.Println(err)
	} else {
		for _, r := range items {
			fmt.Println(r.FriendlyName)
		}
	}
}

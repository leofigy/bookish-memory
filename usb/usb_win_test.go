//go:build windows

package usb

import (
	"fmt"
	"testing"

	"golang.org/x/sys/windows/registry"
)

func TestGetUSBK(t *testing.T) {

	key, err := GetUSBKey(registryEntry)

	if err != nil && err == registry.ErrNotExist {
		fmt.Println("soft error this machine has no usb devices connected yet")
		fmt.Println(err)
		return
	}

	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	// some usb registry entries found
	defer key.Close()

	info, err := key.Stat()
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	fmt.Println("devices", info.SubKeyCount)

	if info.SubKeyCount <= 0 {
		fmt.Println("no devices here")
		return
	}

	// getting subkeys
	subkeys, err := key.ReadSubKeyNames(int(info.SubKeyCount))

	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	for _, k := range subkeys {
		fmt.Println(k)
	}

}

//go:build windows

package usb

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/sys/windows/registry"
)

const (
	registryEntry = "SYSTEM\\CurrentControlSet\\Enum\\USBSTOR"
	readAccess    = registry.ENUMERATE_SUB_KEYS | registry.QUERY_VALUE
)

func GetUSBKey(path string) (*registry.Key, error) {
	reg, err := registry.OpenKey(registry.LOCAL_MACHINE, path, readAccess)

	if err != nil {
		// soft error
		if err == registry.ErrNotExist {
			log.Println(err)
			return nil, err
		}

		fmt.Printf("Failed to open '%s' Error: %v", path, err)
		return nil, err
	}

	return &reg, nil
}

type Inspector struct{}

func NewInspektor() USBInspektor {
	return &Inspector{}
}

func (s *Inspector) GetDevicesHistory() ([]Info, error) {
	devices := make([]Info, 0)

	key, err := GetUSBKey(registryEntry)

	if err != nil && err == registry.ErrNotExist {
		fmt.Println("soft error this machine has no usb devices connected yet")
		fmt.Println(err)
		return devices, nil
	}

	defer key.Close()

	info, err := key.Stat()
	if err != nil {
		return devices, err
	}

	if info.SubKeyCount <= 0 {
		return devices, nil
	}

	subkeys, err := key.ReadSubKeyNames(
		int((info.SubKeyCount)),
	)

	if err != nil {
		return devices, err
	}

	for _, k := range subkeys {
		// TODO fetch the subkey to get the actual time
		devices = append(
			devices, Info{
				time.Now(),
				k,
			},
		)
	}

	return devices, nil
}

//go:build windows

package usb

import (
	"fmt"
	"log"

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

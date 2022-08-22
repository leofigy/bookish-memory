//go:build windows

package disk

import (
	"golang.org/x/sys/windows"
)

var getDiskFreeSpaceExW = windows.MustLoadDLL("Kernel32.dll").MustFindProc("GetDiskFreeSpaceExW")

type WinUnit struct{}

func (w *WinUnit) getSpace() Space {
	return Space{}
}

func CalcSpace(wd string) (Unit, error) {
	return nil, nil
}

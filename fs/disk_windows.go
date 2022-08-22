//go:build windows

package disk

import (
	"golang.org/x/sys/windows"
)

var getDiskFreeSpaceExW = windows.MustLoadDLL("Kernel32.dll").MustFindProc("GetDiskFreeSpaceExW")

type WinUnit struct{}

func Space(wd string) (int64, error) {

}

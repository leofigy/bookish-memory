//go:build windows

package disk

import (
	"log"
	"unsafe"

	"golang.org/x/sys/windows"
)

var getDiskFreeSpaceExW = windows.MustLoadDLL("Kernel32.dll").MustFindProc("GetDiskFreeSpaceExW")

type WinUnit struct {
	dir string
	Space
}

func (w *WinUnit) GetSpace() Space {
	var avail, total, free int64

	code, _, err := getDiskFreeSpaceExW.Call(
		uintptr(
			unsafe.Pointer(
				windows.StringToUTF16Ptr((w.dir)),
			),
		),
		uintptr(
			unsafe.Pointer(
				&avail,
			),
		),
		uintptr(
			unsafe.Pointer(
				&total,
			),
		),
		uintptr(
			unsafe.Pointer(
				&free,
			),
		),
	)

	if code == 0 {
		log.Println("bad code")
		log.Println(err)
	}

	w.Total = total
	w.Used = total - free
	return w.Space
}

func CalcSpace(wd string) (Unit, error) {

	if wd == "" {
		wd = "C:/"
	}

	proc := &WinUnit{
		dir: wd,
	}
	return proc, nil
}

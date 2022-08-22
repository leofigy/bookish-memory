//go:build !windows

package disk

import (
	"os"

	"golang.org/x/sys/unix"
)

type UDisk struct{}

func CalcSpace(wd string) (Unit, error) {
	Stats(wd)
	return nil, nil
}

func Stats(wd string) (unix.Statfs_t, error) {
	var (
		err  error
		info unix.Statfs_t
	)
	if wd == "" {
		wd, err = os.Getwd()
		// defaulting
		if err != nil {
			return info, err
		}
	}
	unix.Statfs(wd, &info)
	return info, nil
}

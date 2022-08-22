package disk

import (
	"os"
	"golang.org/x/sys/unix"
)

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

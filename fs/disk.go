//go:build !windows

package disk

import (
	"os"

	"golang.org/x/sys/unix"
)

type NixUnit struct{
	dir string
	info unix.Statfs_t
	Space
}

func (n *NixUnit) GetSpace() Space {
	info, err := Stats(n.dir)
	if err != nil {
		log.Println(err)
		return n.Space,
	}
	n.info = info

	// TODO update calc here for current values

	return n.Space,
}

func CalcSpace(wd string) (Unit, error) {
	proc := &NixUnit{
		dir: wd,
	}
	return proc, nil
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

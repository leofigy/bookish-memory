package fsevents

/*
	FSEvents log
	default locations
	-> /.fseventsd (macos)
	-> /private/var/.fseventsd
	directory contains gzip formats
	Name is last event ID stored in the FSEvent log
	file plus 1.
*/

import (
	"fmt"
	"io/ioutil"
)

func InspectDir(base string) error {
	files, err := ioutil.ReadDir(base)
	if err != nil {
		return err
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
	return nil
}

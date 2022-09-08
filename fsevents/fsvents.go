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

type FSEvent struct {
	ID                        string
	IDHex                     int32
	FullPath                  string
	Filename                  string
	Type                      string
	Flags                     string
	AproxDatesPlusMinusOneDay string
	Mask                      int32
	NodeID                    string
	RecordEndOffSet           string
	Source                    string
	SourceModifiedDate        string
}

type FSEventHandler interface {
	WorkingDir() string
	Inspect(string) error
	Parse(chan (FSEvent))
}

// TODO steps
// check if is a gzip file create a temporary directory to unzip information
// check the header format
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

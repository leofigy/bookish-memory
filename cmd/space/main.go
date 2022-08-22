package main

import (
	"fmt"

	fs "gihub.com/leofigy/bookish-memory/fs"
)

func main() {
	inf, err := fs.Stats("")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(inf.Bavail * uint64(inf.Bsize))
}

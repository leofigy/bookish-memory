package main

import (
	"fmt"

	fs "gihub.com/leofigy/bookish-memory/fs"
)

func main() {
	inf, err := fs.CalcSpace("")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(inf.GetSpace())
}

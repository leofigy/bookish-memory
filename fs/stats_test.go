package disk

import (
	"fmt"
	"testing"
)

func TestSpace(t *testing.T) {
	data, err := CalcSpace("")
	// error reading space
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	// error getting total space less than 0
	if data.GetSpace().Total <= 0 {
		fmt.Println("invalid case space is 0")
		t.FailNow()
	}

	fmt.Println("test passed")
	fmt.Printf("space here is %d", data.GetSpace().Total)
}

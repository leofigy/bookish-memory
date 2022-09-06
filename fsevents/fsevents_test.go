package fsevents

import (
	"fmt"
	"os"
	"testing"
)

func TestInspectDir(t *testing.T) {

	base := os.Getenv("GO_TEST_MAC_FSEVENT_DIR")
	// GO_TEST_MAC_FSEVENT_DIR
	// /Users/leofigy/repos/bookish-memory/tmp/.fseventsd
	if len(base) == 0 {
		fmt.Println("GO_TEST_MAC_FSEVENT_DIR not set skipping")
		return
	}

	err := InspectDir(base)

	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

}

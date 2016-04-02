package gocliboil

import (
	"fmt"
	"os"
)

func Exit(code int) {
	if code == 0 {
		fmt.Println("exiting with 0")
	} else {
		fmt.Fprintf(os.Stderr, "exiting with %d\n", code)
	}
	os.Exit(code)
}

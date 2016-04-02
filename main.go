package main

import (
	"flag"
	"github.com/mmb/gocliboil/gocliboil"
)

func main() {
	var exitCode int

	flag.IntVar(&exitCode, "exitCode", 0, "exit code to return")
	flag.Parse()

	gocliboil.Exit(exitCode)
}

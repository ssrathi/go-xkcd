/*
CLI interface to read XKCD comics, implemented in Go language.

This project is part of a learning exercise to implement a XKCD comic reader
in Go language.

This CLI can be used to download the latest XKCD, or a specified XKCD with a number,
or a random XKCD from all published ones.

See "go-xkcd --help" for usage and examples.
*/
package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/ssrathi/go-xkcd/xkcd"
)

func init() {
	// Enable logging only if a specific ENV variable is set.
	if os.Getenv("GOXKCD_DBG") != "1" {
		log.SetOutput(ioutil.Discard)
		log.SetFlags(0)
	} else {
		// Print file and line numbers in each log line.
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	}
}

func main() {
	xkcd.Execute()
}

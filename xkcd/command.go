// Package cmd is the entry point for gogit command line parsing.
package xkcd

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	comicNum int,
	getRandom bool,
)

// Execute parses CLI arguments and execute with the given options.
func Execute() {
	progName := os.Args[0]
	args := os.Args[1:]

	// Prepare flags for user input.
	flag.IntVar(&comicNum, "n", 0, "Get a specific comic number")
	flag.BoolVar(&getRandom, "r", false, "Get a random comic")

	flag.Usage = func() {
		fmt.Printf("go-xkcd - Read XKCD from command line.\n\n")
		fmt.Printf("usage: %s [<args>]\n", progName)
		flag.PrintDefaults()
	}

	flag.Parse()

	// Create a REST client to get the comic from XKCD website.
	client := NewXkcdClient()
	comic, err := client.GetComic(comicNum)
	fmt.Printf("%+v\n", comic)
}

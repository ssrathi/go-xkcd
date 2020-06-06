package xkcd

import (
	"flag"
	"fmt"
	"os"
)

var (
	comicNum  int
	getRandom bool
	outputFmt string
	savePath  string
)

// Check validates if there is an error and exits the program if any.
func Check(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

// Execute parses CLI arguments and execute with the given options.
func Execute() {
	progName := os.Args[0]

	// Prepare flags for user input.
	flag.IntVar(&comicNum, "n", 0, "Get a specific comic number")
	flag.BoolVar(&getRandom, "r", false, "Get a random comic")
	flag.StringVar(&outputFmt, "o", "text", "Output format (text/json)")
	flag.StringVar(&savePath, "s", ".", "Path to save the comic image")

	flag.Usage = func() {
		fmt.Printf("go-xkcd - Read XKCD from command line.\n\n")
		fmt.Printf("usage: %s [<args>]\n", progName)
		flag.PrintDefaults()
	}

	flag.Parse()
	if outputFmt != "text" && outputFmt != "json" {
		fmt.Printf("Invalid output format '%s'. Valid values {text/json}\n", outputFmt)
		os.Exit(1)
	}

	// Create a REST client to get the comic from XKCD website.
	client := NewClient()
	comic, err := client.GetComicMetadata(comicNum)
	Check(err)

	savePath, err = client.GetComicImage(comic.Img, savePath)
	Check(err)

	if outputFmt == "text" {
		fmt.Println(comic.PrettyStr())
	} else {
		jsonOutput, err := comic.JSONStr()
		Check(err)
		fmt.Println(jsonOutput)
	}

	fmt.Printf("\nComic image saved at %s\n", savePath)
}

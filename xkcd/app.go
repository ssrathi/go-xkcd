package xkcd

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var (
	comicNum  int
	getRandom bool
	outputFmt string
	savePath  string
)

// Run parses the CLI arguments and executes with the given options.
// It returns the status code (0: success, 1: failure)
func Run() int {
	progName := os.Args[0]

	// Prepare flags for user input.
	flag.IntVar(&comicNum, "n", 0, "Get a specific comic number (overrides -r)")
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
		return 1
	}

	// Create a REST client to get the comic from XKCD website.
	client := NewClient()
	comic, err := client.GetComicMetadata(comicNum)
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}

	// If random comic is asked for, then generate a random comic nunber.
	// Comic 404 doesn't exist as a joke (404 is not-found status code).
	if comicNum == 0 && getRandom {
		rand.Seed(time.Now().UnixNano())
		randNum := rand.Intn(comic.Num + 1)
		for randNum == 404 {
			randNum = rand.Intn(comic.Num + 1)
		}

		if comic, err = client.GetComicMetadata(randNum); err != nil {
			fmt.Println(err.Error())
			return 1
		}
	}

	if savePath, err = client.GetComicImage(comic.Img, savePath); err != nil {
		fmt.Println(err.Error())
		return 1
	}

	var out string
	if outputFmt == "text" {
		out, err = comic.PrettyStr()
	} else {
		out, err = comic.JSONStr()
	}

	if err != nil {
		fmt.Println(err.Error())
		return 1
	}

	fmt.Println(out)
	fmt.Printf("\nComic image saved at %s\n", savePath)
	return 0
}

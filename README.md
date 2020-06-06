[![Builds](https://github.com/ssrathi/go-xkcd/workflows/Build/badge.svg?branch=master)](https://github.com/ssrathi/go-xkcd/actions?query=branch%3Amaster+workflow%3ABuild)
[![Go Report Card](https://goreportcard.com/badge/github.com/ssrathi/go-xkcd)](https://goreportcard.com/report/github.com/ssrathi/go-xkcd)
[![GoDoc](https://godoc.org/github.com/ssrathi/go-xkcd?status.svg)](https://godoc.org/github.com/ssrathi/go-xkcd)

# go-xkcd
CLI interface to read XKCD comics, implemented in Go language

# Usage
```
go-xkcd - Read XKCD from command line.

usage: go-xkcd [<args>]
  -n int
    	Get a specific comic number
  -o string
    	Output format (text/json) (default "text")
  -r	Get a random comic
  -s string
    	Path to save the comic image (default ".")
```

# Examples
```
$ go-xkcd
XKCD Number: 2316
Title: Hair Growth Rate
Date Published: 2020-6-5
Alt Text: Hourly haircuts would be annoying, but they'd be easier to do yourself, since you'd have adjacent hairs as a guide. Growing it out would be a huge pain, though.
Image Link: https://imgs.xkcd.com/comics/hair_growth_rate.png


Comic image saved at <path>/hair_growth_rate.png
```

```
$ go-xkcd -n 1 -o json
{
    "month": "1",
    "num": 1,
    "link": "",
    "year": "2006",
    "news": "",
    "safe_title": "Barrel - Part 1",
    "transcript": "[[A boy sits in a barrel which is floating in an ocean.]]\nBoy: I wonder where I'll float next?\n[[The barrel drifts into the distance. Nothing else can be seen.]]\n{{Alt: Don't we all.}}",
    "alt": "Don't we all.",
    "img": "https://imgs.xkcd.com/comics/barrel_cropped_(1).jpg",
    "title": "Barrel - Part 1",
    "day": "1"
}

Comic image saved at <path>/barrel_cropped_(1).jpg
```

```
$ go-xkcd -r
XKCD Number: 2277
Title: Business Greetings
Date Published: 2020-3-6
Alt Text: We have email and social media now, so we probably don't need to keep exchanging business cards by pressing them gently against each others' faces with an open palm and smearing them around.
Image Link: https://imgs.xkcd.com/comics/business_greetings.png


Comic image saved at <path>/business_greetings.png
```

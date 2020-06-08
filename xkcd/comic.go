// Package xkcd implements a http client to read XKCD JSON API data.
package xkcd

import (
	"encoding/json"
	"fmt"
	"time"
)

// Comic is the XKCD json format for returning a response.
// Generated on https://mholt.github.io/json-to-go/
type Comic struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

// PrettyStr returns a human readable string representation of a XKCD comic.
func (c Comic) PrettyStr() (string, error) {
	dateStr, err := c.Date()
	if err != nil {
		return "", fmt.Errorf("Failed to parse date from server response: "+
			"%s", err.Error())
	}

	out := fmt.Sprintf(
		"XKCD Number: %d\nTitle: %s\nDate Published: %s\nAlt Text: %s\n"+
			"Image Link: %s\n", c.Num, c.Title, dateStr, c.Alt, c.Img)
	return out, nil
}

// JSONStr returns a JSON formatted representation of a XKCD comic.
func (c Comic) JSONStr() (string, error) {
	json, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		return "", fmt.Errorf("Failed to parse server response: %s", err.Error())
	}

	return string(json), nil
}

// Date returns a properly formatted string representation of the publication date.
func (c Comic) Date() (string, error) {
	tm, err := time.Parse("2006-1-2",
		fmt.Sprintf("%s-%s-%s", c.Year, c.Month, c.Day))
	if err != nil {
		return "", err
	}

	return tm.Format("02-Jan-2006"), nil
}

// Package xkcd implements a http client to read XKCD JSON API data.
package xkcd

import (
	"encoding/json"
	"fmt"
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
func (c Comic) PrettyStr() string {
	return fmt.Sprintf(
		"XKCD Number: %d\nTitle: %s\nDate Published: %s-%s-%s\n"+
			"Alt Text: %s\nImage Link: %s\n", c.Num, c.Title, c.Year,
		c.Month, c.Day, c.Alt, c.Img)
}

// JSONStr returns a JSON formatted representation of a XKCD comic.
func (c Comic) JSONStr() (string, error) {
	json, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		return "", fmt.Errorf("Failed to parse server response: %s", err.Error())
	}

	return string(json), nil
}

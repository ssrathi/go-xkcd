package xkcd

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// BaseURL is the XKCD website address.
const BaseURL string = "https://xkcd.com"

// Client is a wrapper on httpClient for fetching XKCD data from their servers.
type Client struct {
	httpClient *http.Client
}

// NewClient returns a new XKCD client.
func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{},
	}
}

// GetComic fetches a specified comic number from the XKCD servers. If 'number' is
// 0, then it fetches the latest comic.
func (xc *Client) GetComic(number int) (Comic, error) {
	var endpoint string
	if number > 0 {
		endpoint = fmt.Sprintf("%s/%d/info.0.json", BaseURL, number)
	} else {
		endpoint = fmt.Sprintf("%s/info.0.json", BaseURL)
	}

	response, err := xc.httpClient.Get(endpoint)
	if err != nil {
		return Comic{}, fmt.Errorf("failed to get comic from server: %s", err.Error())
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return Comic{}, fmt.Errorf("invalid comic number. Failed with %d",
			response.StatusCode)
	}

	var comic Comic
	err = json.NewDecoder(response.Body).Decode(&comic)
	if err != nil {
		return Comic{}, fmt.Errorf("failed to parse server response: %s", err.Error())
	}

	return comic, nil
}

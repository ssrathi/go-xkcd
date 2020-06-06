package xkcd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"path/filepath"
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

// GetComicMetadata fetches a specified comic number from the XKCD servers. If
// 'number' is 0, then it fetches the latest comic.
func (xc *Client) GetComicMetadata(number int) (Comic, error) {
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
		return Comic{}, fmt.Errorf("invalid comic number: %s",
			http.StatusText(response.StatusCode))
	}

	var comic Comic
	err = json.NewDecoder(response.Body).Decode(&comic)
	if err != nil {
		return Comic{}, fmt.Errorf("failed to parse server response: %s", err.Error())
	}

	return comic, nil
}

// GetComicImage downloads and saves the image from the given URL to the given
// path on the directory.
func (xc *Client) GetComicImage(imgURL, dirPath string) (string, error) {
	// Fetch the image and save it to the disk.
	response, err := xc.httpClient.Get(imgURL)
	if err != nil {
		return "", fmt.Errorf("failed to get comic image from server: %s", err.Error())
	}
	defer response.Body.Close()

	imgPath, _ := filepath.Abs(dirPath)
	imgPath = filepath.Join(imgPath, path.Base(imgURL))

	imgData, _ := ioutil.ReadAll(response.Body)
	err = ioutil.WriteFile(imgPath, imgData, 0644)
	if err != nil {
		return "", fmt.Errorf("failed to save comic image to disk: %s", err.Error())
	}

	return imgPath, nil
}

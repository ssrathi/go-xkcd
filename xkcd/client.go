package xkcd

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const BaseURL string = "https://xkcd.com"

type XkcdClient struct {
	client *http.Client,
}

func NewXkcdClient() *XkcdClient {
	return &XkcdClient {
		client: &http.Client{},
	}
}

func (xc *XkcdClient) GetComic(number int) (Comic, error) {
	var endpoint string
	if number > 0 {
		endpoint = fmt.Sprintf("%s/%d/info.0.json", BaseURL, number)
	} else {
		endpoint = fmt.Sprintf("%s/info.0.json", BaseURL))
	}

	response, err := xc.client.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect server: %s", err.Error())
	}
	defer response.Body.Close()

	var comic Comic
	err := json.NewDecoder(response.Body).Decode(&comic)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse server reponse: %s", err.Error())
	}

	return comic, nil
}

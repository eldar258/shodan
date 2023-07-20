package shodan

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const BaseUrl = "https://api.shodan.io"

const key = "PUT_YOU_KEY"

type Client struct {
	apiKey string
}

func New() *Client {
	return &Client{apiKey: key}
}

func (c *Client) Info() (*APIInfo, error)  {
	res, err := http.Get(fmt.Sprintf("%s/api-info?key=%s", BaseUrl, c.apiKey))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var result APIInfo
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, err
}

func (c *Client)HostSearch(q string)  (*HostSearch, error) {
	res, err := http.Get(fmt.Sprintf("%s/shodan/host/search?key=%s&query=%s", BaseUrl, c.apiKey, q))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var result HostSearch //TODO
	return &result, json.NewDecoder(res.Body).Decode(&result)
}

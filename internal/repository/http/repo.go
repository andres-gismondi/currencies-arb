package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

const path = "https://api.swissborg.io/v1/challenge/rates"

var (
	ErrDecodingResponse = errors.New("error decoding response")
	ErrHTTPResponse     = errors.New("error getting response")
)

type Client struct {
	client *http.Client
}

func NewClient() Client {
	return Client{
		client: &http.Client{},
	}
}

func (c Client) Get(ctx context.Context) (map[string]string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", path, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return nil, ErrHTTPResponse
	}

	var currencies map[string]string
	err = json.NewDecoder(res.Body).Decode(&currencies)
	if err != nil {
		return nil, ErrDecodingResponse
	}

	return currencies, nil
}

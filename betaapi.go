package customerio

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type BetaAPIClient struct {
	Key       string
	URL       string
	UserAgent string
	Client    *http.Client
}

// NewBetaAPIClient prepares a client for use with the Customer.io API, see: https://customer.io/docs/api/#apicoreintroduction
// using an App API Key from https://fly.customer.io/settings/api_credentials?keyType=app
func NewBetaAPIClient(key string, opts ...option) *BetaAPIClient {
	client := &BetaAPIClient{
		Key:       key,
		Client:    http.DefaultClient,
		URL:       "https://beta-api.customer.io",
		UserAgent: DefaultUserAgent,
	}

	for _, opt := range opts {
		opt.betaapi(client)
	}
	return client
}

func (c *BetaAPIClient) doRequest(ctx context.Context, verb, requestPath string, body interface{}) ([]byte, int, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return nil, 0, err
	}

	req, err := http.NewRequest(verb, c.URL+requestPath, bytes.NewBuffer(b))
	if err != nil {
		return nil, 0, err
	}

	req = req.WithContext(ctx)

	req.Header.Set("Authorization", "Bearer "+c.Key)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("User-Agent", c.UserAgent)

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	return respBody, resp.StatusCode, nil
}

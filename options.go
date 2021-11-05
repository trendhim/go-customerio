package customerio

import "net/http"

type option struct {
	api     func(*APIClient)
	betaapi func(*BetaAPIClient)
	track   func(*CustomerIO)
}

type region struct {
	ApiURL     string
	BetaApiURL string
	TrackURL   string
}

var (
	RegionUS = region{
		ApiURL:     "https://api.customer.io",
		BetaApiURL: "https://beta-api.customer.io",
		TrackURL:   "https://track.customer.io",
	}
	RegionEU = region{
		ApiURL:     "https://api-eu.customer.io",
		BetaApiURL: "https://beta-api-eu.customer.io",
		TrackURL:   "https://track-eu.customer.io",
	}
)

func WithRegion(r region) option {
	return option{
		api: func(a *APIClient) {
			a.URL = r.ApiURL
		},
		betaapi: func(a *BetaAPIClient) {
			a.URL = r.BetaApiURL
		},
		track: func(c *CustomerIO) {
			c.URL = r.TrackURL
		},
	}
}

func WithHTTPClient(client *http.Client) option {
	return option{
		api: func(a *APIClient) {
			a.Client = client
		},
		betaapi: func(a *BetaAPIClient) {
			a.Client = client
		},
		track: func(c *CustomerIO) {
			c.Client = client
		},
	}
}

func WithUserAgent(ua string) option {
	return option{
		api: func(a *APIClient) {
			a.UserAgent = ua
		},
		betaapi: func(a *BetaAPIClient) {
			a.UserAgent = ua
		},
		track: func(c *CustomerIO) {
			c.UserAgent = ua
		},
	}
}

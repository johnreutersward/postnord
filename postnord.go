// Package postnord provides a client library to access postnords API.
package postnord

import (
	"net/http"
	"net/url"
)

const (
	libraryVersion = 0.1
	userAgent      = "go-postnord/" + libraryVersion
	baseURL        = "http://logistics.postennorden.com/wsp/rest-services/ntt-service-rest/api/"
)

type Client struct {
	// Locale may be set to en, sv, no, fi or da. English (en) is the default locale.
	Locale     string
	ConsumerID string
	UserAgent  string
	baseURL    *url.URL
	httpClient *http.Client
}

// NewClient returns a postnord api access client. If no http client is provided http.DefaultClient will be used.
func NewClient(ConsumerID string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(baseURL)

	c := &Client{
		Locale:     "en",
		UserAgent:  userAgent,
		baseURL:    baseURL,
		ConsumerID: ConsumerID,
		httpClient: httpClient,
	}

	return c
}

type Shipment struct {
}

func (c *Client) Shipment(ID string) (*Shipment, error) {

}

func (c *Client) get(endpoint string, v interface{}) (*http.Request, error) {

}

// Package postnord provides a client library to access postnords API.
package postnord

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

const (
	libraryVersion = "0.1"
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

type TrackingInformationResponse struct {
	shipments []Shipment `xml:"shipments"`
}

type Shipment struct {
	ID string `xml:"shipmentId"`
}

func (c *Client) Shipment(ID string) (*Shipment, error) {
	endp := "shipment.xml?"

	vals := url.Values{}
	vals.Set("id", ID)

	v := new(TrackingInformationResponse)

	_, err := c.get(endp, vals, v)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%v", v)

	return nil, nil
}

func (c *Client) get(endp string, vals url.Values, v interface{}) (*http.Response, error) {
	vals.Set("consumerId", c.ConsumerID)
	vals.Set("locale", c.Locale)

	// can't use url.Values.Encode() because order is not guaranteed
	qs := fmt.Sprintf("id=%s&locale=%s&consumerId=%s", vals.Get("id"), vals.Get("locale"), vals.Get("consumerId"))

	ref, err := url.Parse(endp + qs)
	if err != nil {
		return nil, err
	}

	url := c.baseURL.ResolveReference(ref)

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.UserAgent)

	// debug
	dump, err := httputil.DumpRequestOut(req, false)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(dump[:]))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// debug
	dump, err = httputil.DumpResponse(resp, true)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(dump[:]))

	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("api error HTTP %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	if v != nil {
		err = xml.NewDecoder(resp.Body).Decode(v)
	}

	return resp, err
}

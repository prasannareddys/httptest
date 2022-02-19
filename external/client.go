package external

import (
	"fmt"
	"net/http"
	"net/url"
	"sync"
)

var (
	clientOnce sync.Once
	mClient    *Client
)

const DefaultContentType = "application/json"

type HttpClient interface {
	Get(url string) (*http.Response, error)
}

type Client struct {
	client *http.Client
}

func newClient() *Client {
	return &Client{client: &http.Client{
		Transport: &validURLRoundTripper{
			defaultRoundTripper: http.DefaultTransport,
		},
	}}
}

func GetClient() *Client {
	clientOnce.Do(func() {
		mClient = newClient()
	})
	return mClient
}

func (c *Client) Get(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Content-Type", DefaultContentType)
	if err != nil {
		return nil, fmt.Errorf("Filed to request url %s with error %s ", url, err.Error())
	}

	res, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Filed to request url %s with error %s ", url, err.Error())
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Filed to request url %s with status %d ", url, res.StatusCode)
	}

	return res, nil
}

func (c *Client) Do(r *http.Request) (*http.Response, error) {
	return c.client.Do(r)
}

type validURLRoundTripper struct {
	defaultRoundTripper http.RoundTripper
}

func (v validURLRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {

	scheme := r.URL.Scheme
	isHTTP := scheme == "http" || scheme == "https"

	if !isHTTP {
		u := "http://" + r.URL.String()
		r.URL, _ = url.Parse(u)
	}

	return v.defaultRoundTripper.RoundTrip(r)

}

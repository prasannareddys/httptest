package api

import (
	"fmt"
	"io"

	"github.com/prasanna-eyewa/httptest/external"
)

type Api interface {
	CallURL() (*CallURLResponseBody, error)
}

type Client struct {
	reqUrl         string
	externalClient external.HttpClient
}

type CallURLResponseBody struct {
	ResponseBody, RequestURL string
}

func NewApiClient(reqUrl string, client external.HttpClient) *Client {
	return &Client{
		reqUrl:         reqUrl,
		externalClient: client,
	}
}

func (ac Client) CallURL() (*CallURLResponseBody, error) {
	res, err := ac.externalClient.Get(ac.reqUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to request url %s with error %s ", ac.reqUrl, err.Error())
	}
	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response of url %s with error %s", ac.reqUrl, err.Error())
	}

	return &CallURLResponseBody{ResponseBody: string(resBody), RequestURL: res.Request.Referer()}, nil
}

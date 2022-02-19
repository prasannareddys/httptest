package api

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestClient_CallURL(t *testing.T) {
	m := &externalClientMock{}
	npc := NewApiClient("test.com", m)
	_, err := npc.CallURL()
	if err != nil {
		t.Fail()
	}
}

type externalClientMock struct{}

func (e externalClientMock) Get(url string) (*http.Response, error) {
	responseBody := ioutil.NopCloser(bytes.NewReader([]byte("dummy")))
	return &http.Response{StatusCode: 200, Body: responseBody, Request: &http.Request{Header: map[string][]string{"refer": {"test.com"}}}}, nil
}

func TestClient_CallURLError(t *testing.T) {
	m := &externalClientFailMock{}
	npc := NewApiClient("test.com", m)
	_, err := npc.CallURL()
	if err == nil {
		t.Fail()
	}
}

type externalClientFailMock struct{}

func (e externalClientFailMock) Get(url string) (*http.Response, error) {
	return nil, fmt.Errorf("dummmy error")
}

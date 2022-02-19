package external

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_Get(t *testing.T) {
	expected := "dummy data"
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(expected))
	}))
	defer svr.Close()
	c := GetClient()
	res, err := c.Get(svr.URL)
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}
	defer res.Body.Close()
	resBody, _ := io.ReadAll(res.Body)
	sr := string(resBody)
	if string(resBody) != expected {
		t.Errorf("expected %s got %s", expected, sr)
	}
}

func TestClient_GetErrot(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer svr.Close()
	c := GetClient()
	_, err := c.Get(svr.URL)
	if err == nil {
		t.Errorf("expected err to be not nil")
	}
}

package testutil

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func NewTestServer(t *testing.T, hf func(w http.ResponseWriter, r *http.Request)) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(hf))
	t.Cleanup(ts.Close)
	return ts
}

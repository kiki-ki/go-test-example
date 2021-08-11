package testutil

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

type AssertResponseWant struct {
	StatusCode int
	Body       string
}

func AssertResponse(t *testing.T, got *http.Response, want AssertResponseWant, path string) {
	t.Helper()

	if want.StatusCode != got.StatusCode {
		t.Errorf("statusCode: want=%d, got=%d", want.StatusCode, got.StatusCode)
	}
	assertResponseBodyWithFile(t, got, path)
}

func assertResponseBodyWithFile(t *testing.T, res *http.Response, path string) {
	t.Helper()

	rs := GetStrFromFile(t, path)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
			t.Fatal(err)
	}
	var actual bytes.Buffer
	if err := json.Indent(&actual, body, "", "  "); err != nil {
			t.Fatal(err)
	}
	assert.JSONEq(t, rs, actual.String())
}

func GetStrFromFile(t *testing.T, path string) string {
	t.Helper()

	b, err := ioutil.ReadFile(path)
	if err != nil {
			t.Fatal(err)
	}
	return string(b)
}

type URLParam struct {
	Key, Val string
}

func NewRequestWithURLParams(t *testing.T, method string, target string, body io.Reader, params ...URLParam) *http.Request {
	t.Helper()

	r := httptest.NewRequest(method, target, body)
	return addURLParams(t, r, params...)
}

func addURLParams(t *testing.T, r *http.Request, params ...URLParam) *http.Request {
	t.Helper()

	ctx := chi.NewRouteContext()
	for _, p := range params {
		ctx.URLParams.Add(p.Key, p.Val)
	}
	newR := r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
	return newR
}

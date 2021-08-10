package handler_test

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/kiki-ki/go-test-example/internal/app/handler"
)

func TestHello_Hello(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/hello", nil)
	handler.NewHelloHandler().Hello(w, r)

	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)

	want := struct {
		status int
		body   string
	}{status: 200, body: "hello"}

	if err != nil {
		t.Errorf("cannot read Hello() response: %v", err)
	}
	if resp.StatusCode != want.status {
		t.Errorf("status: got = %d, want = %d", resp.StatusCode, want.status)
	}
	if string(body) != want.body {
		t.Errorf("body: got = %s, want = %s", string(body), want.body)
	}
}

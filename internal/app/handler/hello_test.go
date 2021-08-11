package handler_test

import (
	"net/http/httptest"
	"testing"

	"github.com/kiki-ki/go-test-example/internal/app/handler"
	"github.com/kiki-ki/go-test-example/internal/testutil"
)

func TestHelloHandler_Hello(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/hello", nil)
	want := testutil.AssertResponseWant{StatusCode: 200, Body: "hello"}

	handler.NewHelloHandler().Hello(w, r)

	res := w.Result()
	defer res.Body.Close()


	testutil.AssertResponse(t, res, want, "./testdata/hello/hello_res.golden")
}

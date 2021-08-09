package handler

import (
	"net/http"

	"github.com/go-chi/render"
)

func NewHelloHandler() HelloHandler {
	return &helloHandler{}
}

type HelloHandler interface {
	Hello(http.ResponseWriter, *http.Request)
}

type helloHandler struct{}

func (h *helloHandler) Hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, "hello")
}

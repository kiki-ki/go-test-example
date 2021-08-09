package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kiki-ki/go-test-example/internal/app/handler"
)

func main() {
	r := chi.NewRouter()
	setRouter(r)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("!! %+v", err)
	}
}

func setRouter(r chi.Router) {
	// logging
	r.Use(middleware.Logger)

	// routing
	hello := handler.NewHelloHandler()
	r.Route("/hello", func(r chi.Router) {
		r.Get("/", hello.Hello)
	})
}

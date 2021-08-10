package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kiki-ki/go-test-example/internal/app/handler"
	"github.com/kiki-ki/go-test-example/internal/interface/database"
)

func main() {
	sqlDB := database.NewSqlDB()
	db := database.NewDB(sqlDB)
	defer db.Close()

	r := chi.NewRouter()
	setRouter(r, db)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("!! %+v", err)
	}
}

func setRouter(r chi.Router, db database.DB) {
	// logging
	r.Use(middleware.Logger)

	// routing
	helloH := handler.NewHelloHandler()
	r.Route("/hello", func(r chi.Router) {
		r.Get("/", helloH.Hello)
	})

	userH := handler.NewUserHandler(db)
	r.Route("/users", func(r chi.Router) {
		r.Get("/", userH.Index)
		r.Get("/{userId}", userH.Show)
		r.Put("/{userId}", userH.Update)
		r.Post("/", userH.Create)
		r.Delete("/{userId}", userH.Delete)
	})
}

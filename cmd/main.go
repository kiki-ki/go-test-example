package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kiki-ki/go-test-example/internal/app/handler"
	"github.com/kiki-ki/go-test-example/internal/app/model"
)

func main() {
	db, err := sql.Open("mysql", "user:pass@tcp(localhost:3314)/gte_dev")
	if err != nil {
		log.Fatal(err)
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}
	defer dbmap.Db.Close()
	dbmap.AddTableWithName(model.User{}, "users")
	u := model.User{Name: "太郎", Email: "taro@chan.com", Age: 20}
	err = dbmap.Insert(&u)
	if err != nil {
		fmt.Println(err.Error())
	}

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

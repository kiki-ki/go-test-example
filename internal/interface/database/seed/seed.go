package main

import (
	"github.com/kiki-ki/go-test-example/internal/app/model"
	"github.com/kiki-ki/go-test-example/internal/interface/database"
)

func main() {
	sqlDB := database.NewSqlDB()
	db := database.NewDB(sqlDB)
	defer db.Close()

	conn := db.Conn()
	user1 := model.User{
		Name:  "user1",
		Email: "user1@exa.com",
		Age:   20,
	}
	user2 := model.User{
		Name:  "user2",
		Email: "user2@exa.com",
		Age:   20,
	}
	if err := conn.Insert(&user1, &user2); err != nil {
		panic(err.Error())
	}
}

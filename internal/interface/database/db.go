package database

import (
	"database/sql"
	"log"

	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kiki-ki/go-test-example/internal/app/model"
)

type DB interface {
	Connect() *gorp.DbMap
	Close()
}

func NewDB() DB {
	sqlDB, err := sql.Open("mysql", "user:pass@tcp(localhost:3314)/gte_dev")
	if err != nil {
		panic(err.Error())
	}
	dbmap := &gorp.DbMap{Db: sqlDB, Dialect: gorp.MySQLDialect{}}
	dbmap.TraceOn("[gorp]", &logger{})
	addTableSettings(dbmap)
	return &db{
		connection: dbmap,
	}
}

func addTableSettings(conn *gorp.DbMap) {
	conn.AddTableWithName(model.User{}, "users").SetKeys(true, "Id")
}

type db struct {
	connection *gorp.DbMap
}

func (db *db) Connect() *gorp.DbMap {
	return db.connection
}

func (db *db) Close() {
	err := db.connection.Db.Close()
	if err != nil {
		panic(err.Error())
	}
}

type logger struct{}

func (l *logger) Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

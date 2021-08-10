package database

import (
	"database/sql"
	"log"

	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kiki-ki/go-test-example/internal/app/model"
)

type DB interface {
	Conn() *gorp.DbMap
	Close() error
}

func NewDB(sqlDB *sql.DB) DB {
	dbmap := &gorp.DbMap{Db: sqlDB, Dialect: gorp.MySQLDialect{}}
	dbmap.TraceOn("[gorp]", &logger{})
	addTableSettings(dbmap)
	return &db{
		connection: dbmap,
	}
}

func NewSqlDB() *sql.DB {
	// TODO: env対応
	sqlDB, err := sql.Open("mysql", "user:pass@tcp(localhost:3314)/gte_dev")
	if err != nil {
		panic(err.Error())
	}
	return sqlDB
}

func addTableSettings(conn *gorp.DbMap) {
	conn.AddTableWithName(model.User{}, "users").SetKeys(true, "Id")
}

type db struct {
	connection *gorp.DbMap
}

func (db *db) Conn() *gorp.DbMap {
	return db.connection
}

func (db *db) Close() error {
	err := db.connection.Db.Close()
	if err != nil {
		return err
	}
	return nil
}

type logger struct{}

func (l *logger) Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

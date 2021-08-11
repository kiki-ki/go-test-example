package testutil

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/kiki-ki/go-test-example/internal/interface/database"
)

func NewMockDB(t *testing.T) (database.DB, sqlmock.Sqlmock) {
	t.Helper()

	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	db := database.NewDB(sqlDB)
	return db, mock
}

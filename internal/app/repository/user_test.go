package repository_test

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/kiki-ki/go-test-example/internal/app/model"
	"github.com/kiki-ki/go-test-example/internal/app/repository"
	"github.com/kiki-ki/go-test-example/internal/testutil"
)

func TestUserRepository_Find(t *testing.T) {
	db, mock := testutil.NewMockDB(t)
	defer db.Close()

	want := model.User{
		Id: 1,
		Name: "taro",
		Email: "taro@chan.com",
		Age: 5,
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "age"}).
		AddRow(want.Id, want.Name, want.Email, want.Age)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM users WHERE id = ?`)).
		WithArgs(want.Id).
		WillReturnRows(rows)

	got, err := repository.NewUserRepository().Find(want.Id, db.Conn())
	if err != nil {
		t.Error(err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
	if want != got {
		t.Errorf("want: %v, got: %v", want, got)
	}
}

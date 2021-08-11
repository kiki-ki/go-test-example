package handler_test

import (
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/kiki-ki/go-test-example/internal/app/handler"
	"github.com/kiki-ki/go-test-example/internal/app/model"
	"github.com/kiki-ki/go-test-example/internal/testutil"
)

func TestUserHandler_Show(t *testing.T) {
	w := httptest.NewRecorder()
	r := testutil.NewRequestWithURLParams(
		t, "GET", "/dummy", nil,
		testutil.URLParam{Key: "userId", Val: "1"},
	)
	want := testutil.AssertResponseWant{
		StatusCode: 200, Body: "a",
	}

	db, mock := testutil.NewMockDB(t)
	defer db.Close()

	u := model.User{
		Id:    1,
		Name:  "taro",
		Email: "taro@chan.com",
		Age:   5,
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "age"}).
		AddRow(u.Id, u.Name, u.Email, u.Age)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM users WHERE id = ?`)).
		WithArgs(u.Id).
		WillReturnRows(rows)

	h := handler.NewUserHandler(db)
	h.Show(w, r)

	res := w.Result()
	defer res.Body.Close()

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
	testutil.AssertResponse(t, res, want, "./testdata/user/show_res.golden")
}

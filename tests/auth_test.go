package tests

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"net/http"
	"net/http/httptest"
	"net/url"
	"regexp"
	"strings"
	"testing"
	"valorize-app/handlers"
)

func mockDb() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()

	if err != nil {
		panic(err.Error())
	}

	mockedDB, _ := gorm.Open("mysql", db)

	return mockedDB, mock
}

func TestLogin(t *testing.T) {
	// Setup
	form := url.Values{}
	form.Add("username", "javier123454321")
	form.Add("password", "test")
	req := httptest.NewRequest("POST", "/login", strings.NewReader(form.Encode()))
	req.Form = form
	rec := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(req, rec)

	fakeDb, mock := mockDb()
	mock.MatchExpectationsInOrder(false)
	mock.ExpectBegin()
	columns := []string{"id", "username", "email", "password", "deleted_at"}

	mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `users`")).
		WillReturnRows(sqlmock.NewRows(columns).
			AddRow("4", "javier123454321", "javier@example.com", "$2a$10$e.kv61mInrLwSwD5YUIoLerxI3FEzh5w.PAmXXrk3.RDuozLRoTv2", nil))

	auth := handlers.AuthHandler{DB: fakeDb}

	if assert.NoError(t, auth.Login(c), "/login fails") {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

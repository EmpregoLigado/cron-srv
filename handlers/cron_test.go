package handlers

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/rafaeljesus/kyp-users/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var KYP_USERS_DB = os.Getenv("KYP_USERS_DB")
var env *Env

func TestMain(m *testing.M) {
	db, err := models.NewDB(KYP_USERS_DB)
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{})
	env = &Env{db}
	code := m.Run()
	os.Exit(code)
}

func TestUsersCreate(t *testing.T) {
	response := `{"email":"foo@test.com", "password":"12345678"}`
	e := echo.New()
	req, _ := http.NewRequest(echo.POST, "/v1/users", strings.NewReader(response))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))

	if assert.NoError(t, env.UsersCreate(ctx)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func TestUsersShow(t *testing.T) {
	response := `{"id":"1"}`
	e := echo.New()
	req, _ := http.NewRequest(echo.GET, "/v1/users/1", strings.NewReader(response))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")

	if assert.NoError(t, env.UsersShow(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestUsersAuthenticate(t *testing.T) {
	response := `{"email":"foo@test.com", "password":"12345678"}`
	e := echo.New()
	req, _ := http.NewRequest(echo.POST, "/v1/users", strings.NewReader(response))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))

	if assert.NoError(t, env.UsersAuthenticate(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

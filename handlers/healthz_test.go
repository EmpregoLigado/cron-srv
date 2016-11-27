package handlers

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const (
	responseJSON = `{"alive":true}`
)

func TestHealthzIndex(t *testing.T) {
	e := echo.New()
	req, _ := http.NewRequest(echo.GET, "/v1/healthz", strings.NewReader(responseJSON))

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))
	env := Env{}

	if assert.NoError(t, env.HealthzIndex(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, responseJSON, rec.Body.String())
	}
}

package api

import (
	"encoding/json"
	"github.com/EmpregoLigado/cron-srv/mock"
	"github.com/EmpregoLigado/cron-srv/models"
	"github.com/nbari/violetear"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestEventsIndex(t *testing.T) {
	schedulerMock := mock.NewScheduler()
	repoMock := mock.NewRepo()
	h := NewAPIHandler(repoMock, schedulerMock)

	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/v1/events", nil)
	if err != nil {
		t.Errorf("Expected initialize request %s", err)
	}

	r := violetear.New()
	r.HandleFunc("/v1/events", h.EventsIndex, "GET")
	r.ServeHTTP(res, req)

	events := []models.Event{}
	if err := json.NewDecoder(res.Body).Decode(&events); err != nil {
		t.Errorf("Expected to decode response json %s", err)
	}

	if len(events) == 0 {
		t.Errorf("Expected response to not be empty %s", strconv.Itoa(len(events)))
	}

	if res.Code != http.StatusOK {
		t.Errorf("Expected status %d to be equal %d", res.Code, http.StatusOK)
	}
}

func TestEventsIndexByStatus(t *testing.T) {
	schedulerMock := mock.NewScheduler()
	repoMock := mock.NewRepo()
	h := NewAPIHandler(repoMock, schedulerMock)

	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/v1/events?status=active", nil)
	if err != nil {
		t.Errorf("Expected initialize request %s", err)
	}

	r := violetear.New()
	r.HandleFunc("/v1/events", h.EventsIndex, "GET")
	r.ServeHTTP(res, req)

	if !repoMock.ByStatus {
		t.Errorf("Expected to search by status")
	}

	if res.Code != http.StatusOK {
		t.Errorf("Expected status %d to be equal %d", res.Code, http.StatusOK)
	}
}

func TestEventsIndexByExpression(t *testing.T) {
	schedulerMock := mock.NewScheduler()
	repoMock := mock.NewRepo()
	h := NewAPIHandler(repoMock, schedulerMock)

	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/v1/events?expression=* * * * *", nil)
	if err != nil {
		t.Errorf("Expected initialize request %s", err)
	}

	r := violetear.New()
	r.HandleFunc("/v1/events", h.EventsIndex, "GET")
	r.ServeHTTP(res, req)

	if !repoMock.ByExpression {
		t.Errorf("Expected to search by expression")
	}

	if res.Code != http.StatusOK {
		t.Errorf("Expected status %d to be equal %d", res.Code, http.StatusOK)
	}
}

func TestEventCreate(t *testing.T) {
	schedulerMock := mock.NewScheduler()
	repoMock := mock.NewRepo()
	h := NewAPIHandler(repoMock, schedulerMock)

	res := httptest.NewRecorder()
	body := strings.NewReader(`{"url":"http://foo.com"}`)
	req, err := http.NewRequest("POST", "/v1/events", body)
	if err != nil {
		t.Errorf("Expected initialize request %s", err)
	}

	r := violetear.New()
	r.HandleFunc("/v1/events", h.EventsCreate, "POST")
	r.ServeHTTP(res, req)

	if !repoMock.Created {
		t.Error("Expected repo create to be called")
	}

	if !schedulerMock.Created {
		t.Error("Expected scheduler create to be called")
	}

	if res.Code != http.StatusCreated {
		t.Errorf("Expected status %d to be equal %d", res.Code, http.StatusOK)
	}
}

func TestEventShow(t *testing.T) {
	schedulerMock := mock.NewScheduler()
	repoMock := mock.NewRepo()
	h := NewAPIHandler(repoMock, schedulerMock)

	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/v1/events/1", nil)
	if err != nil {
		t.Errorf("Expected initialize request %s", err)
	}

	r := violetear.New()
	r.AddRegex(":id", `^\d+$`)
	r.HandleFunc("/v1/events/:id", h.EventsShow, "GET")
	r.ServeHTTP(res, req)

	if !repoMock.Found {
		t.Error("Expected repo findEventById to be called")
	}

	if res.Code != http.StatusOK {
		t.Errorf("Expected status %d to be equal %d", res.Code, http.StatusOK)
	}
}

func TestEventsUpdate(t *testing.T) {
	schedulerMock := mock.NewScheduler()
	repoMock := mock.NewRepo()
	h := NewAPIHandler(repoMock, schedulerMock)

	res := httptest.NewRecorder()
	body := strings.NewReader(`{"url":"http://foo.com"}`)
	req, err := http.NewRequest("PUT", "/v1/events/1", body)
	if err != nil {
		t.Errorf("Expected initialize request %s", err)
	}

	r := violetear.New()
	r.AddRegex(":id", `^\d+$`)
	r.HandleFunc("/v1/events/:id", h.EventsUpdate, "PUT")
	r.ServeHTTP(res, req)

	if !repoMock.Updated {
		t.Error("Expected repo update to be called")
	}

	if !schedulerMock.Updated {
		t.Error("Expected scheduler update to be called")
	}

	if res.Code != http.StatusOK {
		t.Errorf("Expected status %d to be equal %d", res.Code, http.StatusOK)
	}
}

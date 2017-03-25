package api

import (
	"encoding/json"
	"github.com/nbari/violetear"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthzIndex(t *testing.T) {
	h := new(APIHandler)

	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/v1/healthz", nil)
	if err != nil {
		t.Errorf("Expected initialize request %s", err)
	}

	r := violetear.New()
	r.HandleFunc("/v1/healthz", h.HealthzIndex, "GET")
	r.ServeHTTP(res, req)

	response := make(map[string]string)
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		t.Errorf("Expected to decode response json %s", err)
	}

	if response["status"] != "up" {
		t.Errorf("Expected status to equal %s", response["status"])
	}

	if res.Code != http.StatusOK {
		t.Error("Expected status %s to be equal %s", res.Code, http.StatusOK)
	}
}

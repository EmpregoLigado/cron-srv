package api

import (
	"encoding/json"
	"github.com/EmpregoLigado/cron-srv/repo"
	"github.com/EmpregoLigado/cron-srv/scheduler"
	"net/http"
)

type APIHandler struct {
	Repo      repo.Repo
	Scheduler scheduler.Scheduler
}

func NewAPIHandler(r repo.Repo, s scheduler.Scheduler) *APIHandler {
	return &APIHandler{
		Repo:      r,
		Scheduler: s,
	}
}

func JSON(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if v != nil || code == http.StatusNoContent {
		json.NewEncoder(w).Encode(v)
	}
}

package api

import (
	"github.com/EmpregoLigado/cron-srv/repo"
	"github.com/EmpregoLigado/cron-srv/scheduler"
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

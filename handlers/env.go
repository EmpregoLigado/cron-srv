package handlers

import (
	"github.com/EmpregoLigado/cron-srv/models"
)

type Env struct {
	Repo  models.Repo
	Sched models.Sched
}

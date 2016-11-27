package handlers

import (
	"github.com/rafaeljesus/cron-srv/models"
)

type Env struct {
	Repo  models.Repo
	Sched models.Sched
}

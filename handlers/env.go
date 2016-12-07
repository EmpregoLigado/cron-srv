package handlers

import (
	"github.com/rafaeljesus/cron-srv/models"
)

type Env struct {
	Repo  models.Repo
	Sched models.Sched
}

func (env *Env) ScheduleAll() error {
	crons := []models.Cron{}
	query := models.Query{}
	if err := env.Repo.Search(&query, &crons); err != nil {
		return err
	}

	for _, cron := range crons {
		if err := env.Sched.Create(&cron); err != nil {
			return err
		}
	}

	return nil
}

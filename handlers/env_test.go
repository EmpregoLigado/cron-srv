package handlers

import (
	"github.com/rafaeljesus/cron-srv/models"
	"testing"
)

func TestScheduleAll(t *testing.T) {
	env := Env{&models.RepoMock{}, &models.SchedMock{}}
	if err := env.ScheduleAll(); err != nil {
		t.Error("Failed to schedule all crons", err)
	}
}

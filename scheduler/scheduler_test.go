package scheduler

import (
	"github.com/EmpregoLigado/cron-srv/models"
	"github.com/EmpregoLigado/cron-srv/repo"
	"testing"
)

func TestScheduleAll(t *testing.T) {
	repoMock := repo.NewMock()
	s := New()
	if err := s.ScheduleAll(repoMock); err != nil {
		t.Errorf("Expected to schedule all events %s", err)
	}
}

func TestSchedulerCreate(t *testing.T) {
	s := New()
	c := &models.Cron{Id: 1, Expression: "* * * * *"}
	if err := s.Create(c); err != nil {
		t.Errorf("Expected to schedule a cron %s", err)
	}
}

func TestSchedulerFind(t *testing.T) {
	s := New()
	c := &models.Cron{Id: 1, Expression: "* * * * *"}
	if err := s.Create(c); err != nil {
		t.Errorf("Expected to schedule a cron %s", err)
	}

	_, err := s.Find(c.Id)
	if err != nil {
		t.Errorf("Expected to find a cron %s", err)
	}
}

func TestSchedulerUpdate(t *testing.T) {
	s := New()
	c := &models.Cron{Id: 1, Expression: "* * * * *"}
	if err := s.Create(c); err != nil {
		t.Errorf("Expected to schedule a cron %s", err)
	}

	c.Status = "active"
	if err := s.Update(c); err != nil {
		t.Errorf("Expected to update a scheduled cron %s", err)
	}
}

func TestSchedulerDelete(t *testing.T) {
	s := New()
	c := &models.Cron{Id: 1, Expression: "* * * * *"}
	if err := s.Create(c); err != nil {
		t.Errorf("Expected to schedule a cron %s", err)
	}

	if err := s.Delete(c.Id); err != nil {
		t.Errorf("Expected to delete a scheduled cron %s", err)
	}
}

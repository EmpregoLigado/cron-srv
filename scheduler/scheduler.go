package scheduler

import (
	"errors"
	"github.com/EmpregoLigado/cron-srv/models"
	"github.com/EmpregoLigado/cron-srv/repo"
	"github.com/EmpregoLigado/cron-srv/runner"
	"github.com/robfig/cron"
	"sync"
)

var (
	ErrEventNotExist = errors.New("finding a scheduled event requires a existent cron id")
)

type Scheduler interface {
	Create(cron *models.Event) (err error)
	Update(cron *models.Event) (err error)
	Delete(id uint) (err error)
	Find(id uint) (cron *cron.Cron, err error)
	ScheduleAll(repo repo.Repo) (err error)
}

type scheduler struct {
	sync.RWMutex
	Kv   map[uint]*cron.Cron
	Cron *cron.Cron
}

func New() Scheduler {
	s := &scheduler{
		Kv:   make(map[uint]*cron.Cron),
		Cron: cron.New(),
	}

	s.Cron.Start()

	return s
}

func (s *scheduler) ScheduleAll(repo repo.Repo) (err error) {
	crons := []models.Event{}
	query := new(models.Query)
	if err = repo.FindEvents(&crons, query); err != nil {
		return
	}

	for _, cron := range crons {
		if err = s.Create(&cron); err != nil {
			return
		}
	}

	return
}

func (s *scheduler) Create(event *models.Event) (err error) {
	s.Cron.AddFunc(event.Expression, func() {
		rc := &runner.Config{
			Url:     event.Url,
			Retries: event.Retries,
			Timeout: event.Timeout,
		}

		r := runner.New()
		r.Run() <- rc
	})

	s.Lock()
	defer s.Unlock()

	s.Kv[event.Id] = s.Cron

	return
}

func (s *scheduler) Find(id uint) (cron *cron.Cron, err error) {
	s.Lock()
	defer s.Unlock()

	cron, found := s.Kv[id]
	if !found {
		err = ErrEventNotExist
		return
	}

	return
}

func (s *scheduler) Update(cron *models.Event) (err error) {
	if err = s.Delete(cron.Id); err != nil {
		return
	}

	return s.Create(cron)
}

func (s scheduler) Delete(id uint) (err error) {
	s.Lock()
	defer s.Unlock()

	_, found := s.Kv[id]
	if !found {
		err = ErrEventNotExist
		return
	}

	s.Kv[id].Stop()
	s.Kv[id] = nil

	return
}

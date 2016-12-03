package models

import (
	"github.com/robfig/cron"
	"github.com/sethgrid/pester"
	"log"
)

type Sched interface {
	Create(cron *Cron) error
	Update(cron *Cron) error
	Delete(id uint) error
}

type Scheduler struct {
	Kv   map[uint]*cron.Cron
	Cron *cron.Cron
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		Kv:   make(map[uint]*cron.Cron),
		Cron: cron.New(),
	}
}

func (s *Scheduler) Create(cron *Cron) error {
	runJob := func() {
		client := pester.New()
		client.Concurrency = 4
		client.MaxRetries = cron.MaxRetries
		client.Backoff = pester.ExponentialBackoff

		_, err := client.Get(cron.Url)
		if err != nil {
			log.Println("Failed to run cron job at %s", cron.Url)
			return
		}

		log.Println("Cron job sent to url %s", cron.Url)
	}

	s.Cron.AddFunc(cron.Expression, runJob)
	s.Kv[cron.Id] = s.Cron

	return nil
}

func (s *Scheduler) Update(cron *Cron) error {
	if err := s.Delete(cron.Id); err != nil {
		return err
	}

	return s.Create(cron)
}

func (s *Scheduler) Delete(id uint) error {
	s.Kv[id].Stop()
	s.Kv[id] = nil
	return nil
}

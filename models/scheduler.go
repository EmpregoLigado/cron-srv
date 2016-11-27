package models

import (
	"github.com/robfig/cron"
	"github.com/sethgrid/pester"
	"log"
)

type Sched interface {
	Create(cron *Cron) error
}

type Scheduler struct {
	Kv   map[int]interface{}
	Cron *cron.Cron
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		Kv:   make(map[int]interface{}),
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

	return nil
}

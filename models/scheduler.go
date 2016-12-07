package models

import (
	"fmt"
	"github.com/robfig/cron"
	"net/http"
	"strconv"
	"time"
)

type Sched interface {
	Create(cron *Cron) error
	Update(cron *Cron) error
	Delete(id uint) error
}

type retriable func(retriable, int)

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

func (s *Scheduler) Start() {
	s.Cron.Start()
}

func (s *Scheduler) Create(cron *Cron) error {
	runJob := func(fn retriable, retries int) {
		_, err := http.Get(cron.Url)
		if err != nil {
			fmt.Printf("Retrying request to %s\nRetry count %s\n", cron.Url, strconv.Itoa(retries))

			if retries == 0 {
				fmt.Printf("Max retries reached %s\nFailed to send job to %s\n", retries, cron.Url)
				return
			}

			secs := time.Duration(cron.RetryTimeout) * time.Second
			time.Sleep(secs)

			fn(fn, retries-1)
		} else {
			fmt.Printf("Cron job sent to %s\n", cron.Url)
		}
	}

	s.Cron.AddFunc(cron.Expression, func() {
		runJob(runJob, cron.MaxRetries)
	})

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

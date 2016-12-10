package models

import (
	log "github.com/Sirupsen/logrus"
	"github.com/robfig/cron"
	"net/http"
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
			l := log.WithFields(log.Fields{
				"url":     cron.Url,
				"retries": retries,
			})

			l.Info("Retrying to send event")

			if retries == 0 {
				l.Info("Max retries reached")
				return
			}

			secs := time.Duration(cron.RetryTimeout) * time.Second
			time.Sleep(secs)

			fn(fn, retries-1)
		} else {
			log.WithFields(log.Fields{
				"url": cron.Url,
			}).Info("Cron job event sent")
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

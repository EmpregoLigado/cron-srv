package scheduler

import (
	"github.com/EmpregoLigado/cron-srv/models"
	"github.com/EmpregoLigado/cron-srv/repo"
	log "github.com/Sirupsen/logrus"
	"github.com/robfig/cron"
	"net/http"
	"sync"
	"time"
)

type Scheduler interface {
	Create(cron *models.Cron) (err error)
	Update(cron *models.Cron) (err error)
	Delete(id uint) (err error)
	ScheduleAll(repo repo.Repo) (err error)
}

type retriable func(retriable, int)

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
	crons := []models.Cron{}
	query := new(models.Query)
	if err = repo.Search(query, &crons); err != nil {
		return
	}

	for _, cron := range crons {
		if err = s.Create(&cron); err != nil {
			return
		}
	}

	return
}

func (s *scheduler) Create(cron *models.Cron) (err error) {
	runJob := func(fn retriable, retries int) {
		_, err := http.Get(cron.Url)
		if err == nil {
			log.WithField("url", cron.Url).Info("Cron job event sent")
			return
		}

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
	}

	s.Cron.AddFunc(cron.Expression, func() {
		runJob(runJob, cron.MaxRetries)
	})

	s.Lock()
	defer s.Unlock()

	s.Kv[cron.Id] = s.Cron

	return
}

func (s *scheduler) Update(cron *models.Cron) (err error) {
	if err = s.Delete(cron.Id); err != nil {
		return
	}

	return s.Create(cron)
}

func (s scheduler) Delete(id uint) (err error) {
	s.Lock()
	defer s.Unlock()

	s.Kv[id].Stop()
	s.Kv[id] = nil

	return
}

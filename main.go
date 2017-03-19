package main

import (
	"github.com/EmpregoLigado/cron-srv/api"
	"github.com/EmpregoLigado/cron-srv/conf"
	"github.com/EmpregoLigado/cron-srv/models"
	"github.com/EmpregoLigado/cron-srv/scheduler"
	log "github.com/Sirupsen/logrus"
	"github.com/nbari/violetear"
	"net/http"
)

func main() {
	db, err := models.NewDB(models.DBConfig{
		Url: conf.CRON_SRV_DB,
	})

	if err != nil {
		log.WithError(err).Fatal("Failed to init database connection!")
		return
	}
	defer db.Close()
	event := new(models.Event)
	db.AutoMigrate(event)

	sc := scheduler.New()
	go func() {
		if err := sc.ScheduleAll(db); err != nil {
			log.WithError(err).Fatal("Failed to schedule crons from database!")
		}
	}()

	h := api.NewAPIHandler(db, sc)

	router := violetear.New()
	router.LogRequests = true
	router.RequestID = "X-Request-ID"
	router.AddRegex(":id", `^\d+$`)
	router.HandleFunc("/v1/healthz", h.HealthzIndex, "GET")
	router.HandleFunc("/v1/events", h.EventsIndex, "GET")
	router.HandleFunc("/v1/events", h.EventsCreate, "POST")
	router.HandleFunc("/v1/events/:id", h.EventsShow, "GET")
	router.HandleFunc("/v1/events/:id", h.EventsUpdate, "PUT")
	router.HandleFunc("/v1/events/:id", h.EventsDelete, "DELETE")

	log.WithField("port", conf.CRON_SRV_PORT).Info("Starting Cron Service")
	log.Fatal(http.ListenAndServe(":"+conf.CRON_SRV_PORT, router))
}

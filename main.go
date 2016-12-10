package main

import (
	"github.com/EmpregoLigado/cron-srv/handlers"
	"github.com/EmpregoLigado/cron-srv/models"
	log "github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/middleware"
	"os"
)

var CRON_SRV_PORT = os.Getenv("CRON_SRV_PORT")
var CRON_SRV_DB = os.Getenv("CRON_SRV_DB")

func main() {
	db, err := models.NewDB(CRON_SRV_DB)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("Failed to init database connection!")
	}

	db.AutoMigrate(&models.Cron{})
	sc := models.NewScheduler()
	sc.Start()

	env := &handlers.Env{db, sc}

	go func() {
		if err := env.ScheduleAll(); err != nil {
			log.Panic(err)
		}
	}()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	v1 := e.Group("/v1")
	v1.GET("/healthz", env.HealthzIndex)

	v1.GET("/crons", env.CronIndex)
	v1.POST("/cron", env.CronCreate)
	v1.GET("/cron/:id", env.CronShow)
	v1.PUT("/cron/:id", env.CronUpdate)
	v1.DELETE("/cron/:id", env.CronDelete)

	log.WithFields(log.Fields{
		"port": CRON_SRV_PORT,
	}).Info("Starting Cron Service")

	e.Run(fasthttp.New(":" + CRON_SRV_PORT))
}

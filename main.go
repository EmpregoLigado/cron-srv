package main

import (
	"github.com/EmpregoLigado/cron-srv/handlers"
	"github.com/EmpregoLigado/cron-srv/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/middleware"
	"log"
	"os"
)

var CRON_SRV_PORT = os.Getenv("CRON_SRV_PORT")
var CRON_SRV_DB = os.Getenv("CRON_SRV_DB")

func main() {
	db, err := models.NewDB(CRON_SRV_DB)
	if err != nil {
		log.Panic(err)
	}

	db.AutoMigrate(&models.Cron{})
	sc := models.NewScheduler()
	sc.Start()

	env := &handlers.Env{db, sc}

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

	log.Print("Starting Cron Service at port ", CRON_SRV_PORT)

	e.Run(fasthttp.New(":" + CRON_SRV_PORT))
}

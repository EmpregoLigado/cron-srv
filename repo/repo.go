package repo

import "github.com/EmpregoLigado/cron-srv/models"

type Repo interface {
	CreateEvent(event *models.Event) (err error)
	FindEventById(event *models.Event, id int) (err error)
	UpdateEvent(event *models.Event) (err error)
	DeleteEvent(event *models.Event) (err error)
	FindEvents(events *[]models.Event, query *models.Query) (err error)
}

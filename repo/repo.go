package repo

import "github.com/EmpregoLigado/cron-srv/models"

type Repo interface {
	CreateCron(cron *models.Cron) (err error)
	FindCronById(cron *models.Cron, id int) (err error)
	UpdateCron(cron *models.Cron) (err error)
	DeleteCron(cron *models.Cron) (err error)
	Search(query *models.Query, crons *[]models.Cron) (err error)
}

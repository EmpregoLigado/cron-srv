package models

type Repo interface {
	CreateCron(*Cron) error
	FindCronById(*Cron, int) error
	UpdateCron(*Cron) error
	DeleteCron(*Cron) error
	Search(*Query, *[]Cron) error
}

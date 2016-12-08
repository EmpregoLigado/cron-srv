package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Cron struct {
	Id           uint       `json:"id",sql:"primary_key"`
	Url          string     `json:"url",sql:"not null"`
	Expression   string     `json:"expression",sql:"not null"`
	Status       string     `json:"status",sql:"not null"`
	MaxRetries   int        `json:"max_retries, sql:"DEFAULT": 1`
	RetryTimeout int        `json:"retry_timeout, sql:"DEFAULT": 5`
	CreatedAt    time.Time  `json:"created_at",sql:"not null"`
	UpdatedAt    time.Time  `json:"updated_at",sql:"not null"`
	DeletedAt    *time.Time `json:"created_at, omitempty"`
}

func (repo *DB) CreateCron(c *Cron) error {
	return repo.Create(c).Error
}

func (repo *DB) FindCronById(c *Cron, id int) error {
	return repo.Find(c, id).Error
}

func (repo *DB) UpdateCron(c *Cron) error {
	return repo.Save(c).Error
}

func (repo *DB) DeleteCron(c *Cron) error {
	return repo.Delete(c).Error
}

func (repo *DB) Search(q *Query, crons *[]Cron) error {
	if q.IsEmpty() {
		return repo.Find(crons).Error
	}

	var r *gorm.DB
	if q.Status != "" {
		r = repo.Where("status = ?", q.Status)
	}

	if q.Expression != "" {
		r = repo.Where("expression = ?", q.Expression)
	}

	return r.Find(crons).Error
}

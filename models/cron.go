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

func (db *DB) CreateCron(c *Cron) error {
	return db.Create(c).Error
}

func (db *DB) FindCronById(c *Cron, id int) error {
	return db.Find(c, id).Error
}

func (db *DB) UpdateCron(c *Cron) error {
	return db.Save(c).Error
}

func (db *DB) DeleteCron(c *Cron) error {
	return db.Delete(c).Error
}

func (db *DB) Search(q *Query, crons *[]Cron) error {
	if q.IsEmpty() {
		return db.Find(crons).Error
	}

	var r *gorm.DB
	if q.Status != "" {
		r = db.Where("status = ?", q.Status)
	}

	if q.Expression != "" {
		r = db.Where("expression = ?", q.Expression)
	}

	return r.Find(crons).Error
}

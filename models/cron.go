package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Cron struct {
	Id         uint       `json:"id", sql:"primary_key"`
	Url        string     `json:"url", sql:"not null"`
	Expression string     `json:"expression", sql:"not null"`
	Status     string     `json:"status", sql:"not null"`
	CreatedAt  time.Time  `json:"created_at", sql:"not null"`
	UpdatedAt  time.Time  `json:"updated_at", sql:"not null`
	DeletedAt  *time.Time `json:"created_at", omitempty"`
}

func (repo *DB) Create(u *Cron) *gorm.DB {
	return repo.Create(u)
}

func (repo *DB) FindById(c *Cron, id int) *gorm.DB {
	return repo.Find(c, id)
}

func (repo *DB) Update(u *Cron) *gorm.DB {
	return repo.Save(u)
}

func (repo *DB) Delete(u *Cron) *gorm.DB {
	return repo.Delete(u)
}

func (repo *DB) Search(q *Query, crons *[]Cron) *gorm.DB {
	if q.IsEmpty() {
		return repo.Find(crons)
	}

	var r *gorm.DB
	if q.Status != "" {
		r = repo.Where("status = ?", q.Status)
	}

	if q.Expression != "" {
		r = repo.Where("expression = ?", q.Expression)
	}

	return r.Find(crons)
}

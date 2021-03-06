package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Event struct {
	Id         uint       `json:"id",sql:"primary_key"`
	Url        string     `json:"url",sql:"not null"`
	Expression string     `json:"expression",sql:"not null"`
	Status     string     `json:"status",sql:"not null"`
	Retries    int64      `json:"retries`
	Timeout    int64      `json:"timeout`
	CreatedAt  time.Time  `json:"created_at",sql:"not null"`
	UpdatedAt  time.Time  `json:"updated_at",sql:"not null"`
	DeletedAt  *time.Time `json:"created_at,omitempty"`
}

func (db *DB) CreateEvent(c *Event) error {
	return db.Create(c).Error
}

func (db *DB) FindEventById(c *Event, id int) error {
	return db.Find(c, id).Error
}

func (db *DB) UpdateEvent(c *Event) error {
	return db.Save(c).Error
}

func (db *DB) DeleteEvent(c *Event) error {
	return db.Delete(c).Error
}

func (db *DB) FindEvents(events *[]Event, q *Query) error {
	if q.IsEmpty() {
		return db.Find(events).Error
	}

	var r *gorm.DB
	if q.Status != "" {
		r = db.Where("status = ?", q.Status)
	}

	if q.Expression != "" {
		r = db.Where("expression = ?", q.Expression)
	}

	return r.Find(events).Error
}

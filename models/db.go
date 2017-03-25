package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DB struct {
	*gorm.DB
}

type DBConfig struct {
	Url         string
	MaxIdleConn int
	MaxOpenConn int
	LogMode     bool
}

func NewDB(c DBConfig) (db *DB, err error) {
	conn, err := gorm.Open("postgres", c.Url)
	if err != nil {
		return
	}

	if err = conn.DB().Ping(); err != nil {
		return
	}

	if c.MaxIdleConn == 0 {
		c.MaxIdleConn = 10
	}

	if c.MaxIdleConn == 0 {
		c.MaxIdleConn = 100
	}

	conn.DB().SetMaxIdleConns(c.MaxIdleConn)
	conn.DB().SetMaxOpenConns(c.MaxOpenConn)
	conn.LogMode(c.LogMode)

	db = &DB{conn}

	return
}

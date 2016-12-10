package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DB struct {
	*gorm.DB
}

func NewDB(dbname string) (*DB, error) {
	conn, err := gorm.Open("postgres", dbname)
	if err != nil {
		return nil, err
	}

	if err := conn.DB().Ping(); err != nil {
		return nil, err
	}

	conn.DB().SetMaxIdleConns(10)
	conn.DB().SetMaxOpenConns(100)

	return &DB{conn}, nil
}

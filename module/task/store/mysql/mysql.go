package mysql

import (
	"gorm.io/gorm"
)

type taskSQL struct {
	db *gorm.DB
}

func NewTaskSQL (db *gorm.DB ) *taskSQL {
	return &taskSQL{db: db}
}
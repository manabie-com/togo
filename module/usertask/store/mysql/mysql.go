package mysql

import (
	"gorm.io/gorm"
)

type userTaskSQL struct {
	db *gorm.DB
}

func NewUserTaskSQL (db *gorm.DB ) *userTaskSQL {
	return &userTaskSQL{db: db}
}
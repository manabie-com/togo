package mysql

import (
	"gorm.io/gorm"
)

type userConfigSQL struct {
	db *gorm.DB
}

func NewUserConfigSQL (db *gorm.DB ) *userConfigSQL {
	return &userConfigSQL{db: db}
}
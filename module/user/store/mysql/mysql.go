package mysql

import (
	"gorm.io/gorm"
)

type userSQL struct {
	db *gorm.DB
}

func NewUserSQL(db *gorm.DB) *userSQL {
	return &userSQL{db: db}
}

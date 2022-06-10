package mysql

import (
	"gorm.io/gorm"
)

type userMongo struct {
	db *gorm.DB
}

func NewUserMongo(db *gorm.DB) *userMongo {
	return &userMongo{db: db}
}

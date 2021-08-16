package storage

import (
	"gorm.io/gorm"
)

type authStorage struct {
	db *gorm.DB
}

func NewAuthStorage(db *gorm.DB) *authStorage {
	return &authStorage{
		db: db,
	}
}
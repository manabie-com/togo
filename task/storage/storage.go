package storage

import "gorm.io/gorm"

type taskStorage struct {
	db *gorm.DB
}

func NewTaskStorage(db *gorm.DB) *taskStorage {
	return &taskStorage{
		db: db,
	}
}

package model

import (
	"errors"

	"github.com/jinzhu/gorm"
)

var (
	ErrSomeTasksAreNotSatisfying = errors.New("some tasks are not satisfying")
	ErrExceedingTaskLimit        = errors.New("exceeding task limit")
)

func where(query interface{}, args ...interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(query, args...)
	}
}

func preload(column string, conditions ...interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload(column, conditions...)
	}
}

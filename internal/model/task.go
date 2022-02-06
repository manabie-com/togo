package model

import "gorm.io/gorm"

type Task struct {
	*gorm.Model
	UserID int
	Name   string
}

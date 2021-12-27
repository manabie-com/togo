package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Content     string
	UserID      int
	CreatedDate string
}

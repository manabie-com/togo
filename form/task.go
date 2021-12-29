package form

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Content     string `json:"content"`
}


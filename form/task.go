package form

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Content     string `json:"content"`
	UserID      string `json:"user_id"`
	CreatedDate string `json:"created_date"`
}


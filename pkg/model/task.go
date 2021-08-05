package model

import "time"

// Task reflects tasks in DB
type Task struct {
	ID          int       `json:"id" gorm:"column=id;PRIMARY"`
	Content     string    `json:"content" gorm:"column=content"`
	UserID      int       `json:"user_id" gorm:"column=user_id"`
	CreatedDate time.Time `json:"created_date" gorm:"column=created_date"`
}

func (m *Task) TableName() string {
	return `tasks`
}

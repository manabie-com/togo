package entities

import "time"

type Task struct {
	ID          uint64    `gorm:"primary_key:auto_increment" json:"id"`
	UserId      uint64    `json:"userId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsCompleted bool      `json:"isCompleted"`
	CreatedAt   time.Time `json:"createdAt"`
}

func (b *Task) TableName() string {
	return "tasks"
}

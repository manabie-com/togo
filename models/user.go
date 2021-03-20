package models

import "time"

type User struct {
	Username  string    `gorm:"primary_key" json:"username"`
	Password  string    `gorm:"type:varchar(100);not null;" json:"password"`
	MaxTodo   int64     `gorm:"int64;default:3" json:"max_todo"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

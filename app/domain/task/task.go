package task

import (
	"time"

	"github.com/ansidev/togo/domain/user"
)

type Task struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title" gorm:"type:varchar(255);not null"`
	UserID    int64     `json:"user_id"`
	User      user.User `json:"-"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (m Task) TableName() string {
	return "task"
}

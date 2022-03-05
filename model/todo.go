package model

import (
	"time"

	"gorm.io/gorm"

	"github.com/khangjig/togo/codetype"
)

type Todo struct {
	ID        int64               `json:"id"`
	UserID    int64               `json:"user_id"`
	Title     string              `json:"title"`
	Content   string              `json:"content"`
	Status    codetype.TodoStatus `json:"status"`
	EditedAt  *time.Time          `json:"edited_at"`
	CreatedAt time.Time           `json:"created_at"`
	UpdatedAt time.Time           `json:"updated_at"`
	DeletedAt *gorm.DeletedAt     `json:"-"`
}

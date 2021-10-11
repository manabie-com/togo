package storages

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Task reflects tasks in DB
type Task struct {
	ID        uuid.UUID `gorm:"primaryKey;type=uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Content     string `json:"content"`
	UserID      string `json:"user_id"`
	CreatedDate string `json:"created_date"`
}

// User reflects users data from DB
type User struct {
	ID       string `json:"user_id" gorm:"primaryKey"`
	Password string `json:"password"`
	MaxTodo  int    `json:"max_todo"`
}

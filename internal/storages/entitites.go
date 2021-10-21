package storages

import (
	"database/sql"
	"time"
)

//Models is the wrapper for database
type Models struct {
	DB DBModel
}

//NewModels returns models with db pool
func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

// Task reflects tasks in DB
type Task struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// User reflects users data from DB
type User struct {
	ID        int
	Email     string
	Username  string
	Password  string
	MaxTodo   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

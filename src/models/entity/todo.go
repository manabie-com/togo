package entity

import "time"

type Status int

const (
	KStatusUncheck Status = 0
	KStatusCheck   Status = 1
)

type Todo struct {
	ID        int       `json:"id" db:"id"`
	UserID    int       `json:"user_id" db:"user_id"`
	Content   string    `json:"content" db:"content"`
	Status    Status    `json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type UserTodoConfig struct {
	UserID    int       `json:"user_id" db:"user_id"`
	MaxTodo   int       `json:"max_todo" db:"max_todo"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

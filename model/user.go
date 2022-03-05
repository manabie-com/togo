package model

import (
	"time"

	"github.com/khangjig/togo/codetype"
)

type User struct {
	ID        int64           `json:"id"`
	Email     string          `json:"email"`
	Name      string          `json:"name"`
	Gender    codetype.Gender `json:"gender"`
	Password  string          `json:"-"`
	MaxTodo   int             `json:"-"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

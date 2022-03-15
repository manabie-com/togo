package domain

import (
	"time"
)

type TodoLimit struct {
	Id        string    `json:"id"`
	UserId    string    `json:"user_id"`
	Limit     int       `json:"limit"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

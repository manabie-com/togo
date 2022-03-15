package domain

import (
	"time"
)

type Todo struct {
	Id          string    `json:"trans_id"`
	UserId      string    `json:"user_id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

package domain

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	tableName      struct{}  `pg:"users"`
	ID             uuid.UUID `json:"id"`
	DailyTaskLimit int       `json:"dailyTaskLimit"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

func NewUser(id uuid.UUID, dailyTaskLimit int) *User {
	return &User{
		ID:             id,
		DailyTaskLimit: dailyTaskLimit,
		UpdatedAt:      time.Now(),
	}
}

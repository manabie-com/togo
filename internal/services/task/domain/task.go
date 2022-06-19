package domain

import (
	"github.com/google/uuid"
	"time"
)

type Task struct {
	tableName   struct{}  `pg:"tasks"`
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"userId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsCompleted bool      `json:"isCompleted"`
	DueDate     time.Time `json:"dueDate"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func NewTask(id, userID uuid.UUID, title, description string, dueDate time.Time) *Task {
	return &Task{
		ID:          id,
		UserID:      userID,
		Title:       title,
		Description: description,
		DueDate:     dueDate,
		UpdatedAt:   time.Now(),
	}
}

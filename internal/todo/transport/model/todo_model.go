package model

import (
	"manabieAssignment/internal/core/entity"
	"time"
)

type TodoModel struct {
	UserID  int64
	Name    string
	Content string
}

func (t TodoModel) ToDomainModel() entity.Todo {
	return entity.Todo{
		UserID:    t.UserID,
		Name:      t.Name,
		Content:   t.Content,
		CreatedAt: time.Now(),
	}
}

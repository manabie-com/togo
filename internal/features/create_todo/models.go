package createtodo

import (
	"time"

	"github.com/google/uuid"
)

// Todo represents information about a todo.
type Todo struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	UserID      uuid.UUID `json:"userID"`
	DateCreated time.Time `json:"dateCreated"`
	DateUpdated time.Time `json:"dateUpdated"`
}

// NewTodo contains information needed to create a new todo.
type NewTodo struct {
	Title   string
	Content string
	UserID  uuid.UUID
}

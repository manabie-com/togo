package entities

import "time"

// todos table
type Todo struct {
	ID        int
	UserID    int
	Task      string
	DueDate   *time.Time
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

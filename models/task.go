package models

import "time"

type Task struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name" validate:"required"`
	Content   string    `json:"content"`
	Status    Status    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UserId    int32     `json:"userId"`
}

type Status string

const (
	toDo  Status = "todo"
	doing Status = "doing"
	done  Status = "done"
)

func (e Status) isValidStatus() bool {
	switch e {
	case toDo, doing, done:
		return true
	}
	return false
}

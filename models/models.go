package models

import "time"

type User struct {
	ID         int32   `json:"id"`
	Name       string  `json:"name"`
	Password   string  `json:"password"`
	IsPayment  bool    `json:"isPayment"`
	LimitTasks int     `json:"limitTasks"`
	Tasks      *[]Task `json:"tasks"`
}

type Task struct {
	ID           int32           `json:"id"`
	Name         string          `json:"name"`
	Content      string          `json:"content"`
	Status       Status          `json:"status"`
	CreatedAt    time.Time       `json:"createdAt"`
	UserId       int32           `json:"userId"`
	ChildrenTask *[]ChildrenTask `json:"childrenTask"`
}

type ChildrenTask struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	Content   string    `json:"content"`
	Status    Status    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	TaskID    int32     `json:"taskId"`
}
type Status string

const (
	toDo  Status = "todo"
	doing Status = "doing"
	done  Status = "done"
)

func (e Status) IsValid() bool {
	switch e {
	case toDo, doing, done:
		return true
	}
	return false
}

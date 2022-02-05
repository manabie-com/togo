package togo

import "fmt"

// Task represents a task resource
type Task struct {
	ID     int64
	UserID int64
	Name   string
}

func (t *Task) String() string {
	return fmt.Sprintf("(ID: '%d', UserID: '%d', Name: '%s')", t.ID, t.UserID, t.Name)
}

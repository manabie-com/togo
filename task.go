package togo

import "fmt"

// Task represents a task resource
type Task struct {
	ID          int64
	Description string
}

func (t *Task) String() string {
	return fmt.Sprintf("ID: '%d', Description: '%s'", t.ID, t.Description)
}

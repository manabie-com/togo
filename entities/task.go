package entities

import "fmt"

type Task struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	Completed   bool   `json:"completed"`
	UserID      int    `json:"user_id"`
}

func (t Task) IsValid() error {
	if len(t.Name) == 0 {
		return fmt.Errorf("Task name is required")
	}
	if t.UserID == 0 {
		return fmt.Errorf("Task user id is required")
	}
	return nil
}

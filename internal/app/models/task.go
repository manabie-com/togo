package models

import (
	"errors"
	"html"
	"strings"
)

type Task struct {
	ID          uint64 `json:"id"`
	Content     string `json:"content"`
	UserID      uint64 `json:"user_id"`
	CreatedDate string `json:"created_date"`
}

// Prepare cleans the inputs
func (t *Task) Prepare() {
	t.Content = html.EscapeString(strings.TrimSpace(t.Content))
	t.CreatedDate = html.EscapeString(strings.TrimSpace(t.CreatedDate))
}

// Validate validates the inputs
func (t *Task) Validate(action string) error {
	switch strings.ToLower(action) {
	case "retrieve_tasks":
		if t.CreatedDate == "" {
			return errors.New("create date is required")
		}
	case "add_task":
		if t.Content == "" {
			return errors.New("content is required")
		}
	default:
		return errors.New("need required")
	}
	return nil
}

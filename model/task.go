package model

import "github.com/manabie-com/togo/internal/storages"

// For get tasks
type (
	GetTaskInput struct {
		CreatedDate string `json:"created_date"`
	}

	GetTaskResponse struct {
		Data  []*storages.Task `json:"data"`
		Error string           `json:"error,omitempty"`
	}
)

// For create task
type (
	CreateTaskInput struct {
		Content string `json:"content"`
	}

	CreateTaskResponse struct {
		Data  *storages.Task `json:"data"`
		Error string         `json:"error,omitempty"`
	}
)

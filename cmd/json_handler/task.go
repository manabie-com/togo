package main

import (
	"net/http"

	"github.com/kozloz/togo/internal/tasks"
)

// CreateTaskResult is a JSON object we'll return in our API
type CreateTaskResult struct {
	Error
}

// TaskHandler is a handler that can be attached to a router. This handler specifically manages the Task resource.
type TaskHandler struct {
	op tasks.Operation
}

// CreateTask is a handler that creates the task for the user
func (t *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	_ = t.op.Create(0, "")
}

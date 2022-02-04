package main

import (
	"encoding/json"
	"net/http"

	"github.com/kozloz/togo/internal/errors"
	"github.com/kozloz/togo/internal/tasks"
)

// CreateTaskResult is a JSON object we'll return in our API
type CreateTaskResult struct {
	Error `json:"error"`
}

// TaskHandler is a handler that can be attached to a router. This handler specifically manages the Task resource.
type TaskHandler struct {
	op tasks.Operation
}

// CreateTask is a handler that creates the task for the user
func (t *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	res := CreateTaskResult{}
	//_, err := t.op.Create(0, "")
	var err error
	err = errors.MaxLimit
	if err != nil {
		resErr := CustomErrorToJSON(err)
		res.Error = resErr
		resBody, _ := json.Marshal(res)
		w.Write(resBody)
	}
}

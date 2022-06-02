package controller

import (
	"encoding/json"
	"net/http"
	"togo/common/response"
	"togo/models"
	"togo/service"
)

// Define an interface for the task controller
// `CreateTask` will be the entry point for creating tasks
type TaskController interface {
	// Accept POST body, validate `Task`, then create `Task`
	CreateTask(w http.ResponseWriter, r *http.Request)
}

// Define a Controller struct that contains
// the `Task` Service (business logic for `Task`) attribute
type taskcontroller struct {
	taskservice service.TaskService
}

// Define a Constructor
// Dependency Injection for `Task` Controller
func NewTaskController(service service.TaskService) TaskController {
	return &taskcontroller{
		taskservice: service,
	}
}

// Create a `Task` record
// Route: POST /tasks
// Access: protected
func (c *taskcontroller) CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get POST body
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response.ErrorResponse{
			Status:  "Fail",
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	// Validate Task
	err = c.taskservice.Validate(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response.ErrorResponse{
			Status:  "Fail",
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	// Create Task
	c.taskservice.Create(&task)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.SuccessResponse{
		Status: "Success",
		Code:   http.StatusOK,
		Data:   task,
	})

}

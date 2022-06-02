package service

import (
	"errors"
	"time"
	"togo/models"

	repository "togo/repository/task"

	"github.com/rs/xid"
)

// Define `Task` Service Interface with the following
// Methods which will be utilized by the `TaskController`
type TaskService interface {
	// Check for missing fields
	Validate(task *models.Task) error

	// Generate an ID and set creation time
	Create(task *models.Task) (*models.Task, error)
}

// Define struct with `TaskRepository` as the attribute
// This attribute is responsible for `Task` database interactions
type taskservice struct {
	taskrepository repository.TaskRepository
}

// Define a Constructor
// Dependency Injection for `Task` Service
func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskservice{
		taskrepository: repo,
	}
}

// Validate the `Task`
func (s *taskservice) Validate(task *models.Task) error {
	if task == nil {
		return errors.New("the task is empty")
	}
	if task.Title == "" {
		return errors.New("the task title is empty")
	}
	return nil
}

// Create `Task`
// Generate ID and set the time
func (s *taskservice) Create(task *models.Task) (*models.Task, error) {
	// Generate xid using 3rd party library
	guid := xid.New()

	// Get only id string
	task.TaskID = guid.String()

	// Get current time
	task.CreatedAt = time.Now()

	return s.taskrepository.CreateTask(task)
}

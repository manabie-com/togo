package service

import (
	"errors"
	"time"
	"togo/models"

	"togo/repository"

	"github.com/rs/xid"
)

type TaskService interface {
	Validate(task *models.Task) error
	Create(task *models.Task) (*models.Task, error)
}

type taskservice struct {
	taskrepository repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskservice{
		taskrepository: repo,
	}
}

// Validate the task
func (s *taskservice) Validate(task *models.Task) error {
	if task == nil {
		return errors.New("the task is empty")
	}
	if task.Title == "" {
		return errors.New("the task title is empty")
	}
	return nil
}

// Create task
// Generate ID and set the time
func (s *taskservice) Create(task *models.Task) (*models.Task, error) {
	// Generate xid using 3rd party library
	guid := xid.New()

	// Get only id string
	task.TaskID = guid.String()

	// Get current time
	task.CreatedAt = time.Now()

	return s.taskrepository.Create(task)
}

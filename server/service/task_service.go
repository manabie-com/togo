package service

import (
	"errors"
	"time"
	"togo/models"
	taskrepo "togo/repository/task"
	userrepo "togo/repository/user"

	"github.com/rs/xid"
)

// Define `Task` Service Interface with the following
// Methods which will be utilized by the `TaskController`
type TaskService interface {
	// Check for missing fields
	Validate(task *models.Task) error

	// Generate an ID and set creation time
	Create(task *models.Task) (*models.Task, error)

	// Assign `User` to `Task`
	GetUser(token string, task *models.Task) error

	// Check if `User` exceeds limited number of tasks per day
	GetLimit(token string) error
}

// Define struct with `TaskRepository` as the attribute
// This attribute is responsible for `Task` database interactions
type taskservice struct {
	taskrepository taskrepo.TaskRepository
	userrepository userrepo.UserRepository
}

// Define a Constructor
// Dependency Injection for `Task` Service
func NewTaskService(
	taskrepository taskrepo.TaskRepository,
	userrepository userrepo.UserRepository) TaskService {
	return &taskservice{
		taskrepository: taskrepository,
		userrepository: userrepository,
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

// Retrieve `User` by token
func (s *taskservice) GetUser(token string, task *models.Task) error {
	// Check if session exists
	user, err := s.userrepository.GetUserByToken(token)
	if err != nil {
		return errors.New("session not found")
	}

	// Assign `User` to `Task`
	task.CreatedBy = user.ID
	return nil
}

// Retrieve user by token
func (s *taskservice) GetLimit(token string) error {
	// Get current date
	now := time.Now()
	year, month, day := now.Date()
	today := time.Date(year, month, day, 0, 0, 0, 0, now.Location())

	// Get `User` limit
	user, _ := s.userrepository.GetUserByToken(token)

	// Get `User` tasks by current date
	count, err := s.taskrepository.CountTask(user.ID, today)
	if err != nil {
		return err
	}

	// Check if limit reached
	if count < user.Limit {
		return nil
	} else {
		return errors.New("number of maximum tasks reached")
	}

}

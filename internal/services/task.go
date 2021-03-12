package services

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/entities"
	"github.com/manabie-com/togo/internal/storages"
)

// TaskService defines set of methods or properties needed
// to be implement that will be injected as dependency in
// other components
type TaskService interface {
	ListTasksByUserAndDate(ctx context.Context, userID, createdDate string) ([]*entities.Task, error)
	AddTask(ctx context.Context, task entities.Task) (*entities.Task, error)
}

var (
	// ErrTaskLimitOfDayReached ...
	ErrTaskLimitOfDayReached = errors.New("You have reached the limit of task of provided date")
)

// TaskSvc ...
type TaskSvc struct {
	taskRepo storages.TaskRepository
	userRepo storages.UserRepository
}

// TaskServiceConfiguration ...
type TaskServiceConfiguration struct {
	TaskRepo storages.TaskRepository
	UserRepo storages.UserRepository
}

// NewTaskService ...
func NewTaskService(config TaskServiceConfiguration) *TaskSvc {
	return &TaskSvc{
		taskRepo: config.TaskRepo,
		userRepo: config.UserRepo,
	}
}

// ListTasksByUserAndDate ...
func (ts *TaskSvc) ListTasksByUserAndDate(ctx context.Context, userID, createdDate string) ([]*entities.Task, error) {
	// Validate created date
	tmpTask := entities.Task{CreatedDate: createdDate}
	if isCreatedDateValid := tmpTask.ValidateCreatedDate(); !isCreatedDateValid {
		return nil, entities.ErrTaskInvalidCreatedDate
	}

	// Save to db
	tasks, err := ts.taskRepo.GetTasksByUserIDAndDate(ctx, userID, createdDate)
	if err != nil {
		return nil, ErrInternalError
	}
	return tasks, nil
}

// AddTask ...
func (ts *TaskSvc) AddTask(ctx context.Context, task entities.Task) (*entities.Task, error) {
	if len(task.Content) == 0 {
		return nil, entities.ErrTaskInvalidContent
	}
	task.CreatedDate = time.Now().Format("2006-01-02")
	taskCountByDate, err := ts.taskRepo.CountTasksOfUserByDate(ctx, task.UserID, task.CreatedDate)
	if err != nil {
		return nil, ErrInternalError
	}

	userTaskLimit, err := ts.userRepo.GetUserTaskLimit(ctx, task.UserID)
	if err != nil {
		return nil, ErrInternalError
	}

	if taskCountByDate >= userTaskLimit {
		return nil, ErrTaskLimitOfDayReached
	}
	task.ID = uuid.New().String()
	t, err := ts.taskRepo.SaveTask(ctx, task)
	if err != nil {
		return nil, ErrInternalError
	}
	return t, nil
}

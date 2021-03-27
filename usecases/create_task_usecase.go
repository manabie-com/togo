package usecases

import (
	"context"
	"errors"
	"github.com/manabie-com/togo/domains"
)

var (
	ErrorUserNotFound                 = errors.New("user not found")
	ErrorReachedLimitCreateTaskPerDay = errors.New("reached limit create tasks per day")
)

type (
	CreateTaskUseCase interface {
		Execute(context.Context, TaskInput) (*TaskOutput, error)
	}

	TaskInput struct {
		UserId  int64  `json:"user_id" validate:"required"`
		Content string `json:"content" validate:"required"`
	}

	createTaskInteractor struct {
		taskRepo domains.TaskRepository
		userRepo domains.UserRepository
	}
)

func NewCreateTaskUseCase(taskRepository domains.TaskRepository, userRepository domains.UserRepository) CreateTaskUseCase {
	return createTaskInteractor{
		taskRepo: taskRepository,
		userRepo: userRepository,
	}
}

// Execute create task with dependencies
func (i createTaskInteractor) Execute(ctx context.Context, req TaskInput) (*TaskOutput, error) {
	// check user exists
	user, err := i.userRepo.GetUserById(ctx, req.UserId)
	if err != nil {
		if err == domains.ErrorNotFound {
			return nil, ErrorUserNotFound
		}
		return nil, err
	}

	count, err := i.taskRepo.GetCountCreatedTaskTodayByUser(ctx, user.Id)
	if err != nil {
		if err != domains.ErrorNotFound {
			return nil, err
		}
	}

	// Check user limit today
	if user.MaxTodo > 0 && count >= user.MaxTodo {
		return nil, ErrorReachedLimitCreateTaskPerDay
	}

	task, err := i.taskRepo.CreateTask(ctx, &domains.Task{
		Content: req.Content,
		UserId:  req.UserId,
	})

	if err != nil {
		return nil, err
	}

	tasks := transformTasksToSliceTaskOutput([]*domains.Task{task})
	return tasks[0], nil
}

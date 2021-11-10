package domain

import (
	"context"
	"github.com/manabie-com/togo/internal/domain/entity"
)

type taskUseCase struct {
	taskRepository TaskRepository
	userRepository UserRepository
}

func NewTaskUseCase(taskRepository TaskRepository, userRepository UserRepository) TaskUseCase {
	return taskUseCase{
		taskRepository: taskRepository,
		userRepository: userRepository,
	}
}

func (t taskUseCase) CreateTask(ctx context.Context, content string, username string) error {
	taskCountInDay, err := t.taskRepository.CountTaskInDayByUsername(ctx, username)
	if err != nil {
		return err
	}
	user, err := t.userRepository.GetUser(ctx, username)
	if err != nil {
		return err
	}
	if taskCountInDay+1 > user.MaxTodo {
		return ErrorMaximumTaskPerDay
	}
	err = t.taskRepository.Create(ctx, content, username)
	if err != nil {
		return err
	}
	return nil
}

func (t taskUseCase) GetTask(ctx context.Context, username string, date string) ([]entity.Task, error) {
	return t.taskRepository.GetTaskByUsernameAndDate(ctx, username, date)
}

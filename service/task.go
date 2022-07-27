package service

import (
	"errors"
	"togo/dto"
	"togo/models"
	"togo/repository"
)

type TaskService struct {
	taskRepo *repository.TaskRepository
	userRepo *repository.UserRepository
}

func NewTaskService(
	taskRepo *repository.TaskRepository,
	userRepo *repository.UserRepository,
) *TaskService {
	return &TaskService{taskRepo, userRepo}
}

func (t *TaskService) Create(createTaskDto *dto.CreateTaskDto) (*models.Task, error) {
	userID := createTaskDto.UserID
	user, err := t.userRepo.GetByID(userID)
	if err != nil {
		return nil, err
	}

	numberOfTaskToday, err := t.taskRepo.GetNumberOfUserTaskOnToday(userID)
	if err != nil {
		return nil, err
	}

	if numberOfTaskToday >= user.LimitCount {
		return nil, errors.New("limit_max")
	}

	task := &models.Task{
		UserID:      createTaskDto.UserID,
		Description: createTaskDto.Description,
		EndedAt:     createTaskDto.EndedAt,
	}
	return t.taskRepo.Create(task)
}

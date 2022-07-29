package service

import (
	"errors"
	"togo/internal/models"
	"togo/internal/response"
	"togo/internal/task/dto"
	taskRepo "togo/internal/task/repository"
	userRepo "togo/internal/user/repository"
	"togo/utils"

	"gorm.io/gorm"
)

type TaskService interface {
	Create(createTaskDto *dto.CreateTaskDto, userID int) (*response.TaskResponse, error)
	GetListByUserID(userID int) ([]*response.TaskResponse, error)
}

type taskService struct {
	taskRepo taskRepo.TaskRepository
	userRepo userRepo.UserRepository
}

func NewTaskService(
	taskRepo taskRepo.TaskRepository,
	userRepo userRepo.UserRepository,
) TaskService {
	return &taskService{taskRepo, userRepo}
}

func (t *taskService) Create(createTaskDto *dto.CreateTaskDto, userID int) (*response.TaskResponse, error) {
	user, err := t.userRepo.GetByID(userID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user_not_found")
	}

	numberOfTaskToday, err := t.taskRepo.GetNumberOfUserTaskOnToday(userID)
	if err != nil {
		return nil, err
	}

	if numberOfTaskToday >= user.LimitCount {
		return nil, errors.New("limit_max")
	}

	task := &models.Task{
		UserID:      userID,
		Description: createTaskDto.Description,
		EndedAt:     createTaskDto.EndedAt,
	}
	task, err = t.taskRepo.Create(task)
	if err != nil {
		return nil, err
	}

	var res response.TaskResponse
	err = utils.MarshalDto(&task, &res)
	if err != nil {
		return nil, err
	}
	return &res, err
}

func (t *taskService) GetListByUserID(userID int) ([]*response.TaskResponse, error) {
	user, err := t.userRepo.GetByID(userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("user_not_found")
		}
		return nil, err
	}

	tasks, err := t.taskRepo.GetListByUserID(int(user.ID))
	if err != nil {
		return nil, err
	}

	var res []*response.TaskResponse
	err = utils.MarshalDto(&tasks, &res)
	if err != nil {
		return nil, err
	}
	return res, err
}

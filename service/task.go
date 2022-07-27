package service

import (
	"errors"
	"togo/dto"
	"togo/models"
	"togo/repository"
	"togo/utils"

	"gorm.io/gorm"
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

func (t *TaskService) Create(createTaskDto *dto.CreateTaskDto, userID int) (*dto.TaskResponse, error) {
	user, err := t.userRepo.GetByID(userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("user_not_found")
		}
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
		UserID:      userID,
		Description: createTaskDto.Description,
		EndedAt:     createTaskDto.EndedAt,
	}
	task, err = t.taskRepo.Create(task)
	if err != nil {
		return nil, err
	}

	var res dto.TaskResponse
	err = utils.MarshalDto(&task, &res)
	if err != nil {
		return nil, err
	}
	return &res, err
}

func (t *TaskService) GetListByUserID(userID int) ([]*dto.TaskResponse, error) {
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

	var res []*dto.TaskResponse
	err = utils.MarshalDto(&tasks, &res)
	if err != nil {
		return nil, err
	}
	return res, err
}

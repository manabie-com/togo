package services

import (
	"errors"

	"github.com/manabie-com/togo/entities"
	"github.com/manabie-com/togo/repositories"
)

type ITaskService interface {
	GetAllTask() ([]entities.Task, error)
	CreateTask(task *entities.Task) (error, error)
}

type TaskService struct {
	TaskRepo repositories.ITaskRepository
	UserRepo repositories.IUserRepository
}

func NewTaskService(taskRepo repositories.ITaskRepository, userRepo repositories.IUserRepository) ITaskService {
	return &TaskService{
		TaskRepo: taskRepo,
		UserRepo: userRepo,
	}
}

func (taskService *TaskService) GetAllTask() ([]entities.Task, error) {
	tasks, err := taskService.TaskRepo.GetAllTask()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (taskService *TaskService) CreateTask(task *entities.Task) (error, error) {
	countUserTaskToday, errCount := taskService.TaskRepo.CountTaskByUserIdAndCreatedAt(task.UserId, task.CreatedAt)
	limitUserTaskToday, errCountLimit := taskService.UserRepo.GetLimitTaskPerDay(task.UserId)

	if errCount != nil {
		return errCount, nil
	}

	if errCountLimit != nil {
		return errCountLimit, nil
	}

	if countUserTaskToday < int64(limitUserTaskToday) {
		err := taskService.TaskRepo.CreateTask(task)

		if err != nil {
			return err, nil
		}

		return nil, nil
	} else {
		return nil, errors.New("the task created today exceeded the allowed limit")
	}
}

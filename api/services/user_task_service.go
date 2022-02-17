package services

import (
	"time"

	"github.com/kier1021/togo/api/apierrors.go"
	"github.com/kier1021/togo/api/dto"
	"github.com/kier1021/togo/api/models"
	"github.com/kier1021/togo/api/repositories"
)

type UserTaskService struct {
	userTaskRepo repositories.IUserTaskRepository
}

func NewUserTaskService(userTaskRepo repositories.IUserTaskRepository) *UserTaskService {
	return &UserTaskService{
		userTaskRepo: userTaskRepo,
	}
}

func (srv *UserTaskService) CreateUser(userDto dto.CreateUserDTO) (map[string]interface{}, error) {

	// Check if user already exists
	existingUser, err := srv.userTaskRepo.GetUser(map[string]interface{}{"user_name": userDto.UserName})
	if err != nil {
		return nil, err
	}

	if existingUser != nil {
		return nil, apierrors.UserAlreadyExists
	}

	// Create the user
	lastInsertID, err := srv.userTaskRepo.CreateUser(
		models.User{
			UserName: userDto.UserName,
			MaxTasks: userDto.MaxTasks,
		},
	)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"info": map[string]interface{}{
			"_id":       lastInsertID,
			"user_name": userDto.UserName,
			"max_tasks": userDto.MaxTasks,
		},
	}, nil
}

func (srv *UserTaskService) AddTaskToUser(taskDto dto.CreateTaskDTO) (map[string]interface{}, error) {

	// Get the existing user
	existingUser, err := srv.userTaskRepo.GetUser(map[string]interface{}{
		"user_name": taskDto.UserName,
	})

	if err != nil {
		return nil, err
	}

	// Check if user exists
	if existingUser == nil {
		return nil, apierrors.UserDoesNotExists
	}

	// Get the user tasks
	existingUserTask, err := srv.userTaskRepo.GetUserTask(map[string]interface{}{
		"user_name": taskDto.UserName,
		"ins_day":   taskDto.InsDay,
	})

	if err != nil {
		return nil, err
	}

	// If there is an existing task, check if the maximum tasks has already been reached
	if existingUserTask != nil {
		if len(existingUserTask.Tasks) >= existingUser.MaxTasks {
			return nil, apierrors.MaxTasksReached
		}
	}

	// Upsert the user task
	if err := srv.userTaskRepo.AddTaskToUser(
		models.User{
			ID:       existingUser.ID,
			UserName: existingUser.UserName,
		},
		models.Task{
			Title:       taskDto.Title,
			Description: taskDto.Description,
		},
		taskDto.InsDay,
	); err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"info": map[string]interface{}{
			"title":       taskDto.Title,
			"description": taskDto.Description,
			"user_name":   taskDto.UserName,
		},
	}, nil
}

func (srv *UserTaskService) GetTasksOfUser(getTaskDto dto.GetTaskOfUserDTO) (map[string]interface{}, error) {

	// Set the default ins_day to the current date
	if getTaskDto.InsDay == "" {
		getTaskDto.InsDay = time.Now().Format("2006-01-02")
	}

	// Get the existing user
	existingUser, err := srv.userTaskRepo.GetUser(map[string]interface{}{
		"user_name": getTaskDto.UserName,
	})
	if err != nil {
		return nil, err
	}

	// Check if user exists
	if existingUser == nil {
		return nil, apierrors.UserDoesNotExists
	}

	// Get the current user task
	currentUserTask, err := srv.userTaskRepo.GetUserTask(map[string]interface{}{
		"user_name": getTaskDto.UserName,
		"user_id":   existingUser.ID,
		"ins_day":   getTaskDto.InsDay,
	})

	if err != nil {
		return nil, err
	}

	// Set the tasks.
	// If there are no current task, the tasks will be empty.
	// Otherwise, tasks will be set
	tasks := []map[string]interface{}{}
	if currentUserTask != nil {
		for _, task := range currentUserTask.Tasks {
			tasks = append(tasks, map[string]interface{}{
				"title":       task.Title,
				"description": task.Description,
			})
		}
	}

	return map[string]interface{}{
		"user_task": map[string]interface{}{
			"user_id":   existingUser.ID,
			"user_name": existingUser.UserName,
			"max_tasks": existingUser.MaxTasks,
			"ins_day":   getTaskDto.InsDay,
			"tasks":     tasks,
		},
	}, nil
}

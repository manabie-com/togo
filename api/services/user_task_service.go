package services

import (
	"encoding/json"
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

	dateToday := time.Now().Format("2006-01-02")

	// Check if user already exists
	existingUser, err := srv.userTaskRepo.GetUser(map[string]interface{}{"user_name": userDto.UserName})
	if err != nil {
		return nil, err
	}

	if existingUser != nil {
		return nil, apierrors.UserAlreadyExists
	}

	// Create the user
	user := models.User{
		UserName: userDto.UserName,
		MaxTasks: userDto.MaxTasks,
		InsDay:   dateToday,
	}

	lastInsertID, err := srv.userTaskRepo.CreateUser(user)
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

	dateToday := time.Now().Format("2006-01-02")

	// Get the user
	existingUser, err := srv.userTaskRepo.GetUser(map[string]interface{}{
		"user_name": taskDto.UserName,
	})
	if err != nil {
		return nil, err
	}

	if existingUser == nil {
		return nil, apierrors.UserDoesNotExists
	}

	if len(existingUser.Tasks) >= existingUser.MaxTasks {
		return nil, apierrors.MaxTasksReached
	}

	userTask := models.UserTask{
		User: models.User{
			UserName: taskDto.UserName,
			InsDay:   dateToday,
			MaxTasks: existingUser.MaxTasks,
		},
		Task: models.Task{
			Title:       taskDto.Title,
			Description: taskDto.Description,
		},
	}

	if err := srv.userTaskRepo.AddTaskToUser(userTask); err != nil {
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

	if getTaskDto.InsDay == "" {
		getTaskDto.InsDay = time.Now().Format("2006-01-02")
	}

	existingUser, err := srv.userTaskRepo.GetUser(map[string]interface{}{
		"user_name": getTaskDto.UserName,
	})
	if err != nil {
		return nil, err
	}

	if existingUser == nil {
		return nil, apierrors.UserDoesNotExists
	}

	currentUserTask, err := srv.userTaskRepo.GetUser(map[string]interface{}{
		"user_name": getTaskDto.UserName,
		"ins_day":   getTaskDto.InsDay,
	})
	if err != nil {
		return nil, err
	}

	if currentUserTask == nil {

		// Set the ins_day to the given ins_day
		// Set the tasks to empty tasks
		existingUser.InsDay = getTaskDto.InsDay
		existingUser.Tasks = []models.Task{}

		var userTask map[string]interface{}
		jsonByte, _ := json.Marshal(existingUser)
		json.Unmarshal(jsonByte, &userTask)

		return map[string]interface{}{
			"user_task": userTask,
		}, nil
	}

	var userTask map[string]interface{}
	jsonByte, _ := json.Marshal(currentUserTask)
	json.Unmarshal(jsonByte, &userTask)

	return map[string]interface{}{
		"user_task": userTask,
	}, nil
}

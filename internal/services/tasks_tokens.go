package services

import (
	"github.com/manabie-com/togo/internal/repositories"
)


type TaskService interface {
	AddTask(task *repositories.Task) (*repositories.Task, error)
	ListTask(userID, createdAt string) (*[]repositories.Task, error)
}

type UserService interface {
	ValidateUser(userID, password string) bool
}

type TokenService interface {
	GetAuthToken(userID, password string) (string, error)
	ValidToken(token string) (userID string, valid bool)
}

type ToDoService struct {
	TaskService TaskService
	TokenService TokenService
}

func NewToDoService(taskService TaskService, tokenService TokenService) *ToDoService{
	return &ToDoService {
		TaskService: taskService,
		TokenService: tokenService,
	}
}


func (s *ToDoService) ListTasks(userID, createdAt string) (*[]repositories.Task, error) {
	return s.TaskService.ListTask(userID, createdAt)
}

func (s *ToDoService) AddTask(task *repositories.Task) (*repositories.Task, error) {
	return s.TaskService.AddTask(task)
}

func (s *ToDoService) GetAuthToken(userID, password string) (string, error) {
	return s.TokenService.GetAuthToken(userID, password)
}

func (s *ToDoService) ValidToken(token string) (userID string, valid bool) {
	return s.TokenService.ValidToken(token)
}


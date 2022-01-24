package service

import (
	taskservice "github.com/trinhdaiphuc/togo/internal/service/task"
	userservice "github.com/trinhdaiphuc/togo/internal/service/user"
)

type Service struct {
	UserService userservice.UserService
	TaskService taskservice.TaskService
}

func NewService(user userservice.UserService, task taskservice.TaskService) *Service {
	return &Service{
		UserService: user,
		TaskService: task,
	}
}

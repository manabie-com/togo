package service

import (
	taskservice "todo/internal/service/task"
	userservice "todo/internal/service/user"
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

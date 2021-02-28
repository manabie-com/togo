package up

import (
	"context"
)

type UserService interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
}

type TaskService interface {
	ListTasks(context.Context, *ListTasksRequest) (*ListTasksResponse, error)
	AddTask(context.Context, *AddTaskRequest) (*Task, error)
}

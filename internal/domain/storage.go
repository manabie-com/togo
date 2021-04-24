package domain

import "fmt"

type TaskStore interface {
	AddTaskWithLimitPerDay(task Task, limit int) error
	GetTasksByUserID(userID string) ([]Task, error)
}

type UserStore interface {
	FindUserByID(string) (User, error)
	CreateUser(User) error
	GetUserTasksPerDay(userID string) (int, error)
}

type UserExist string

func (u UserExist) Error() string {
	return fmt.Sprintf("user with ID %s exists already", string(u))
}

type UserNotFound string

func (u UserNotFound) Error() string {
	return fmt.Sprintf("user with ID %s not found", string(u))
}

type TaskLimitReached struct{}

func (u TaskLimitReached) Error() string {
	return "task limit reached"
}

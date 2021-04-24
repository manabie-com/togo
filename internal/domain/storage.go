package domain

import (
	"errors"
	"fmt"
)

type TaskStore interface {
	AddTaskWithLimitPerDay(task Task, limit int) error
	GetTasksByUserIDAndDate(userID string, date string) ([]Task, error)
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

var TaskLimitReached = errors.New("task limit reached")

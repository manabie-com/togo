package domain

import "fmt"

type TaskStore interface {
	AddTaskWithLimitPerDay(task Task, limit int) error
	GetTasksByUserID(userID string) ([]Task, error)
}

type UserStore interface {
	FindUserByID(string) (User, error)
	CreateUser(Id string, Password string) error
	GetUserTasksPerDay(userID string) (int, error)
}

type UserNotFound struct {
	ID string
}

func (u UserNotFound) Error() string {
	return fmt.Sprintf("user with ID %s not found", u.ID)
}

type TaskLimitReached struct{}

func (u TaskLimitReached) Error() string {
	return "task limit reached"
}

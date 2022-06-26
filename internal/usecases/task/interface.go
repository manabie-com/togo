package task

import (
	"example.com/m/v2/internal/models"
)

type Reader interface {
	FindTaskByUser(username, createDate string) ([]models.Task, error)
}

type Writer interface {
	AddTask(task *models.Task) error
}

type TaskRepository interface {
	Reader
	Writer
}

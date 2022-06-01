package repository

import "togo/models"

type TaskRepository interface {
	CreateTask(task *models.Task) (*models.Task, error)
}

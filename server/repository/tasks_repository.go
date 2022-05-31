package repository

import "togo/models"

type TaskRepository interface {
	Create(task *models.Task) (*models.Task, error)
}

package service

import "togo/domain/model"

type TaskService interface {
	CreateTask(task *model.Task) (*model.Task, error)
}

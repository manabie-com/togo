package task

import "github.com/manabie-com/togo/internal/storages"

type TaskStorageInterface interface {
	CreateTask(task *storages.Task) error
	RetrieveTasks(userID, createdDate string) ([]*storages.Task, error)
}

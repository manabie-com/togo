package task

import "github.com/manabie-com/togo/internal/storages"

type TaskUsecaseInterface interface {
	RetrieveTasks(userID, createdDate string) ([]*storages.Task, error)
}

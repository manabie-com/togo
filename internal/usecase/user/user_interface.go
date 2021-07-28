package user

import "github.com/manabie-com/togo/internal/storages"

type UserUsecaseInterface interface {
	ValidateUser(id, password string) error
	CreateTask(task *storages.Task) error
}

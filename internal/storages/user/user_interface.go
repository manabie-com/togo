package user

import "github.com/manabie-com/togo/internal/storages"

type UserStorageInterface interface {
	GetUser(id, password string) error
	GetUsersTasks(userID string, craetedDate string) (storages.User, error)
}

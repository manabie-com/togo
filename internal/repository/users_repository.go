package repository

import (
	"github.com/manabie-com/togo/internal/model"
)
type UsersRepository interface {
	GetUserByIdAndOPassword(id string, password string) (model.Users, error)
	GetUserById(id string) (model.Users, error)
}
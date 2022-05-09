package repository

import (
	"github.com/jfzam/togo/domain/entity"
)

type UserRepository interface {
	SaveUser(*entity.User) (*entity.User, map[string]string)
	GetUser(uint64) (*entity.User, error)
	GetUsers() ([]entity.User, error)
	GetUserByUsernameAndPassword(*entity.User) (*entity.User, map[string]string)
}

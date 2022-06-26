package user

import (
	"github.com/manabie-com/togo/internal/models"
)

type Reader interface {
	GetByUsername(username string) (*models.User, error)
	Login(username, password string) (*models.User, error)
}

type Writer interface {
	Create(u *models.User) error
}

type UserRepository interface {
	Reader
	Writer
}

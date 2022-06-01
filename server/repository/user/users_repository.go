package repository

import "togo/models"

type UserRepository interface {
	Register(user *models.User) (*models.User, error)
	GetUser(email string, user *models.User) (models.User, error)
	Login(user *models.User) error
}

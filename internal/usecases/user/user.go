package user

import (
	"github.com/manabie-com/togo/internal/models"
)

// UserUsecase is the definition for collection of methods related to the `users` table use case
type UserUseCase interface {
	Login(username, password string) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	Create(u *models.User) error
}

type userUseCaseRepository struct {
	repository UserRepository
}

// NewUserUsecase returns a UserUsecase attached with methods related to the `users` table use case
func NewUserUseCase(repository UserRepository) UserUseCase {
	return &userUseCaseRepository{repository: repository}
}

// GetUserByUsername implements UserUseCase
func (uc *userUseCaseRepository) GetUserByUsername(username string) (*models.User, error) {
	return uc.repository.GetByUsername(username)
}

// Login implements UserUseCase
func (uc *userUseCaseRepository) Login(username, password string) (*models.User, error) {
	return uc.repository.Login(username, password)
}

// Create implements UserUseCase
func (uc *userUseCaseRepository) Create(u *models.User) error {
	return uc.repository.Create(u)
}

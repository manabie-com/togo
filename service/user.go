package service

import (
	"togo/dto"
	"togo/models"
	"togo/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(
	userRepo *repository.UserRepository,
) *UserService {
	return &UserService{userRepo}
}

func (t *UserService) Create(createUserDto *dto.CreateUserDto) (*models.User, error) {
	user := &models.User{
		Name:       createUserDto.Name,
		LimitCount: createUserDto.LimitCount,
	}
	return t.userRepo.Create(user)
}

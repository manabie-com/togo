package service

import (
	"togo/dto"
	"togo/models"
	"togo/repository"
	"togo/utils"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(
	userRepo *repository.UserRepository,
) *UserService {
	return &UserService{userRepo}
}

func (t *UserService) Create(createUserDto *dto.CreateUserDto) (*dto.UserResponse, error) {
	user := &models.User{
		Name:       createUserDto.Name,
		LimitCount: createUserDto.LimitCount,
	}
	user, err := t.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	var res dto.UserResponse
	err = utils.MarshalDto(&user, &res)
	if err != nil {
		return nil, err
	}
	return &res, err
}

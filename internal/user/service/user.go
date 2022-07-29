package service

import (
	"errors"
	"togo/internal/models"
	"togo/internal/response"
	"togo/internal/user/dto"
	"togo/internal/user/repository"
	"togo/utils"
)

type UserService interface {
	Create(createUserDto *dto.CreateUserDto) (*response.UserResponse, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(
	userRepo repository.UserRepository,
) UserService {
	return &userService{userRepo}
}

func (t *userService) Create(createUserDto *dto.CreateUserDto) (*response.UserResponse, error) {
	userExist, err := t.userRepo.GetByName(createUserDto.Name)
	if err != nil {
		return nil, err
	}
	if userExist != nil {
		return nil, errors.New("user_with_name_exist")
	}

	user := &models.User{
		Name:       createUserDto.Name,
		LimitCount: createUserDto.LimitCount,
	}
	user, err = t.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	var res response.UserResponse
	err = utils.MarshalDto(&user, &res)
	if err != nil {
		return nil, err
	}
	return &res, err
}

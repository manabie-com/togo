package service

import (
	"errors"
	"togo/internal/dto"
	"togo/internal/models"
	"togo/internal/repository"
	"togo/utils"

	"gorm.io/gorm"
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
	userExist, err := t.userRepo.GetByName(createUserDto.Name)
	if err != nil && err != gorm.ErrRecordNotFound {
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

	var res dto.UserResponse
	err = utils.MarshalDto(&user, &res)
	if err != nil {
		return nil, err
	}
	return &res, err
}

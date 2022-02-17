package services

import (
	"github.com/kier1021/togo/api/apierrors.go"
	"github.com/kier1021/togo/api/dto"
	"github.com/kier1021/togo/api/models"
	"github.com/kier1021/togo/api/repositories"
)

type UserService struct {
	userRepo repositories.IUserRepository
}

func NewUserService(userRepo repositories.IUserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (srv *UserService) CreateUser(userDto dto.CreateUserDTO) (map[string]interface{}, error) {

	// Check if user already exists
	existingUser, err := srv.userRepo.GetUser(map[string]interface{}{"user_name": userDto.UserName})
	if err != nil {
		return nil, err
	}

	if existingUser != nil {
		return nil, apierrors.UserAlreadyExists
	}

	// Create the user
	lastInsertID, err := srv.userRepo.CreateUser(
		models.User{
			UserName: userDto.UserName,
			MaxTasks: userDto.MaxTasks,
		},
	)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"info": map[string]interface{}{
			"_id":       lastInsertID,
			"user_name": userDto.UserName,
			"max_tasks": userDto.MaxTasks,
		},
	}, nil
}

func (srv *UserService) GetUsers() (map[string]interface{}, error) {
	users, err := srv.userRepo.GetUsers(map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"users": users,
	}, nil
}

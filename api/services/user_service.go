package services

import (
	"github.com/kier1021/togo/api/apierrors"
	"github.com/kier1021/togo/api/dto"
	"github.com/kier1021/togo/api/models"
	"github.com/kier1021/togo/api/repositories"
	"github.com/kier1021/togo/api/validation"
)

// UserService holds the business logic for user entity
type UserService struct {
	userRepo repositories.IUserRepository
}

// NewUserService is the constructor for UserService
func NewUserService(userRepo repositories.IUserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

// CreateUser creates a user
func (srv *UserService) CreateUser(userDto dto.CreateUserDTO) (map[string]interface{}, error) {

	// Validate the data
	v := validation.NewValidator()
	err := v.Struct(userDto)
	if err != nil {
		return nil, err
	}

	// Check if user already exists
	existingUser, err := srv.userRepo.GetUser(map[string]interface{}{"user_name": userDto.UserName})
	if err != nil {
		return nil, err
	}

	if existingUser != nil {
		return nil, apierrors.NewUserAlreadyExistsError("user_name", userDto.UserName)
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

// GetUsers return all the users
func (srv *UserService) GetUsers() (map[string]interface{}, error) {

	// Get the users from repository
	users, err := srv.userRepo.GetUsers(map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"users": users,
	}, nil
}

package services

import (
	"errors"
	"log"
	"reflect"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/kier1021/togo/api/apierrors"
	"github.com/kier1021/togo/api/dto"
	"github.com/kier1021/togo/api/models"
	"github.com/kier1021/togo/api/repositories"
)

const (
	UserAlreadyExistsError = "*apierror.UserAlreadyExistsError"
	UserDoesNotExistsError = "*apierror.UserDoesNotExistsError"
	MaxTasksReachedError   = "*apierror.MaxTasksReachedError"
	ValidationErrors       = "validator.ValidationErrors"
)

type testCase struct {
	name             string
	createUserDTO    dto.CreateUserDTO
	createTaskDTO    dto.CreateTaskDTO
	getTaskOfUserDTO dto.GetTaskOfUserDTO
	expected         map[string]interface{}
	hasError         bool
	errType          string
}

func assertTestCase(t *testing.T, tc testCase, res interface{}, err error) {

	if tc.hasError {

		if !assertErrorType(err, tc.errType) {
			t.Errorf("Test '%s' failed. Expecting an error type %s got %T", tc.name, tc.errType, err)
			return
		}

	} else {

		if err != nil {
			t.Errorf("Test '%s' failed. Expecting a nil error got %T", tc.name, err)
			return
		}

	}

	if !reflect.DeepEqual(res, tc.expected) {
		t.Errorf("Test '%s' failed.\n Expected value %v. \n Got %v.", tc.name, tc.expected, res)
	}

}

func assertErrorType(err error, errType string) bool {
	switch errType {
	case UserAlreadyExistsError:
		_, ok := err.(*apierrors.UserAlreadyExistsError)
		return ok
	case UserDoesNotExistsError:
		_, ok := err.(*apierrors.UserDoesNotExistsError)
		return ok
	case MaxTasksReachedError:
		_, ok := err.(*apierrors.MaxTasksReachedError)
		return ok
	case ValidationErrors:
		var ve validator.ValidationErrors
		return errors.As(err, &ve)
	default:
		log.Fatal("Error type provided is not supported. Test failed.")
		return false
	}
}

func TestCreateUser(t *testing.T) {
	userMockRepo := repositories.NewUserMockRepository()
	userSrv := NewUserService(userMockRepo)

	testCases := []testCase{
		{
			name: "User should be created successfully.",
			createUserDTO: dto.CreateUserDTO{
				UserName: "Unit Test User 1",
				MaxTasks: 5,
			},
			expected: map[string]interface{}{
				"info": map[string]interface{}{
					"_id":       "620e6baff70a3fd2fc8811a0",
					"user_name": "Unit Test User 1",
					"max_tasks": 5,
				},
			},
		},
		{
			name: "validator.ValidationErrors must be returned when the given UserName is empty",
			createUserDTO: dto.CreateUserDTO{
				UserName: "   ",
				MaxTasks: 5,
			},
			hasError: true,
			errType:  ValidationErrors,
		},
		{
			name: "validator.ValidationErrors must be returned when the given MaxTasks is 0",
			createUserDTO: dto.CreateUserDTO{
				UserName: "Unit Test User 1",
				MaxTasks: 0,
			},
			hasError: true,
			errType:  ValidationErrors,
		},
		{
			name: "*apierrors.UserAlreadyExistError must be returned when the given UserName already exists.",
			createUserDTO: dto.CreateUserDTO{
				UserName: "Test User 1",
				MaxTasks: 5,
			},
			hasError: true,
			errType:  UserAlreadyExistsError,
		},
	}

	for _, tc := range testCases {
		res, err := userSrv.CreateUser(tc.createUserDTO)
		assertTestCase(t, tc, res, err)
	}
}

func TestGetUsers(t *testing.T) {
	userMockRepo := repositories.NewUserMockRepository()
	userSrv := NewUserService(userMockRepo)

	testCases := []testCase{
		{
			name: "Should return all users",
			expected: map[string]interface{}{
				"users": []models.User{
					{
						ID:       "620e6b6e20bdcb887326931a",
						UserName: "Test User 1",
						MaxTasks: 3,
					},
					{
						ID:       "620e6b79657f405b5f79b32d",
						UserName: "Test User 2",
						MaxTasks: 4,
					},
					{
						ID:       "620e6b7e64b5c80f08aaddcd",
						UserName: "Test User 3",
						MaxTasks: 2,
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		res, err := userSrv.GetUsers()
		assertTestCase(t, tc, res, err)
	}
}

package service

import (
	"testing"
	"time"
	"togo/models"
	"togo/utils/security"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Register(user *models.User) (*models.User, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(*models.User), args.Error(1)
}

func (m *MockUserRepository) GetUserByEmail(email string, user *models.User) (models.User, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(models.User), args.Error(1)
}

func (m *MockUserRepository) GetUserByToken(token string) (models.User, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(models.User), args.Error(1)
}

func (m *MockUserRepository) Login(user *models.User) error {
	args := m.Called(user)
	return args.Error(1)
}

func TestValidateRegistration(t *testing.T) {
	testService := NewUserService(nil)
	err := testService.ValidateRegistration(nil)

	// Assert Nil
	assert.NotNil(t, err)

	// Assert Error message
	expected := "the user is empty"
	actual := err.Error()
	assert.Equal(t, expected, actual)
}

func TestValidateLogin(t *testing.T) {
	mockUserRepository := new(MockUserRepository)

	// Define mock `User` fields
	id := "a67980sdf78as"
	email := "admin@gmail.com"
	password := "admin"
	token := "890asdfdf"
	limit := 3
	hashedPassword, _ := security.HashPassword(password)

	// Mock User
	mockuser := models.User{
		ID:       id,
		Email:    email,
		Password: hashedPassword,
		Token:    token,
		Limit:    limit,
	}

	// Input User
	inputuser := models.User{
		Email:    email,
		Password: password,
	}

	// Set expectations
	mockUserRepository.On("GetUserByEmail").Return(mockuser, nil)

	// Initialize Service
	testService := NewUserService(mockUserRepository)

	// Get result from tested service method
	result := testService.ValidateLogin(&inputuser)

	// Assert if User ID is saved
	assert.NotNil(t, inputuser.ID)

	// Assert if password is correct
	assert.Equal(t, nil, result)
}

func TestValidateLoginIncorrectPassword(t *testing.T) {
	// Instantiate mock `User` repository
	mockUserRepository := new(MockUserRepository)

	// Define mock `User` fields
	id := "a67980sdf78as"
	email := "admin@gmail.com"
	password := "admin"
	incorrectPassword := "admin2"
	token := "890asdfdf"
	limit := 3
	hashedPassword, _ := security.HashPassword(password)

	// Mock User
	mockuser := models.User{
		ID:       id,
		Email:    email,
		Password: hashedPassword,
		Token:    token,
		Limit:    limit,
	}

	// Input User
	inputuser := models.User{
		Email:    email,
		Password: incorrectPassword,
	}

	// Set expectations
	mockUserRepository.On("GetUserByEmail").Return(mockuser, nil)

	// Initialize Service
	testService := NewUserService(mockUserRepository)

	// Get result from tested service method
	result := testService.ValidateLogin(&inputuser)

	// Assert if User ID is saved
	assert.NotNil(t, inputuser.ID)

	// Assert if password is incorrect
	expected := "password incorrect"
	assert.Equal(t, expected, result.Error())
}

func TestGenerateJWT(t *testing.T) {
	// Instantiate mock `User` repository
	mockUserRepository := new(MockUserRepository)

	// Initialize Service
	testService := NewUserService(mockUserRepository)

	// Define mock `User` fields
	id := "a67980sdf78as"

	// Mock User
	mockuser := models.User{ID: id}

	// Set expiration time of the token
	expiration := time.Now().Add(time.Minute * 60)

	// Get result from tested service method
	_, err := testService.GenerateJWT(&mockuser, expiration)

	// Assert if JWT is Generated if there is no error
	assert.Nil(t, err)

	// Assert if Token is assigned to User
	assert.NotNil(t, mockuser.Token)
}

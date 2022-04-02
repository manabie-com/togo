package service

import (
	"errors"
	"github.com/ansidev/togo/auth/dto"
	authMock "github.com/ansidev/togo/auth/mock"
	"github.com/ansidev/togo/domain/user"
	"github.com/ansidev/togo/errs"
	"github.com/ansidev/togo/test"
	userMock "github.com/ansidev/togo/user/mock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

func TestAuthService(t *testing.T) {
	suite.Run(t, new(AuthServiceTestSuite))
}

type AuthServiceTestSuite struct {
	test.ServiceTestSuite
	mockUserRepository *userMock.MockIUserRepository
	mockCredRepository *authMock.MockICredRepository
}

func (s *AuthServiceTestSuite) SetupSuite() {
	s.ServiceTestSuite.SetupSuite()
	s.mockUserRepository = userMock.NewMockIUserRepository(s.Ctrl)
	s.mockCredRepository = authMock.NewMockICredRepository(s.Ctrl)
}

func (s *AuthServiceTestSuite) TestLogin_WhenUsernameDoesNotExist_ShouldReturnErrUsernameNotFound() {
	authenticateRequest := dto.UsernamePasswordCredential{
		Username: "test_user",
		Password: "test_password",
	}

	s.mockUserRepository.
		EXPECT().
		FindFirstByUsername(authenticateRequest.Username).
		Return(user.User{}, errors.New(errs.ErrUsernameNotFound))

	authorService := NewAuthService(s.mockUserRepository, s.mockCredRepository)
	_, err := authorService.Login(authenticateRequest)

	require.Error(s.T(), err)
	require.Equal(s.T(), errs.ErrUsernameNotFound, errs.Message(err))
	require.Equal(s.T(), errs.ErrCodeUsernameNotFound, errs.ErrorCode(err))
}

func (s *AuthServiceTestSuite) TestLogin_WhenPasswordIsWrong_ShouldReturnErrWrongUserPassword() {
	authenticateRequest := dto.UsernamePasswordCredential{
		Username: "test_user",
		Password: "wrong_password",
	}

	mockUser := user.User{
		ID:           1,
		Username:     "test_user",
		Password:     "$2a$12$IsAJrIc1yhMtlcXC1KfhLOqJSon.NAUMo3KG8NHA9myPm05F85Id2", // test_password
		MaxDailyTask: 5,
		CreatedAt:    time.Date(2022, 2, 22, 2, 34, 56, 0, time.UTC),
		UpdatedAt:    time.Date(2022, 2, 22, 2, 34, 56, 0, time.UTC),
	}
	s.mockUserRepository.
		EXPECT().
		FindFirstByUsername(authenticateRequest.Username).
		Return(mockUser, nil)

	authorService := NewAuthService(s.mockUserRepository, s.mockCredRepository)
	_, err := authorService.Login(authenticateRequest)

	require.Error(s.T(), err)
	require.Equal(s.T(), errs.ErrWrongUserPassword, errs.Message(err))
	require.Equal(s.T(), errs.ErrCodeWrongUserPassword, errs.ErrorCode(err))
}

func (s *AuthServiceTestSuite) TestLogin_WhenSavingTokenWasFailed_ShouldReturnErrCouldNotSaveAuthenticateCredential() {
	authenticateRequest := dto.UsernamePasswordCredential{
		Username: "test_user",
		Password: "test_password",
	}

	mockUser := user.User{
		ID:           1,
		Username:     "test_user",
		Password:     "$2a$12$IsAJrIc1yhMtlcXC1KfhLOqJSon.NAUMo3KG8NHA9myPm05F85Id2", // test_password
		MaxDailyTask: 5,
		CreatedAt:    time.Date(2022, 2, 22, 2, 34, 56, 0, time.UTC),
		UpdatedAt:    time.Date(2022, 2, 22, 2, 34, 56, 0, time.UTC),
	}

	s.mockUserRepository.
		EXPECT().
		FindFirstByUsername(authenticateRequest.Username).
		Return(mockUser, nil)

	s.mockCredRepository.
		EXPECT().
		Save(mockUser).
		Return("", errors.New("fake_error"))

	authorService := NewAuthService(s.mockUserRepository, s.mockCredRepository)
	_, err := authorService.Login(authenticateRequest)

	require.Error(s.T(), err)
	require.Equal(s.T(), errs.ErrCouldNotSaveToken, errs.Message(err))
	require.Equal(s.T(), errs.ErrCodeCouldNotSaveToken, errs.ErrorCode(err))
}

func (s *AuthServiceTestSuite) TestLogin_WhenPasswordIsRight_ShouldReturnToken() {
	authenticateRequest := dto.UsernamePasswordCredential{
		Username: "test_user",
		Password: "test_password",
	}

	mockUser := user.User{
		ID:           1,
		Username:     "test_user",
		Password:     "$2a$12$IsAJrIc1yhMtlcXC1KfhLOqJSon.NAUMo3KG8NHA9myPm05F85Id2", // test_password
		MaxDailyTask: 5,
		CreatedAt:    time.Date(2022, 2, 22, 2, 34, 56, 0, time.UTC),
		UpdatedAt:    time.Date(2022, 2, 22, 2, 34, 56, 0, time.UTC),
	}

	s.mockUserRepository.
		EXPECT().
		FindFirstByUsername(authenticateRequest.Username).
		Return(mockUser, nil)

	expectedToken := uuid.NewString()
	s.mockCredRepository.
		EXPECT().
		Save(mockUser).
		Return(expectedToken, nil)

	authorService := NewAuthService(s.mockUserRepository, s.mockCredRepository)
	token, err := authorService.Login(authenticateRequest)

	require.NoError(s.T(), err)
	require.Equal(s.T(), expectedToken, token)
}

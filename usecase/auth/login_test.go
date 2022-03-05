package auth_test

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/khangjig/togo/model"
	"github.com/khangjig/togo/usecase/auth"
	"github.com/khangjig/togo/util/myerror"
)

func (suite *TestSuite) TestLogin_Success() {
	req := &auth.LoginRequest{
		Email:    "admin@gmail.com",
		Password: "12345678",
	}

	mockUser := &model.User{
		ID:       1,
		Email:    "admin@gmail.com",
		Name:     "admin",
		Gender:   0,
		Password: "$2a$10$kXOeCosi13/0Fk8T7DihYOfHvgnfzaGTWwp1ypqJcWKxXx28Hozom",
	}

	suite.mockUserRepo.On("GetByEmail", suite.ctx, req.Email).Return(mockUser, nil)

	// execute
	_, err := suite.useCase.Login(suite.ctx, req)

	suite.Nil(err)
}

func (suite *TestSuite) TestLogin_EmptyEmail() {
	req := &auth.LoginRequest{
		Email:    "",
		Password: "12345678",
	}

	// execute
	_, err := suite.useCase.Login(suite.ctx, req)
	expectErr := myerror.ErrInvalidEmailPassword()
	myErr := err.(myerror.MyError)

	suite.Equal(expectErr.ErrorCode, myErr.ErrorCode)
	suite.Equal(expectErr.HTTPCode, myErr.HTTPCode)
}

func (suite *TestSuite) TestLogin_NotEmail() {
	req := &auth.LoginRequest{
		Email:    "admin@gmailcom",
		Password: "12345678",
	}

	// execute
	_, err := suite.useCase.Login(suite.ctx, req)
	expectErr := myerror.ErrInvalidEmailPassword()
	myErr := err.(myerror.MyError)

	suite.Equal(expectErr.ErrorCode, myErr.ErrorCode)
	suite.Equal(expectErr.HTTPCode, myErr.HTTPCode)
}

func (suite *TestSuite) TestLogin_UserNotFound() {
	req := &auth.LoginRequest{
		Email:    "admin@gmail.com",
		Password: "12345678",
	}

	suite.mockUserRepo.On("GetByEmail", suite.ctx, req.Email).Return(nil, gorm.ErrRecordNotFound)

	// execute
	_, err := suite.useCase.Login(suite.ctx, req)
	expectErr := myerror.ErrInvalidEmailPassword()
	myErr := err.(myerror.MyError)

	suite.Equal(expectErr.ErrorCode, myErr.ErrorCode)
	suite.Equal(expectErr.HTTPCode, myErr.HTTPCode)
}

func (suite *TestSuite) TestLogin_ErrorQuery() {
	req := &auth.LoginRequest{
		Email:    "admin@gmail.com",
		Password: "12345678",
	}

	suite.mockUserRepo.On("GetByEmail", suite.ctx, req.Email).Return(nil, errors.New("error query"))

	// execute
	_, err := suite.useCase.Login(suite.ctx, req)
	expectErr := myerror.ErrGet(err)
	myErr := err.(myerror.MyError)

	suite.Equal(expectErr.ErrorCode, myErr.ErrorCode)
	suite.Equal(expectErr.HTTPCode, myErr.HTTPCode)
}

func (suite *TestSuite) TestLogin_WrongPassword() {
	req := &auth.LoginRequest{
		Email:    "admin@gmail.com",
		Password: "1234567",
	}

	mockUser := &model.User{
		ID:       1,
		Email:    "admin@gmail.com",
		Name:     "admin",
		Gender:   0,
		Password: "$2a$10$kXOeCosi13/0Fk8T7DihYOfHvgnfzaGTWwp1ypqJcWKxXx28Hozom",
	}

	suite.mockUserRepo.On("GetByEmail", suite.ctx, req.Email).Return(mockUser, nil)

	// execute
	_, err := suite.useCase.Login(suite.ctx, req)
	expectErr := myerror.ErrInvalidEmailPassword()
	myErr := err.(myerror.MyError)

	suite.Equal(expectErr.ErrorCode, myErr.ErrorCode)
	suite.Equal(expectErr.HTTPCode, myErr.HTTPCode)
}

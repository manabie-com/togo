package todo_test

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"

	"github.com/khangjig/togo/usecase/todo"
	"github.com/khangjig/togo/util/myerror"
)

func (suite *TestSuite) TestCreate_Success() {
	req := &todo.CreateRequest{
		Title:   "title",
		Content: "content",
		Status:  0,
	}

	suite.mockUserCacheRepo.On("GetTotalTodoByUserID", suite.ctx, suite.userClaims.ID).Return(0, nil)
	suite.mockTodoRepo.On("Create", suite.ctx, mock.Anything).Return(nil)
	suite.mockUserCacheRepo.On("SetTotalTodoByUserID", suite.ctx, suite.userClaims.ID, 1).Return(nil)

	// execute
	_, err := suite.useCase.Create(suite.ctx, req)

	suite.Nil(err)
}

func (suite *TestSuite) TestCreate_InvalidTitle() {
	req := &todo.CreateRequest{
		Title:   "",
		Content: "content",
		Status:  0,
	}

	// execute
	_, err := suite.useCase.Create(suite.ctx, req)

	expectErr := myerror.ErrTodoTitleInvalid("Invalid title.")
	myErr := err.(myerror.MyError)

	suite.Equal(expectErr.ErrorCode, myErr.ErrorCode)
	suite.Equal(expectErr.HTTPCode, myErr.HTTPCode)
}

// nolint:lll,misspell
func (suite *TestSuite) TestCreate_TooLongTitle() {
	req := &todo.CreateRequest{
		Title:   "Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor. Aenean massa. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus.Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor. Aenean massa. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus.",
		Content: "content",
		Status:  0,
	}

	// execute
	_, err := suite.useCase.Create(suite.ctx, req)

	expectErr := myerror.ErrTodoTitleInvalid("Invalid title.")
	myErr := err.(myerror.MyError)

	suite.Equal(expectErr.ErrorCode, myErr.ErrorCode)
	suite.Equal(expectErr.HTTPCode, myErr.HTTPCode)
}

func (suite *TestSuite) TestCreate_InvalidStatus() {
	req := &todo.CreateRequest{
		Title:   "title",
		Content: "content",
		Status:  7,
	}

	// execute
	_, err := suite.useCase.Create(suite.ctx, req)

	expectErr := myerror.ErrTodoStatusInvalid()
	myErr := err.(myerror.MyError)

	suite.Equal(expectErr.ErrorCode, myErr.ErrorCode)
	suite.Equal(expectErr.HTTPCode, myErr.HTTPCode)
}

func (suite *TestSuite) TestCreate_ErrGetRedis() {
	req := &todo.CreateRequest{
		Title:   "title",
		Content: "content",
		Status:  0,
	}

	suite.mockUserCacheRepo.On("GetTotalTodoByUserID", suite.ctx, suite.userClaims.ID).Return(0, errors.New("error"))

	// execute
	_, err := suite.useCase.Create(suite.ctx, req)

	expectErr := myerror.ErrGetRedis(err)
	myErr := err.(myerror.MyError)

	suite.Equal(expectErr.ErrorCode, myErr.ErrorCode)
	suite.Equal(expectErr.HTTPCode, myErr.HTTPCode)
}

func (suite *TestSuite) TestCreate_TodoMaxLimit() {
	req := &todo.CreateRequest{
		Title:   "title",
		Content: "content",
		Status:  0,
	}

	suite.mockUserCacheRepo.On("GetTotalTodoByUserID", suite.ctx, suite.userClaims.ID).Return(11, nil)

	// execute
	_, err := suite.useCase.Create(suite.ctx, req)

	expectErr := myerror.ErrTodoMaxLimit()
	myErr := err.(myerror.MyError)

	suite.Equal(expectErr.ErrorCode, myErr.ErrorCode)
	suite.Equal(expectErr.HTTPCode, myErr.HTTPCode)
}

func (suite *TestSuite) TestCreate_ErrCreate() {
	req := &todo.CreateRequest{
		Title:   "title",
		Content: "content",
		Status:  0,
	}

	suite.mockUserCacheRepo.On("GetTotalTodoByUserID", suite.ctx, suite.userClaims.ID).Return(0, nil)
	suite.mockTodoRepo.On("Create", suite.ctx, mock.Anything).Return(errors.New("error"))

	// execute
	_, err := suite.useCase.Create(suite.ctx, req)

	expectErr := myerror.ErrCreate(err)
	myErr := err.(myerror.MyError)

	suite.Equal(expectErr.ErrorCode, myErr.ErrorCode)
	suite.Equal(expectErr.HTTPCode, myErr.HTTPCode)
}

func (suite *TestSuite) TestCreate_ErrSetRedis() {
	req := &todo.CreateRequest{
		Title:   "title",
		Content: "content",
		Status:  0,
	}

	suite.mockUserCacheRepo.On("GetTotalTodoByUserID", suite.ctx, suite.userClaims.ID).Return(0, nil)
	suite.mockTodoRepo.On("Create", suite.ctx, mock.Anything).Return(nil)
	suite.mockUserCacheRepo.On("SetTotalTodoByUserID", suite.ctx, suite.userClaims.ID, 1).Return(errors.New("error"))

	// execute
	_, err := suite.useCase.Create(suite.ctx, req)

	expectErr := myerror.ErrSetRedis(err)
	myErr := err.(myerror.MyError)

	suite.Equal(expectErr.ErrorCode, myErr.ErrorCode)
	suite.Equal(expectErr.HTTPCode, myErr.HTTPCode)
}

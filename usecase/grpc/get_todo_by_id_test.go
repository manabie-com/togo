package grpc_test

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/khangjig/togo/model"
	"github.com/khangjig/togo/util/myerror"
)

func (suite *TestSuite) TestGetTodoByID_Success() {
	var todoID int64 = 1

	mockTodo := &model.Todo{
		ID:      1,
		UserID:  1,
		Title:   "title",
		Content: "content",
	}

	suite.mockTodoRepo.On("GetByID", suite.ctx, todoID).Return(mockTodo, nil)

	// execute
	_, err := suite.useCase.GetTodoByID(suite.ctx, todoID)

	suite.Nil(err)
}

func (suite *TestSuite) TestGetTodoByID_NotFound() {
	var todoID int64 = 1

	suite.mockTodoRepo.On("GetByID", suite.ctx, todoID).Return(nil, gorm.ErrRecordNotFound)

	// execute
	_, err := suite.useCase.GetTodoByID(suite.ctx, todoID)

	expectErr := myerror.ErrNotFound()
	myErr := err.(myerror.MyError)

	suite.Equal(expectErr.ErrorCode, myErr.ErrorCode)
	suite.Equal(expectErr.HTTPCode, myErr.HTTPCode)
}

func (suite *TestSuite) TestGetTodoByID_ErrGet() {
	var todoID int64 = 1

	suite.mockTodoRepo.On("GetByID", suite.ctx, todoID).Return(nil, errors.New("error"))

	// execute
	_, err := suite.useCase.GetTodoByID(suite.ctx, todoID)

	expectErr := myerror.ErrGet(err)
	myErr := err.(myerror.MyError)

	suite.Equal(expectErr.ErrorCode, myErr.ErrorCode)
	suite.Equal(expectErr.HTTPCode, myErr.HTTPCode)
}

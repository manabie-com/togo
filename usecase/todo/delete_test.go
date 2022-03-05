package todo_test

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/khangjig/togo/model"
	"github.com/khangjig/togo/util/myerror"
)

func (suite *TestSuite) TestDelete_Success() {
	var todoID int64 = 1

	mockTodo := &model.Todo{
		ID:      1,
		UserID:  1,
		Title:   "title",
		Content: "content",
	}

	suite.mockTodoRepo.On("GetByID", suite.ctx, todoID).Return(mockTodo, nil)
	suite.mockTodoRepo.On("DeleteByID", suite.ctx, todoID, false).Return(nil)

	// execute
	err := suite.useCase.Delete(suite.ctx, todoID)

	suite.Nil(err)
}

func (suite *TestSuite) TestDelete_NotFound() {
	var todoID int64 = 1

	suite.mockTodoRepo.On("GetByID", suite.ctx, todoID).Return(nil, gorm.ErrRecordNotFound)

	// execute
	err := suite.useCase.Delete(suite.ctx, todoID)

	suite.Nil(err)
}

func (suite *TestSuite) TestDelete_ErrGet() {
	var todoID int64 = 1

	suite.mockTodoRepo.On("GetByID", suite.ctx, todoID).Return(nil, errors.New("error"))

	// execute
	err := suite.useCase.Delete(suite.ctx, todoID)

	expectErr := myerror.ErrGet(err)
	myErr := err.(myerror.MyError)

	suite.Equal(expectErr.ErrorCode, myErr.ErrorCode)
	suite.Equal(expectErr.HTTPCode, myErr.HTTPCode)
}

func (suite *TestSuite) TestDelete_NotOwned() {
	var todoID int64 = 1

	mockTodo := &model.Todo{
		ID:      1,
		UserID:  2,
		Title:   "title",
		Content: "content",
	}

	suite.mockTodoRepo.On("GetByID", suite.ctx, todoID).Return(mockTodo, nil)

	// execute
	err := suite.useCase.Delete(suite.ctx, todoID)

	suite.Nil(err)
}

func (suite *TestSuite) TestDelete_ErrDelete() {
	var todoID int64 = 1

	mockTodo := &model.Todo{
		ID:      1,
		UserID:  1,
		Title:   "title",
		Content: "content",
	}

	suite.mockTodoRepo.On("GetByID", suite.ctx, todoID).Return(mockTodo, nil)
	suite.mockTodoRepo.On("DeleteByID", suite.ctx, todoID, false).Return(errors.New("error"))

	// execute
	err := suite.useCase.Delete(suite.ctx, todoID)

	expectErr := myerror.ErrDelete(err)
	myErr := err.(myerror.MyError)

	suite.Equal(expectErr.ErrorCode, myErr.ErrorCode)
	suite.Equal(expectErr.HTTPCode, myErr.HTTPCode)
}

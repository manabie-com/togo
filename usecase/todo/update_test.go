package todo_test

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/khangjig/togo/codetype"
	"github.com/khangjig/togo/model"
	"github.com/khangjig/togo/usecase/todo"
	"github.com/khangjig/togo/util"
	"github.com/khangjig/togo/util/myerror"
)

func (suite *TestSuite) TestUpdate_Success_NotChangedData() {
	mockReq := &todo.UpdateRequest{
		ID:      1,
		Title:   util.String("title"),
		Content: util.String("content"),
		Status:  codetype.TodoStatusOpen.GetPointer(),
	}

	mockTodo := &model.Todo{
		ID:      1,
		UserID:  1,
		Title:   "title",
		Content: "content",
		Status:  codetype.TodoStatusOpen,
	}

	suite.mockTodoRepo.On("GetByID", suite.ctx, mockReq.ID).Return(mockTodo, nil)

	// execute
	_, err := suite.useCase.Update(suite.ctx, mockReq)

	suite.Nil(err)
}

func (suite *TestSuite) TestUpdate_Success_ChangedData() {
	mockReq := &todo.UpdateRequest{
		ID:      1,
		Title:   util.String("title has changed"),
		Content: util.String("content has changed"),
		Status:  codetype.TodoStatusInProgress.GetPointer(),
	}

	mockTodo := &model.Todo{
		ID:      1,
		UserID:  1,
		Title:   "title",
		Content: "content",
		Status:  codetype.TodoStatusOpen,
	}

	suite.mockTodoRepo.On("GetByID", suite.ctx, mockReq.ID).Return(mockTodo, nil)
	suite.mockTodoRepo.On("Update", suite.ctx, mockTodo).Return(nil)

	// execute
	_, err := suite.useCase.Update(suite.ctx, mockReq)

	suite.Nil(err)
}

func (suite *TestSuite) TestUpdate_InvalidTitle() {
	mockReq := &todo.UpdateRequest{
		ID:      1,
		Title:   util.String(""),
		Content: util.String("content"),
		Status:  codetype.TodoStatusInProgress.GetPointer(),
	}

	mockTodo := &model.Todo{
		ID:      1,
		UserID:  1,
		Title:   "title",
		Content: "content",
		Status:  codetype.TodoStatusOpen,
	}

	suite.mockTodoRepo.On("GetByID", suite.ctx, mockReq.ID).Return(mockTodo, nil)

	// execute
	_, err := suite.useCase.Update(suite.ctx, mockReq)

	expectErr := myerror.ErrTodoTitleInvalid("Invalid title.")
	myErr := err.(myerror.MyError)

	suite.Equal(expectErr.ErrorCode, myErr.ErrorCode)
	suite.Equal(expectErr.HTTPCode, myErr.HTTPCode)
}

func (suite *TestSuite) TestUpdate_InvalidStatus() {
	invalidStatus := codetype.TodoStatus(7)

	mockReq := &todo.UpdateRequest{
		ID:      1,
		Title:   util.String("title"),
		Content: util.String("content"),
		Status:  &invalidStatus,
	}

	mockTodo := &model.Todo{
		ID:      1,
		UserID:  1,
		Title:   "title",
		Content: "content",
		Status:  codetype.TodoStatusOpen,
	}

	suite.mockTodoRepo.On("GetByID", suite.ctx, mockReq.ID).Return(mockTodo, nil)

	// execute
	_, err := suite.useCase.Update(suite.ctx, mockReq)

	expectErr := myerror.ErrTodoStatusInvalid()
	myErr := err.(myerror.MyError)

	suite.Equal(expectErr.ErrorCode, myErr.ErrorCode)
	suite.Equal(expectErr.HTTPCode, myErr.HTTPCode)
}

// nolint:lll,misspell
func (suite *TestSuite) TestUpdate_TooLongTitle() {
	mockReq := &todo.UpdateRequest{
		ID:      1,
		Title:   util.String("Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor. Aenean massa. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus.Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor. Aenean massa. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus."),
		Content: util.String("content"),
		Status:  codetype.TodoStatusInProgress.GetPointer(),
	}

	mockTodo := &model.Todo{
		ID:      1,
		UserID:  1,
		Title:   "title",
		Content: "content",
		Status:  codetype.TodoStatusOpen,
	}

	suite.mockTodoRepo.On("GetByID", suite.ctx, mockReq.ID).Return(mockTodo, nil)

	// execute
	_, err := suite.useCase.Update(suite.ctx, mockReq)

	expectErr := myerror.ErrTodoTitleInvalid("Invalid title.")
	myErr := err.(myerror.MyError)

	suite.Equal(expectErr.ErrorCode, myErr.ErrorCode)
	suite.Equal(expectErr.HTTPCode, myErr.HTTPCode)
}

func (suite *TestSuite) TestUpdate_NotFound() {
	mockReq := &todo.UpdateRequest{
		ID:      1,
		Title:   util.String("title"),
		Content: util.String("content"),
		Status:  codetype.TodoStatusInProgress.GetPointer(),
	}

	suite.mockTodoRepo.On("GetByID", suite.ctx, mockReq.ID).Return(nil, gorm.ErrRecordNotFound)

	// execute
	_, err := suite.useCase.Update(suite.ctx, mockReq)

	expectErr := myerror.ErrNotFound()
	myErr := err.(myerror.MyError)

	suite.Equal(expectErr.ErrorCode, myErr.ErrorCode)
	suite.Equal(expectErr.HTTPCode, myErr.HTTPCode)
}

func (suite *TestSuite) TestUpdate_ErrGet() {
	mockReq := &todo.UpdateRequest{
		ID:      1,
		Title:   util.String("title"),
		Content: util.String("content"),
		Status:  codetype.TodoStatusInProgress.GetPointer(),
	}

	suite.mockTodoRepo.On("GetByID", suite.ctx, mockReq.ID).Return(nil, errors.New("error"))

	// execute
	_, err := suite.useCase.Update(suite.ctx, mockReq)

	expectErr := myerror.ErrGet(err)
	myErr := err.(myerror.MyError)

	suite.Equal(expectErr.ErrorCode, myErr.ErrorCode)
	suite.Equal(expectErr.HTTPCode, myErr.HTTPCode)
}

func (suite *TestSuite) TestUpdate_NotOwned() {
	mockReq := &todo.UpdateRequest{
		ID:      1,
		Title:   util.String("title"),
		Content: util.String("content"),
		Status:  codetype.TodoStatusInProgress.GetPointer(),
	}

	mockTodo := &model.Todo{
		ID:      1,
		UserID:  2,
		Title:   "title",
		Content: "content",
		Status:  codetype.TodoStatusOpen,
	}

	suite.mockTodoRepo.On("GetByID", suite.ctx, mockReq.ID).Return(mockTodo, nil)

	// execute
	_, err := suite.useCase.Update(suite.ctx, mockReq)

	expectErr := myerror.ErrNotFound()
	myErr := err.(myerror.MyError)

	suite.Equal(expectErr.ErrorCode, myErr.ErrorCode)
	suite.Equal(expectErr.HTTPCode, myErr.HTTPCode)
}

func (suite *TestSuite) TestUpdate_ErrUpdate() {
	mockReq := &todo.UpdateRequest{
		ID:      1,
		Title:   util.String("title"),
		Content: util.String("content"),
		Status:  codetype.TodoStatusInProgress.GetPointer(),
	}

	mockTodo := &model.Todo{
		ID:      1,
		UserID:  1,
		Title:   "title",
		Content: "content",
		Status:  codetype.TodoStatusOpen,
	}

	suite.mockTodoRepo.On("GetByID", suite.ctx, mockReq.ID).Return(mockTodo, nil)
	suite.mockTodoRepo.On("Update", suite.ctx, mockTodo).Return(errors.New("error"))

	// execute
	_, err := suite.useCase.Update(suite.ctx, mockReq)

	expectErr := myerror.ErrUpdate(err)
	myErr := err.(myerror.MyError)

	suite.Equal(expectErr.ErrorCode, myErr.ErrorCode)
	suite.Equal(expectErr.HTTPCode, myErr.HTTPCode)
}

package todo_test

import (
	"github.com/pkg/errors"

	"github.com/khangjig/togo/codetype"
	"github.com/khangjig/togo/model"
	"github.com/khangjig/togo/usecase/todo"
	"github.com/khangjig/togo/util/myerror"
)

// nolint:lll
func (suite *TestSuite) TestGetList_Success() {
	mockReq := &todo.GetListRequest{
		GetListRequest: codetype.GetListRequest{
			Paginator: codetype.Paginator{
				Page:  1,
				Limit: codetype.PageSizeDefault,
			},
			OrderBy: "title",
		},
		Status: codetype.TodoStatusInProgress.GetPointer(),
	}

	mockTodos := []model.Todo{
		{
			ID:      1,
			UserID:  1,
			Title:   "title",
			Content: "content",
		},
	}

	suite.mockTodoRepo.On("GetList", suite.ctx, suite.userClaims.ID, map[string]interface{}{"status": *mockReq.Status}, mockReq.Search, "title ASC", mockReq.Paginator).Return(mockTodos, int64(1), nil)

	// execute
	_, err := suite.useCase.GetList(suite.ctx, mockReq)

	suite.Nil(err)
}

// nolint:lll
func (suite *TestSuite) TestGetList_ErrGet() {
	progress := codetype.TodoStatusInProgress

	mockReq := &todo.GetListRequest{
		GetListRequest: codetype.GetListRequest{
			Paginator: codetype.Paginator{
				Page:  1,
				Limit: codetype.PageSizeDefault,
			},
			OrderBy: "title",
		},
		Status: &progress,
	}

	suite.mockTodoRepo.On("GetList", suite.ctx, suite.userClaims.ID, map[string]interface{}{"status": *mockReq.Status}, mockReq.Search, "title ASC", mockReq.Paginator).Return(nil, int64(0), errors.New("error"))

	// execute
	_, err := suite.useCase.GetList(suite.ctx, mockReq)

	expectErr := myerror.ErrGet(err)
	myErr := err.(myerror.MyError)

	suite.Equal(expectErr.ErrorCode, myErr.ErrorCode)
	suite.Equal(expectErr.HTTPCode, myErr.HTTPCode)
}

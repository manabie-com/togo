package todo

import (
	"context"
	"strings"
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/khangjig/togo/codetype"
	"github.com/khangjig/togo/model"
	"github.com/khangjig/togo/util"
	"github.com/khangjig/togo/util/jwt"
	"github.com/khangjig/togo/util/myerror"
)

type UpdateRequest struct {
	ID      int64                `json:"-"`
	Title   *string              `json:"title"`
	Content *string              `json:"content"`
	Status  *codetype.TodoStatus `json:"status"`
}

func (u *UseCase) validateUpdate(myTodo *model.Todo, req *UpdateRequest) (bool, error) {
	isChanged := false

	if req.Title != nil {
		*req.Title = strings.TrimSpace(*req.Title)
		if len(*req.Title) == 0 {
			return false, myerror.ErrTodoTitleInvalid("Invalid title.")
		}

		if len(*req.Title) > 128 {
			return false, myerror.ErrTodoTitleInvalid("Title must be less than 128 characters.")
		}

		if myTodo.Title != *req.Title {
			myTodo.Title = *req.Title
			isChanged = true
		}
	}

	if req.Content != nil {
		*req.Content = strings.TrimSpace(*req.Content)

		if len(*req.Content) > 65535 {
			return false, myerror.ErrTodoContentInvalid("Content is too long.")
		}

		if myTodo.Content != *req.Content {
			myTodo.Content = *req.Content
			isChanged = true
		}
	}

	if req.Status != nil {
		if !req.Status.IsValid() {
			return false, myerror.ErrTodoStatusInvalid()
		}

		if myTodo.Status != *req.Status {
			myTodo.Status = *req.Status
			isChanged = true
		}
	}

	return isChanged, nil
}

func (u *UseCase) Update(ctx context.Context, req *UpdateRequest) (*ResponseWrapper, error) {
	myUser, _ := ctx.Value(jwt.MyUserClaim).(*model.User)

	myTodo, err := u.TodoRepo.GetByID(ctx, req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, myerror.ErrNotFound()
		}

		return nil, myerror.ErrGet(err)
	}

	if myTodo.UserID != myUser.ID {
		return nil, myerror.ErrNotFound()
	}

	isChanged, err := u.validateUpdate(myTodo, req)
	if err != nil {
		return nil, err
	}

	if isChanged {
		myTodo.EditedAt = util.Time(time.Now())

		err = u.TodoRepo.Update(ctx, myTodo)
		if err != nil {
			return nil, myerror.ErrUpdate(err)
		}
	}

	return &ResponseWrapper{Todo: myTodo}, nil
}

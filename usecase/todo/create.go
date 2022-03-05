package todo

import (
	"context"
	"strings"

	"github.com/khangjig/togo/codetype"
	"github.com/khangjig/togo/model"
	"github.com/khangjig/togo/util/jwt"
	"github.com/khangjig/togo/util/myerror"
)

type CreateRequest struct {
	Title   string              `json:"title"`
	Content string              `json:"content"`
	Status  codetype.TodoStatus `json:"status"`
}

func (u *UseCase) validateCreate(req *CreateRequest) error {
	req.Title = strings.TrimSpace(req.Title)

	if len(req.Title) == 0 {
		return myerror.ErrTodoTitleInvalid("Invalid title.")
	}

	if len(req.Title) > 128 {
		return myerror.ErrTodoTitleInvalid("Title must be less than 128 characters.")
	}

	req.Content = strings.TrimSpace(req.Content)

	if len(req.Content) > 65535 {
		return myerror.ErrTodoContentInvalid("Content is too long.")
	}

	if !req.Status.IsValid() {
		return myerror.ErrTodoStatusInvalid()
	}

	return nil
}

func (u *UseCase) Create(ctx context.Context, req *CreateRequest) (*ResponseWrapper, error) {
	myUser, _ := ctx.Value(jwt.MyUserClaim).(*model.User)

	err := u.validateCreate(req)
	if err != nil {
		return nil, err
	}

	myTodo := model.Todo{
		UserID:  myUser.ID,
		Title:   req.Title,
		Content: req.Content,
		Status:  req.Status,
	}

	err = u.TodoRepo.Create(ctx, &myTodo)
	if err != nil {
		return nil, myerror.ErrCreate(err)
	}

	return &ResponseWrapper{Todo: &myTodo}, nil
}

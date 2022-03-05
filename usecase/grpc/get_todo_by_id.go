package grpc

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/khangjig/togo/model"
	"github.com/khangjig/togo/util/myerror"
)

func (u *UseCase) GetTodoByID(ctx context.Context, todoID int64) (*model.Todo, error) {
	myTodo, err := u.TodoRepo.GetByID(ctx, todoID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, myerror.ErrNotFound()
		}

		return nil, myerror.ErrGet(err)
	}

	return myTodo, nil
}

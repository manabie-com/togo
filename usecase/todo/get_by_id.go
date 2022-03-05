package todo

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/khangjig/togo/model"
	"github.com/khangjig/togo/util/jwt"
	"github.com/khangjig/togo/util/myerror"
)

func (u *UseCase) GetByID(ctx context.Context, id int64) (*ResponseWrapper, error) {
	myUser, _ := ctx.Value(jwt.MyUserClaim).(*model.User)

	myTodo, err := u.TodoRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, myerror.ErrNotFound()
		}

		return nil, myerror.ErrGet(err)
	}

	if myTodo.UserID != myUser.ID {
		return nil, myerror.ErrNotFound()
	}

	return &ResponseWrapper{Todo: myTodo}, nil
}

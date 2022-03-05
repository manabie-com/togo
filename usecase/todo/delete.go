package todo

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/khangjig/togo/model"
	"github.com/khangjig/togo/util/jwt"
	"github.com/khangjig/togo/util/myerror"
)

func (u *UseCase) Delete(ctx context.Context, id int64) error {
	myUser, _ := ctx.Value(jwt.MyUserClaim).(*model.User)

	myTodo, err := u.TodoRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}

		return myerror.ErrGet(err)
	}

	if myTodo.UserID != myUser.ID {
		return nil
	}

	err = u.TodoRepo.DeleteByID(ctx, myTodo.ID, false)
	if err != nil {
		return myerror.ErrDelete(err)
	}

	return nil
}

package user

import (
	"context"

	"github.com/khangjig/togo/model"
	"github.com/khangjig/togo/util/jwt"
)

func (u *UseCase) GetMe(ctx context.Context) (*ResponseWrapper, error) {
	myUser, _ := ctx.Value(jwt.MyUserClaim).(*model.User)

	return &ResponseWrapper{User: myUser}, nil
}

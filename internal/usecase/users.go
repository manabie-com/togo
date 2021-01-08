package usecase

import (
	"context"
	"github.com/manabie-com/togo/internal/model"
)

func (a Usecase) GetUser(ctx context.Context, id string) (*model.User, error) {
	res, err := a.Store.User().Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a Usecase) AuthenticateUser(ctx context.Context, id, password string) (*model.User, error) {
	u, err := a.GetUser(ctx, id)
	if err != nil {
		return nil, model.NewError(model.ErrAuthenticateUser, err.Error())
	}

	if !u.ComparePassword(password) {
		return nil, model.NewError(model.ErrAuthenticateUser, "")
	}

	return u, nil
}

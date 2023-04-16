package userrepo

import (
	"context"

	"github.com/phathdt/libs/go-sdk/sdkcm"
	"togo/modules/user/usermodel"
)

type UserStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}) (*usermodel.User, error)
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
}

type repo struct {
	store UserStorage
}

func (r *repo) GetUserLimit(ctx context.Context, userId int) (int, error) {
	user, err := r.store.FindUser(ctx, map[string]interface{}{"id": userId})
	if err != nil {
		return 0, sdkcm.ErrCannotGetEntity("user", err)
	}

	return user.LimitTask, nil
}

func NewRepo(store UserStorage) *repo {
	return &repo{store: store}
}

func (r *repo) CreateUser(ctx context.Context, data *usermodel.UserCreate) error {
	if err := r.store.CreateUser(ctx, data); err != nil {
		return sdkcm.ErrCannotCreateEntity("user", err)
	}

	return nil
}

func (r *repo) FindUser(ctx context.Context, conditions map[string]interface{}) (*usermodel.User, error) {
	user, err := r.store.FindUser(ctx, conditions)
	if err != nil {
		return nil, sdkcm.ErrCannotCreateEntity("user", err)
	}

	return user, nil
}

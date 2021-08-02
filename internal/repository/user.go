package repository

import (
	"context"
	"togo/internal/entity"
	"togo/internal/postgresql"
	"togo/utils"
)

func (r *Repo) CreateUser(ctx context.Context, username string, password string) (entity.User, error) {
	user, err := r.q.InsertUser(ctx, postgresql.InsertUserParams{
		Username: username,
		Password: utils.GetHash([]byte(password)),
	})

	if err != nil {
		return entity.User{}, err
	}

	return user.MapToEntity(), nil
}

func (r *Repo) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	user, err := r.q.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	userEntity := user.MapToEntity()
	return &userEntity, nil
}

func (r *Repo) GetUser(ctx context.Context, id int32) (entity.User, error) {
	user, err := r.q.GetUser(ctx, id)
	if err != nil {
		return entity.User{}, err
	}

	return user.MapToEntity(), nil
}

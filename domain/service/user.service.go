package service

import (
	"context"
	"errors"
	"togo/domain/model"
	"togo/domain/repository"
)

type UserService interface {
	Register(ctx context.Context, user model.User) error
	Login(ctx context.Context, username string, password string) (model.User, error)
}

type userService struct {
	db repository.UserRepository
}

func (this userService) Register(ctx context.Context, user model.User) error {
	err := this.db.Create(user)
	return err
}

func (this *userService) createToken(ctx context.Context, user *model.User) (string, error) {
	// TODO: create JWT token and saving it in database
	return "token", nil
}

func (this userService) Login(ctx context.Context, username string, password string) (model.User, error) {
	user, err := this.db.Get(username)
	if err != nil {
		return model.User{}, err
	}

	if user.Password != password {
		return model.User{}, errors.New("username or password is incorrect")
	}
	// create token for this session
	token, err := this.createToken(ctx, &user)
	user.Token = token
	return user, nil
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{db: userRepo}
}

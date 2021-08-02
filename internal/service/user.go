package service

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"togo/common/cmerrors"
	"togo/internal/entity"
)

type UserRepository interface {
	CreateUser(ctx context.Context, username string, password string) (entity.User, error)
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
	GetUser(ctx context.Context, id int32) (entity.User, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) CreateUser(ctx context.Context, username string, password string) (entity.User, error) {
	if _, err := u.repo.GetUserByUsername(ctx, username); err == nil {
		return entity.User{}, cmerrors.ErrUserAlreadyExist
	}

	user, err := u.repo.CreateUser(ctx, username, password)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (u *UserService) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	user, err := u.repo.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, cmerrors.ErrUserNotFound
	}

	return user, nil
}

func (u *UserService) Login(ctx context.Context, username string, password string) (*entity.User, error) {
	user, err := u.repo.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, cmerrors.ErrUserNotFound
	}

	userPass := []byte(password)
	dbPass := []byte(user.Password)

	if passErr := bcrypt.CompareHashAndPassword(dbPass, userPass); passErr != nil {
		return nil, cmerrors.ErrPasswordNotMatch
	}

	return user, nil
}

func (u *UserService) GetUser(ctx context.Context, id int32) (entity.User, error) {
	user, err := u.repo.GetUser(ctx, id)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

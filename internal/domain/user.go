package domain

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

type UserRedisRepo interface {
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
	GetUser(ctx context.Context, id int32) (*entity.User, error)
	SetUser(ctx context.Context, user entity.User) error
}
type UserDomain struct {
	repo    UserRepository
	rdbRepo UserRedisRepo
}

func NewUserDomain(repo UserRepository, rdbRepo UserRedisRepo) *UserDomain {
	return &UserDomain{repo: repo, rdbRepo: rdbRepo}
}

func (u *UserDomain) CreateUser(ctx context.Context, username string, password string) (*entity.User, error) {
	userRdb, err := u.rdbRepo.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	if userRdb != nil {
		return nil, cmerrors.ErrUserAlreadyExist
	}

	if _, err = u.repo.GetUserByUsername(ctx, username); err == nil {
		return nil, cmerrors.ErrUserAlreadyExist
	}

	user, err := u.repo.CreateUser(ctx, username, password)
	if err != nil {
		return nil, err
	}

	_ = u.rdbRepo.SetUser(ctx, user)

	return &user, nil
}

func (u *UserDomain) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	userRdb, err := u.rdbRepo.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	if userRdb != nil {
		return userRdb, nil
	}

	user, err := u.repo.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, cmerrors.ErrUserNotFound
	}

	_ = u.rdbRepo.SetUser(ctx, *user)

	return user, nil
}

func (u *UserDomain) Login(ctx context.Context, username string, password string) (*entity.User, error) {
	var user *entity.User

	userRdb, err := u.rdbRepo.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	if userRdb != nil {
		user = userRdb
	}

	if user == nil {
		userDb, err := u.repo.GetUserByUsername(ctx, username)
		if err != nil {
			return nil, cmerrors.ErrUserNotFound
		}

		_ = u.rdbRepo.SetUser(ctx, *userDb)

		user = userDb
	}

	userPass := []byte(password)
	dbPass := []byte(user.Password)

	if passErr := bcrypt.CompareHashAndPassword(dbPass, userPass); passErr != nil {
		return nil, cmerrors.ErrPasswordNotMatch
	}

	return user, nil
}

func (u *UserDomain) GetUser(ctx context.Context, id int32) (*entity.User, error) {
	userRdb, err := u.rdbRepo.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	if userRdb != nil {
		return userRdb, nil
	}

	user, err := u.repo.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	_ = u.rdbRepo.SetUser(ctx, user)

	return &user, nil
}

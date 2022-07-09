package service

import (
	"context"
	"fmt"

	"github.com/lawtrann/togo"
)

type UserService struct {
	Repo togo.UserRepo
}

func NewUserService(rp togo.UserRepo) *UserService {
	return &UserService{Repo: rp}
}

func (us *UserService) GetUserByName(ctx context.Context, username string) (*togo.User, error) {
	result, err := us.Repo.GetUserByName(ctx, username)
	if err != nil {
		fmt.Println(err)
		return &togo.User{}, err
	}
	return result, nil
}

func (us *UserService) IsExceedPerDay(ctx context.Context, u *togo.User) (bool, error) {
	result, err := us.Repo.IsExceedPerDay(ctx, u)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	return result, nil
}

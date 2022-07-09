package service

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/lawtrann/togo"
)

type TodoService struct {
	Repo togo.TodoRepo

	UserService togo.UserService
}

func NewTodoService(rp togo.TodoRepo) *TodoService {
	return &TodoService{Repo: rp}
}

func (ts TodoService) Add(ctx context.Context, t *togo.Todo, username string) (*togo.Todo, error) {
	var res *togo.Todo
	var u *togo.User

	u, err := ts.UserService.GetUserByName(ctx, username)
	if err != nil {
		fmt.Println(err)
		return &togo.Todo{}, err
	}

	// checkif user exist
	if (togo.User{}) != *u {
		// checkif exceed a limited per day
		isExceedPerDay, err := ts.UserService.IsExceedPerDay(ctx, u)
		if isExceedPerDay || err != nil {
			return &togo.Todo{}, togo.ErrIsExceedLimitedPerDay
		}
		res, err = ts.Repo.Add(ctx, t, u)
		if err != nil {
			fmt.Println(err)
			return &togo.Todo{}, err
		}
	} else {
		// set user's name and limited per day
		u.Username = username
		u.LimitedPerDay = rand.Intn(10) + 1
		res, err = ts.Repo.AddWithNewUser(ctx, t, u)
		if err != nil {
			fmt.Println(err)
			return &togo.Todo{}, err
		}
	}

	return res, nil
}

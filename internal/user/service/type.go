package service

import (
	"time"

	"github.com/manabie-com/togo/model"
)

type User struct {
	ID        int
	TaskLimit int
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateUserArgs struct {
	Username  string
	Password  string
	TaskLimit int
}

type UpdateUserArgs struct {
	UserID    int
	Password  string
	TaskLimit *int
}

type LoginUserArgs struct {
	UserID   int
	Password string
	Username string
}

type LoginUserResponse struct {
	AccessToken string
	AtExpires   int64
}

type GetUserArgs struct {
	UserID int
}

func convertModelUserToServiceUser(args *model.User) *User {
	if args == nil {
		return nil
	}
	return &User{
		ID:        args.ID,
		TaskLimit: args.TaskLimit,
		Password:  args.Password,
		CreatedAt: args.CreatedAt,
		UpdatedAt: args.UpdatedAt,
	}
}

func convertModelTasksToServiceTasks(args []*model.User) []*User {
	var res []*User
	for _, v := range args {
		res = append(res, convertModelUserToServiceUser(v))
	}
	return res
}

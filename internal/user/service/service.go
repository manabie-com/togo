package service

import (
	"context"
	"errors"

	"github.com/manabie-com/togo/pkg/errorx"

	"github.com/manabie-com/togo/pkg/auth"

	"github.com/manabie-com/togo/model"

	taskrepository "github.com/manabie-com/togo/internal/task/repository"
	"github.com/manabie-com/togo/internal/user/repository"
	"gorm.io/gorm"
)

type UserService interface {
	GetUser(context.Context, *GetUserArgs) (*User, error)
	CreateUser(context.Context, *CreateUserArgs) error
	UpdateUser(context.Context, *UpdateUserArgs) error
	Login(context.Context, *LoginUserArgs) (*LoginUserResponse, error)
	Authenticate(ctx context.Context, userID int) (*model.User, error)
}

type userService struct {
	db       *gorm.DB
	userRepo repository.UserRepository
	taskRepo taskrepository.TaskRepository
}

func (u *userService) GetUser(ctx context.Context, args *GetUserArgs) (*User, error) {
	task, err := u.userRepo.GetUser(ctx, &model.User{
		ID: args.UserID,
	})
	if err != nil {
		return nil, err
	}
	return convertModelUserToServiceUser(task), nil
}

func NewUserService(userRepo repository.UserRepository, taskRepo taskrepository.TaskRepository, db *gorm.DB) UserService {
	return &userService{
		db:       db,
		userRepo: userRepo,
		taskRepo: taskRepo,
	}
}

func (u *userService) Login(ctx context.Context, args *LoginUserArgs) (*LoginUserResponse, error) {
	user, err := u.userRepo.GetUser(ctx, &model.User{
		ID: args.UserID,
	})
	if err != nil {
		return nil, err
	}
	isMatch := auth.CheckPasswordHashMatch(args.Password, user.Password)
	if !isMatch {
		return nil, errorx.ErrAuthFailure(errors.New("Mật khẩu không hợp lệ"))
	}
	tokenDetail, err := auth.GenerateToken(user.ID)
	if err != nil {
		return nil, errorx.ErrAuthFailure(err)
	}
	return &LoginUserResponse{
		AccessToken: tokenDetail.AccessToken,
		AtExpires:   tokenDetail.AtExpires,
	}, nil
}

func (u *userService) CreateUser(ctx context.Context, args *CreateUserArgs) error {
	hashedPassword, err := auth.HashPassword(args.Password)
	if err != nil {
		return errorx.ErrInternal(err)
	}
	tx := u.db.Begin()
	if err = u.userRepo.SaveUser(tx, &model.User{
		Password:  hashedPassword,
		LimitTask: args.LimitTask,
	}); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (u *userService) UpdateUser(ctx context.Context, args *UpdateUserArgs) error {
	user, err := u.userRepo.GetUser(ctx, &model.User{
		ID: args.UserID,
	})
	if err != nil {
		return err
	}
	if args.TaskLimit != nil {
		taskCount, err := u.taskRepo.CountByUserID(ctx, user.ID)
		if err != nil {
			return err
		}
		if *args.TaskLimit < int(taskCount) {
			return errorx.ErrInvalidParameter(errors.New("New task limit  must be bigger current task limit "))
		}
		user.LimitTask = *args.TaskLimit
	}

	if args.Password != "" {
		hashedPassword, err := auth.HashPassword(args.Password)
		if err != nil {
			return errorx.ErrInternal(err)
		}
		user.Password = hashedPassword
	}

	tx := u.db.Begin()
	if err = u.userRepo.SaveUser(tx, user); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (u *userService) Authenticate(ctx context.Context, userID int) (*model.User, error) {
	return u.userRepo.GetUser(ctx, &model.User{
		ID: userID,
	})
}

package repo

import (
	"github.com/jinzhu/gorm"
	"github.com/manabie-com/togo/common/context"
	"github.com/manabie-com/togo/domain/model"
)

type UserRepository interface {
	GetUserByUsername(ctx context.Context, userName string) (*model.User, error)
	GetUserById(ctx context.Context, id string) (*model.User, error)
}

type user struct {
	Id       string
	Username string
	Password string
	MaxTodo  int
}

func (u *user) TableName() string {
	return "users"
}

func NewUserRepo() UserRepository {
	return &userRepo{}
}

type userRepo struct {
}

func (u *userRepo) GetUserById(ctx context.Context, id string) (*model.User, error) {
	db := ctx.GetDb()
	userGorm := &user{}
	err := db.First(userGorm, "id = ?", id).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return mapUserGormToModel(userGorm), nil
}

func (u *userRepo) GetUserByUsername(ctx context.Context, userName string) (*model.User, error) {
	db := ctx.GetDb()
	userGorm := &user{}
	err := db.First(userGorm, "username = ?", userName).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return mapUserGormToModel(userGorm), nil
}

func mapUserGormToModel(userGorm *user) *model.User {
	return &model.User{
		Id:       userGorm.Id,
		Username: userGorm.Username,
		Password: userGorm.Password,
		MaxTodo:  userGorm.MaxTodo,
	}
}

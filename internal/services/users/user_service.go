package users

import (
	"context"
	"database/sql"

	"github.com/manabie-com/togo/internal/consts"
	"github.com/manabie-com/togo/internal/models"
)

func NewService(db *sql.DB) IUserService {
	return &UserService{
		DB: db,
	}
}

type IUserService interface {
	Validate(context.Context, string, string) error
	GetUserByID(context.Context, string) (*models.User, error)
}

type UserService struct {
	DB *sql.DB
}

func (s UserService) Validate(context context.Context, userId string, password string) error {
	user, err := models.Users(
		models.UserWhere.ID.EQ(userId),
		models.UserWhere.Password.EQ(password),
	).One(context, s.DB)
	if err != nil || user == nil {
		if err == sql.ErrNoRows {
			return consts.ErrInvalidAuth
		}
		return err
	}

	return nil
}

func (s UserService) GetUserByID(ctx context.Context, userID string) (*models.User, error) {
	return models.Users(
		models.UserWhere.ID.EQ(userID),
	).One(ctx, s.DB)

}

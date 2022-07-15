package user

import (
	"context"
	"database/sql"

	"manabie/todo/models"
	"manabie/todo/pkg/db"
	"manabie/todo/repository/user"
)

type UserService interface {
	Index(ctx context.Context) ([]*models.User, error)
}

type service struct {
	User user.UserRespository
}

func NewUserService(ur user.UserRespository) UserService {
	return &service{
		User: ur,
	}
}

func (s *service) Index(ctx context.Context) (users []*models.User, err error) {
	if err := db.Transaction(ctx, nil, func(ctx context.Context, tx *sql.Tx) error {

		users, err = s.User.Find(ctx, tx)

		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return users, nil
}

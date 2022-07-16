package user

import (
	"context"
	"database/sql"

	"manabie/todo/models"
)

const (
	queryFind = `SELECT * FROM member`
)

type UserRespository interface {
	Find(ctx context.Context, tx *sql.Tx) ([]*models.User, error)
}

type userRespository struct{}

func NewUserRespository() UserRespository {
	return &userRespository{}
}

func (ur *userRespository) Find(ctx context.Context, tx *sql.Tx) ([]*models.User, error) {
	rows, err := tx.QueryContext(ctx, queryFind)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()

	users := make([]*models.User, 0)

	for rows.Next() {

		u := &models.User{}

		if err := rows.Scan(&u.ID, &u.Email, &u.Name, &u.CreatedAt, &u.UpdateAt); err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}

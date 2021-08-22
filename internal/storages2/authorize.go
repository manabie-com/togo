package storages2

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type IAuthorizeRepo interface {
	ValidateUser(ctx context.Context, arg ValidateUserParams) (string, error)
}

type AuthorizeRepo struct {
	*Queries
	db *sqlx.DB
}

func NewAuthorizeRepo(db *sqlx.DB) IAuthorizeRepo {
	return &AuthorizeRepo{
		db: db,
	}
}

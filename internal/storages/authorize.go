package storages

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/manabie-com/togo/internal/tools"
	"net/http"
)

type IAuthorizeRepo interface {
	ValidateUserStore(ctx context.Context, arg ValidateUserParams) (string, *tools.TodoError)
}

type AuthorizeRepo struct {
	*Queries
	db *sqlx.DB
}

func (ar *AuthorizeRepo) ValidateUserStore(ctx context.Context, arg ValidateUserParams) (string, *tools.TodoError) {
	id, err := ar.ValidateUser(ctx, arg)
	if err != nil {
		return "", tools.NewTodoError(http.StatusInternalServerError, err.Error())
	}
	return id, nil
}

func NewAuthorizeRepo(db *sqlx.DB) IAuthorizeRepo {
	return &AuthorizeRepo{
		Queries: New(db),
		db:      db,
	}
}

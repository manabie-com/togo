package storages

import (
	"context"
	"github.com/manabie-com/togo/internal/tools"
	"net/http"
)

type IAuthorizeRepo interface {
	ValidateUserStore(ctx context.Context, arg ValidateUserParams) (string, *tools.TodoError)
}

type AuthorizeRepo struct {
	*Queries
}

func (ar *AuthorizeRepo) ValidateUserStore(ctx context.Context, arg ValidateUserParams) (string, *tools.TodoError) {
	_, err := ar.ValidateUser(ctx, arg)
	if err != nil {
		return "", tools.NewTodoError(http.StatusInternalServerError, err.Error())
	}
	return arg.ID, nil
}

func NewAuthorizeRepo(db DBTX) IAuthorizeRepo {
	return &AuthorizeRepo{
		Queries: New(db),
	}
}

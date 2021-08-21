package api

import (
	"context"
	"database/sql"
	"github.com/manabie-com/togo/internal/dto"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages/repos"
	"github.com/manabie-com/togo/internal/tools"
	"net/http"
)

type AuthorApi struct {
	service dto.IAuthorizeApi
}

func (aa *AuthorApi) Login(ctx context.Context, req *http.Request) (*dto.LoginResponse, error) {
	id := tools.Value(req, "user_id")
	password := tools.Value(req, "password")
	if !id.Valid || !password.Valid {
		return nil, tools.NewTodoError(400, "bad request")
	}
	return aa.service.Login(ctx, dto.LoginRequest{UserId: id.String, Password: password.String})
}

func (aa *AuthorApi) Validate(req *http.Request) (context.Context, error) {
	return aa.service.Validate(req)
}

func NewAuthorApi(db *sql.DB, JWTKey string) AuthorApi {
	return AuthorApi{
		service: services.NewAuthorizeService(repos.NewAuthorizeRepo(db), JWTKey),
	}
}

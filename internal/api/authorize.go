package api

import (
	"context"
	"database/sql"
	"github.com/manabie-com/togo/internal/iservices"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages/repos"
	"github.com/manabie-com/togo/internal/tools"
	"net/http"
)

type IAuthorApi interface {
	Login(ctx context.Context, req *http.Request) (*iservices.LoginResponse, *tools.TodoError)
	Validate(req *http.Request) (context.Context, *tools.TodoError)
}

type AuthorApi struct {
	service     iservices.IAuthorizeService
	requestTool tools.IRequestTool
}

func (aa *AuthorApi) Login(ctx context.Context, req *http.Request) (*iservices.LoginResponse, *tools.TodoError) {
	id := aa.requestTool.Value(req, "user_id")
	password := aa.requestTool.Value(req, "password")
	if !id.Valid || !password.Valid {
		return nil, tools.NewTodoError(400, "bad request")
	}
	return aa.service.Login(ctx, iservices.LoginRequest{UserId: id.String, Password: password.String})
}

func (aa *AuthorApi) Validate(req *http.Request) (context.Context, *tools.TodoError) {
	return aa.service.Validate(req)
}

func NewAuthorApi(db *sql.DB, JWTKey string) AuthorApi {
	return AuthorApi{
		service:     services.NewAuthorizeService(repos.NewAuthorizeRepo(db), JWTKey),
		requestTool: tools.NewRequestTool(),
	}
}

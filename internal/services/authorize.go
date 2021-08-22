package services

import (
	"context"
	"database/sql"
	"net/http"

	authorizeApi "github.com/manabie-com/togo/internal/iservices"
	"github.com/manabie-com/togo/internal/storages/repos"
	"github.com/manabie-com/togo/internal/tools"
)

type AuthorizeService struct {
	repo        repos.IAuthorizeRepo
	JWTKey      string
	contextTool tools.IContextTool
	tokenTool   tools.ITokenTool
}

func (as *AuthorizeService) Login(ctx context.Context, req authorizeApi.LoginRequest) (*authorizeApi.LoginResponse, *tools.TodoError) {
	if !as.repo.ValidateUser(ctx,
		sql.NullString{String: req.UserId, Valid: true},
		sql.NullString{String: req.Password, Valid: true}) {
		return nil, tools.NewTodoError(http.StatusUnauthorized, "incorrect user_id/pwd")
	}
	token, err := as.tokenTool.CreateToken(req.UserId, as.JWTKey)
	if err != nil {
		return nil, err
	}
	return &authorizeApi.LoginResponse{Data: token}, nil
}

func (as *AuthorizeService) Validate(req *http.Request) (context.Context, *tools.TodoError) {
	token := as.tokenTool.GetToken(req)
	id, err := as.tokenTool.ClaimToken(token, as.JWTKey)

	if err != nil {
		return nil, err
	}

	ctx := as.contextTool.WriteUserIDToContext(req.Context(), id)
	return ctx, nil
}

func NewAuthorizeService(repo repos.IAuthorizeRepo, jwtKey string, tokenTool tools.ITokenTool, contextTool tools.IContextTool) authorizeApi.IAuthorizeService {
	return &AuthorizeService{
		repo:        repo,
		JWTKey:      jwtKey,
		tokenTool:   tokenTool,
		contextTool: contextTool,
	}
}

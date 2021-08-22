package services

import (
	"context"
	"github.com/manabie-com/togo/internal/storages"
	"net/http"

	authorizeApi "github.com/manabie-com/togo/internal/iservices"
	"github.com/manabie-com/togo/internal/tools"
)

type AuthorizeService struct {
	repo        storages.IAuthorizeRepo
	JWTKey      string
	contextTool tools.IContextTool
	tokenTool   tools.ITokenTool
}

func (as *AuthorizeService) Login(ctx context.Context, req authorizeApi.LoginRequest) (*authorizeApi.LoginResponse, *tools.TodoError) {
	_, err := as.repo.ValidateUserStore(ctx, storages.ValidateUserParams{ID: req.UserId, Password: req.Password})
	if err != nil {
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

func NewAuthorizeService(repo storages.IAuthorizeRepo, jwtKey string, tokenTool tools.ITokenTool, contextTool tools.IContextTool) authorizeApi.IAuthorizeService {
	return &AuthorizeService{
		repo:        repo,
		JWTKey:      jwtKey,
		tokenTool:   tokenTool,
		contextTool: contextTool,
	}
}

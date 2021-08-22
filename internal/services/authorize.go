package services

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	authorizeApi "github.com/manabie-com/togo/internal/dto"
	"github.com/manabie-com/togo/internal/storages/repos"
	"github.com/manabie-com/togo/internal/tools"
)

type AuthorizeService struct {
	repo   repos.IAuthorizeRepo
	JWTKey string
}

func (as *AuthorizeService) Login(ctx context.Context, req authorizeApi.LoginRequest) (*authorizeApi.LoginResponse, *tools.TodoError) {
	if !as.repo.ValidateUser(ctx,
		sql.NullString{String: req.UserId, Valid: true},
		sql.NullString{String: req.Password, Valid: true}) {
		return nil, tools.NewTodoError(http.StatusUnauthorized, "incorrect user_id/pwd")
	}
	token, err := tools.CreateToken(req.UserId, as.JWTKey)
	if err != nil {
		return nil, err
	}
	return &authorizeApi.LoginResponse{Data: token}, nil
}

func (as *AuthorizeService) Validate(req *http.Request) (context.Context, *tools.TodoError) {
	token := req.Header.Get("Authorization")

	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(as.JWTKey), nil
	})
	if err != nil {
		log.Println(err)
		return nil, tools.NewTodoError(http.StatusInternalServerError, err.Error())
	}

	if !t.Valid {
		return nil, tools.NewTodoError(http.StatusUnauthorized, "Your request is unauthorized")
	}

	id, ok := claims["user_id"].(string)
	if !ok {
		return nil, tools.NewTodoError(http.StatusInternalServerError, "Something went wrong")
	}

	ctx := tools.WriteUserIDToContext(req.Context(), id)
	return ctx, nil
}

func NewAuthorizeService(repo repos.IAuthorizeRepo, jwtKey string) authorizeApi.IAuthorizeService {
	return &AuthorizeService{
		repo:   repo,
		JWTKey: jwtKey,
	}
}

package auth

import (
	"context"

	"github.com/khangjig/togo/model"
)

type IUseCase interface {
	Login(ctx context.Context, req *LoginRequest) (*ResponseWrapper, error)
}

type ResponseWrapper struct {
	Token string      `json:"token"`
	User  *model.User `json:"user"`
}

package auth

import (
	"context"
	"github.com/HoangVyDuong/togo/internal/usecase/auth"
	"github.com/HoangVyDuong/togo/internal/usecase/user"
	authDTO "github.com/HoangVyDuong/togo/pkg/dtos/auth"
)

type Handler interface {
	Auth(ctx context.Context, request authDTO.AuthUserRequest) (response authDTO.AuthUserResponse, err error)
}

type authHandler struct {
	authService auth.Service
	userService user.Service
}

func NewHander(authService auth.Service, userService user.Service) Handler{
	return &authHandler{authService, userService}
}

func (ah *authHandler) Auth(ctx context.Context, request authDTO.AuthUserRequest) (response authDTO.AuthUserResponse, err error) {
	return authDTO.AuthUserResponse{}, nil
}


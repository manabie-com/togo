package auth

import (
	"context"
	"github.com/HoangVyDuong/togo/internal/usecase/auth"
	authDTO "github.com/HoangVyDuong/togo/pkg/dtos/auth"
)

type Handler interface {
	Auth(ctx context.Context, request authDTO.AuthUserRequest) (response authDTO.AuthUserResponse, err error)
}

type authHandler struct {
	authService auth.Service
}

func NewHander(authService auth.Service) Handler{
	return &authHandler{authService}
}

func (ah *authHandler) Auth(ctx context.Context, request authDTO.AuthUserRequest) (response authDTO.AuthUserResponse, err error) {
	// Auth Username, password
	// Create Token
	// Convert To DTO
	return authDTO.AuthUserResponse{}, nil
}


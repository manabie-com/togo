package auth

import (
	"context"
	authHandler "github.com/HoangVyDuong/togo/internal/handler/auth"
	authDTO "github.com/HoangVyDuong/togo/pkg/dtos/auth"
)

type EndpointSample interface {
	auth(ctx context.Context, request interface{}) (response interface{}, err error)
}

type authEndpoint struct {
	authHandler authHandler.Handler
}

func Endpoint(authHandler authHandler.Handler) EndpointSample{
	return &authEndpoint{authHandler}
}

func (ae *authEndpoint) auth(ctx context.Context, request interface{}) (response interface{}, err error) {
	req := request.(authDTO.AuthUserRequest)
	return ae.authHandler.Auth(ctx, req)
}

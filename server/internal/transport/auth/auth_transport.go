package auth

import (
	"context"
	"encoding/json"
	authHandler "github.com/HoangVyDuong/togo/internal/handler/auth"
	"github.com/HoangVyDuong/togo/internal/kit"
	authDTO "github.com/HoangVyDuong/togo/pkg/dtos/auth"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Endpoint interface {
	auth(ctx context.Context, request interface{}) (response interface{}, err error)
}

type authEndpoint struct {
	authHandler authHandler.Handler
}

func WithEndpoint(authHandler authHandler.Handler) Endpoint{
	return &authEndpoint{authHandler}
}

func (ae *authEndpoint) auth(ctx context.Context, request interface{}) (response interface{}, err error) {
	req := request.(authDTO.AuthUserRequest)
	return ae.authHandler.Auth(ctx, req)
}

func WithHandler(router *httprouter.Router, authHandler authHandler.Handler) {
	router.Handler("POST", "/auth", kit.NewServer(
		WithEndpoint(authHandler).auth,
		decodeAuthRequest,
	))

}

func decodeAuthRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req authDTO.AuthUserRequest
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}


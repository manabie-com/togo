package auth

import (
	"context"
	"encoding/json"
	authHandler "github.com/HoangVyDuong/togo/internal/handler/auth"
	authDTO "github.com/HoangVyDuong/togo/pkg/dtos/auth"
	"github.com/HoangVyDuong/togo/pkg/kit"
	"github.com/julienschmidt/httprouter"
	"net/http"
)


func MakeHandler(router *httprouter.Router, authHandler authHandler.Handler) {
	router.Handler("POST", "/auth", kit.WithCORS(kit.NewServer(
		Endpoint(authHandler).auth,
		decodeAuthRequest,
	)))

}

func decodeAuthRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req authDTO.AuthUserRequest
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}


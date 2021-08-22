package dto

import (
	"context"
	"github.com/manabie-com/togo/internal/tools"
	"net/http"
)

type LoginRequest struct {
	UserId   string `json:"user_id"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Data string `json:"data"`
}

func (lr *LoginResponse) ToRes() interface{} {
	return lr
}

type IAuthorizeApi interface {
	Login(ctx context.Context, req LoginRequest) (*LoginResponse, *tools.TodoError)
	Validate(req *http.Request) (context.Context, *tools.TodoError)
}

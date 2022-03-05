package user

import (
	"context"

	"github.com/khangjig/togo/model"
)

type IUseCase interface {
	GetMe(ctx context.Context) (*ResponseWrapper, error)
}

type ResponseWrapper struct {
	User *model.User `json:"user"`
}

package todo

import (
	"context"

	"github.com/khangjig/togo/model"
)

type IUseCase interface {
	Create(ctx context.Context, req *CreateRequest) (*ResponseWrapper, error)
}

type ResponseWrapper struct {
	Todo *model.Todo `json:"todo"`
}

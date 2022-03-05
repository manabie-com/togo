package todo

import (
	"context"

	"github.com/khangjig/togo/model"
)

type IUseCase interface {
	Create(ctx context.Context, req *CreateRequest) (*ResponseWrapper, error)
	Update(ctx context.Context, req *UpdateRequest) (*ResponseWrapper, error)
	GetByID(ctx context.Context, id int64) (*ResponseWrapper, error)
}

type ResponseWrapper struct {
	Todo *model.Todo `json:"todo"`
}

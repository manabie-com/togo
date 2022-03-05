package todo

import (
	"context"

	"github.com/khangjig/togo/model"
)

type IUseCase interface {
	Create(ctx context.Context, req *CreateRequest) (*ResponseWrapper, error)
	Update(ctx context.Context, req *UpdateRequest) (*ResponseWrapper, error)
	GetByID(ctx context.Context, id int64) (*ResponseWrapper, error)
	GetList(ctx context.Context, req *GetListRequest) (*ResponseListWrapper, error)
}

type ResponseWrapper struct {
	Todo *model.Todo `json:"todo"`
}

type ResponseListWrapper struct {
	Todos []model.Todo `json:"todos"`
	Page  int          `json:"page"`
	Limit int          `json:"limit"`
	Total int64        `json:"total"`
}

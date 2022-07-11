package http

import (
	"github.com/labstack/echo/v4"
)

func NewGetRequest(ctx echo.Context) GetRequest {
	return &getRequest{ctx: ctx}
}

type GetRequest interface {
	Bind() error
	GetID() int
}

type getRequest struct {
	ctx echo.Context
	ID  int `param:"id"`
}

func (p *getRequest) Bind() error {
	return p.ctx.Bind(p)
}

func (p *getRequest) GetID() int {
	return p.ID
}

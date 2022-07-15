package http

import (
	"github.com/datshiro/togo-manabie/internal/infras/errors"
	"github.com/labstack/echo/v4"
)

func NewPostRequest(ctx echo.Context) PostRequest {
	return &postRequest{ctx: ctx}
}

type PostRequest interface {
	Bind() error
	Validate() error
	GetName() string
	GetEmail() string
	GetQuota() int
}

type postRequest struct {
	ctx   echo.Context
	Email string `json:"email"`
	Name  string `json:"name"`
	Quota int    `json:"quota"`
}

func (p *postRequest) GetName() string {
	return p.Name
}

func (p *postRequest) GetEmail() string {
	return p.Email
}

func (p *postRequest) GetQuota() int {
	return p.Quota
}

func (p *postRequest) Bind() error {
	return p.ctx.Bind(p)
}

func (p *postRequest) Validate() error {
	if p.Name == "" {
		return errors.NewParamErr("Name cannot be empty ")
	}
	if p.Quota == 0 {
		return errors.NewParamErr("Quota cannot be 0 ")
	}
	return nil
}

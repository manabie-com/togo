package http

import (
	"github.com/datshiro/togo-manabie/internal/infras/errors"
	"github.com/datshiro/togo-manabie/internal/interfaces/domain"
	"github.com/labstack/echo/v4"
)

func NewPostRequest(ctx echo.Context) PostRequest {
	return &postRequest{ctx: ctx}
}

type PostRequest interface {
	Bind() error
	Validate() error
	GetTitle() string
	GetDescription() string
	GetPriority() int
	GetUserID() int
}

type postRequest struct {
	ctx         echo.Context
	UserID      int    `json:"user_id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
}

func (p *postRequest) Bind() error {
	return p.ctx.Bind(p)
}

func (p *postRequest) Validate() error {
	if p.UserID == 0 {
		return errors.NewParamErr("UserID cannot be 0")
	}
	if p.Title == "" {
		return errors.NewParamErr("Title cannot be empty ")
	}
	return nil
}

func (p *postRequest) GetTitle() string {
	return p.Title
}

func (p *postRequest) GetDescription() string {
	return p.Description
}

func (p *postRequest) GetUserID() int {
	return p.UserID
}

func (p *postRequest) GetPriority() int {
	if p.Priority == 0 {
		p.Priority = domain.TaskPriorityLow
	}
	return p.Priority
}

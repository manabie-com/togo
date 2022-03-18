package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/manabie-com/togo/pkg/errorx"
)

type CreateUserRequest struct {
	Username  string `json:"username"  validate:"required"`
	Password  string `json:"password" validate:"required"`
	TaskLimit int    `json:"task_limit" validate:"required"`
}

func (p *CreateUserRequest) Validate() error {
	if err := validator.New().Struct(p); err != nil {
		return errorx.ErrInvalidParameter(err)
	}

	return nil
}

type UpdateUserRequest struct {
	ID        int
	Password  string `json:"password"`
	TaskLimit int    `json:"task_limit"`
}

type LoginRequest struct {
	ID       int    `json:"id" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginUserResponse struct {
	AccessToken string `json:"access_token"`
	AtExpires   int64  `json:"at_expires"`
}

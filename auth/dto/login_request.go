package dto

import (
	"errors"
	apperror "github.com/manabie-com/togo/shared/app_error"
)

var (
	ErrLoginIdEmpty = apperror.NewCustomError(errors.New("loginId can't empty"),
		"loginId can't empty", "ErrLoginIdEmpty")
	ErrPasswordEmpty = apperror.NewCustomError(errors.New("password can't empty"),
		"password can't empty", "ErrPasswordEmpty")
)

type LoginRequest struct {
	LoginId  string `json:"loginId"`
	Password string `json:"password"`
}

func (u LoginRequest) Validate() error {
	if len(u.LoginId) == 0 {
		return ErrLoginIdEmpty
	}
	if len(u.Password) == 0 {
		return ErrPasswordEmpty
	}
	return nil
}

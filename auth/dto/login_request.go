package dto

import "errors"

var (
	ErrIdEmpty       = errors.New("id can't empty")
	ErrPasswordEmpty = errors.New("password can't empty")
)

type UserInput struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}

func (u UserInput) Validate() error {
	if len(u.Id) == 0 {
		return ErrIdEmpty
	}
	if len(u.Password) == 0 {
		return ErrPasswordEmpty
	}
	return nil
}

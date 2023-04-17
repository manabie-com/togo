package common

import (
	"errors"

	"github.com/phathdt/libs/go-sdk/sdkcm"
)

var (
	ErrEmailOrPasswordInvalid = sdkcm.NewCustomError(
		errors.New("email or password invalid"),
		"email or password invalid",
		"ErrUsernameOrPasswordInvalid",
	)

	ErrEmailExisted = sdkcm.NewCustomError(
		errors.New("email has already existed"),
		"email has already existed",
		"ErrEmailExisted",
	)
)

package common

import (
	"errors"

	"github.com/phathdt/libs/go-sdk/sdkcm"
)

var (
	ErrCannotGetUserLimit = sdkcm.NewCustomError(
		errors.New("cannot get user limit"),
		"cannot get user limit",
		"ErrCannotGetUserLimit",
	)

	ErrLimitTaskToday = sdkcm.NewCustomError(
		errors.New("limit task today"),
		"limit task today",
		"ErrLimitTaskToday",
	)

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

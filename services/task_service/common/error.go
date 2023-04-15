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
)

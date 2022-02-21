package model

import "errors"

var (
	ErrorNotFound   = errors.New("no record found")
	ErrorNotAllowed = errors.New("permission denied, you reach maximum posts a day")
)

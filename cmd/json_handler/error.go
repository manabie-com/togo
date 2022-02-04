package main

import (
	"errors"

	customError "github.com/kozloz/togo/internal/errors"
)

// Error represents our Error in JSON
type Error struct {
	ErrorCode int    `json:"err_code"`
	ErrorDesc string `json:"err_desc"`
}

// Converts the error to our JSON Error type
func CustomErrorToJSON(err error) Error {
	var cerr customError.Error
	if errors.As(err, &cerr) {
		return Error{
			ErrorCode: cerr.ErrorCode,
			ErrorDesc: cerr.ErrorDesc,
		}
	}
	return Error{
		ErrorCode: customError.InternalError.ErrorCode,
		ErrorDesc: customError.InternalError.ErrorDesc,
	}
}

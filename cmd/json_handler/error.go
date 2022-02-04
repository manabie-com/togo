package main

import (
	"errors"
	"log"

	customError "github.com/kozloz/togo/internal/errors"
)

// Error represents our Error in JSON
type Error struct {
	ErrorCode int    `json:"err_code"`
	ErrorDesc string `json:"err_desc"`
}

// Converts the error to our JSON Error type
func CustomErrorToJSON(err error) Error {
	var cerr *customError.Error
	if errors.As(err, &cerr) {
		log.Println(cerr.ErrorCode)
	}
	return Error{}
}

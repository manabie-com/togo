package main

import (
	"testing"

	"github.com/kozloz/togo/internal/errors"
)

func TestCustomErrorToJSON(t *testing.T) {
	errResult := CustomErrorToJSON(errors.MaxLimit)
	if errResult.ErrorCode != errors.MaxLimit.ErrorCode {
		t.Errorf("Expected '%d', got '%d'", errors.MaxLimit.ErrorCode, errResult.ErrorCode)
	}
}

package middleware_test

import (
	"github.com/japananh/togo/middleware"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMiddlewareAuthorize_ExtractTokenFromHeaderString(t *testing.T) {
	var tsc = []struct {
		arg    string
		result string
		err    error
	}{
		{"Bear", "", middleware.ErrWrongAuthHeader(nil)},
		{"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjd9LCJleHAiOjE2NTI2MDYzMjEsImlhdCI6MTY1MjUxOTkyMX0.SITq8_S_Nk5eFBjnxiPQDpjVWAo5ya2GIb7cNy6Ieyw", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjd9LCJleHAiOjE2NTI2MDYzMjEsImlhdCI6MTY1MjUxOTkyMX0.SITq8_S_Nk5eFBjnxiPQDpjVWAo5ya2GIb7cNy6Ieyw", nil},
	}

	for _, tc := range tsc {
		output, err := middleware.ExtractTokenFromHeaderString(tc.arg)
		assert.Equal(t, output, tc.result, "token should be equal")
		assert.Equal(t, err, tc.err, "error should be equal")
	}
}

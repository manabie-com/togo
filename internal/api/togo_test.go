package api

import (
	"context"
	"github.com/jmsemira/togo/internal/auth"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type TogoTestSuite struct {
	suite.Suite
}

func (suite *TogoTestSuite) TestCreateTodoHandler() {
	var tests = []struct {
		mockRequest *http.Request
		expected    string
	}{
		{
			mockRequest: func() *http.Request {
				req, _ := http.NewRequest(http.MethodGet, "/create_todo?name=sample1", nil)

				// Set request context with claims
				ctx := context.WithValue(req.Context(), "user", &auth.Claims{
					ID:        1,
					Username:  "user1",
					RateLimit: 1,
				})
				return req.WithContext(ctx)
			}(),
			expected: `
			{
				"status":"OK"
			}
			`,
		},
		{
			mockRequest: func() *http.Request {
				req, _ := http.NewRequest(http.MethodGet, "/create_todo?name=sample1", nil)
				// Set request context with claims
				ctx := context.WithValue(req.Context(), "user", &auth.Claims{
					ID:        1,
					Username:  "user1",
					RateLimit: 1,
				})
				return req.WithContext(ctx)
			}(),
			expected: `
			{
				"status":"error",
				"err_msg": "Rate Limit for today was reached"
			}
			`,
		},
		{
			mockRequest: func() *http.Request {
				req, _ := http.NewRequest(http.MethodGet, "/create_todo", nil)
				// Set request context with claims
				ctx := context.WithValue(req.Context(), "user", &auth.Claims{
					ID:        1,
					Username:  "user1",
					RateLimit: 1,
				})
				return req.WithContext(ctx)
			}(),
			expected: `
			{
				"status":"error",
				"err_msg": "Name is required!"
			}
			`,
		},
	}

	for _, test := range tests {
		recorder := httptest.NewRecorder()
		CreateTodoHandler(recorder, test.mockRequest)

		// then
		suite.JSONEq(test.expected, recorder.Body.String())
	}
}

func TestTogoSuite(t *testing.T) {
	suite.Run(t, new(TogoTestSuite))
	os.Remove("test.db")
}

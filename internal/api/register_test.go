package api

import (
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type RegisterTestSuite struct {
	suite.Suite
}

func (suite *RegisterTestSuite) TestRegisterHandler() {
	var tests = []struct {
		mockRequest *http.Request
		expected    string
	}{
		{
			mockRequest: func() *http.Request {
				req, _ := http.NewRequest(http.MethodGet, "/register?username=user4&password=user4", nil)
				return req
			}(),
			expected: `
			{
				"status":"ok"
			}
			`,
		},
		{
			mockRequest: func() *http.Request {
				req, _ := http.NewRequest(http.MethodGet, "/register?username=user4&password=user4", nil)
				return req
			}(),
			expected: `
			{
				"status":"error",
				"err_msg": "Username Exist!"
			}
			`,
		},
	}

	for _, test := range tests {
		recorder := httptest.NewRecorder()
		RegisterHandler(recorder, test.mockRequest)

		// then
		suite.JSONEq(test.expected, recorder.Body.String())
	}
}

func TestRegisterSuite(t *testing.T) {
	suite.Run(t, new(RegisterTestSuite))
}

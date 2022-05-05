package middleware

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/suite"
	"net/http/httptest"
	"testing"
)

type MiddlewareTestSuite struct {
	suite.Suite
}

func (suite *MiddlewareTestSuite) TestExtractInvalidBearerToken() {
	// given
	expectedToken := "123456"
	request := httptest.NewRequest("", "http://test.com", nil)
	request.Header.Add("authorization", fmt.Sprintf("Bad %s", expectedToken))

	// when
	token, err := extractBearerToken(request)

	// then
	suite.Nil(token)
	suite.Equal(err, errors.New("Invalid Authorization header"))
}

func TestMiddlewareSuite(t *testing.T) {
	suite.Run(t, new(MiddlewareTestSuite))
}

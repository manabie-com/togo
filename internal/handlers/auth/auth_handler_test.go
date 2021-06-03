package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/manabie-com/togo/internal/consts"
	userService "github.com/manabie-com/togo/internal/services/users"
	"github.com/stretchr/testify/assert"
)

type AuthHandlerLoginTestCase struct {
	mockRequest      ILoginRequest
	mockUserService  userService.IUserService
	expectedRespCode int
	expectedErr      error
}

func AuthHandlerLoginTestCases(t *testing.T) map[string]AuthHandlerLoginTestCase {
	return map[string]AuthHandlerLoginTestCase{
		"Should return nil error on correct": {
			mockRequest: &MockLoginRequest{
				ShouldBindIsCalled:        true,
				ShouldValidateIsCalled:    true,
				ShouldGetUserIDIsCalled:   true,
				ShouldGetPasswordIsCalled: true,
				Testing:                   t,
			},
			mockUserService: userService.MockUserService{
				ShouldValidateIsCalled: true,
				Testing:                t,
			},
			expectedRespCode: http.StatusOK,
			expectedErr:      nil,
		},
		"Should return error on validate fail": {
			mockRequest: &MockLoginRequest{
				ValidateResponse:          consts.ErrInvalidRequest,
				ShouldBindIsCalled:        true,
				ShouldValidateIsCalled:    true,
				ShouldGetUserIDIsCalled:   false,
				ShouldGetPasswordIsCalled: false,
				Testing:                   t,
			},
			mockUserService: userService.MockUserService{
				ShouldValidateIsCalled: true,
				Testing:                t,
			},
			expectedRespCode: 0,
			expectedErr:      consts.ErrInvalidRequest,
		},
	}
}

func TestAuthHandlerLogin(t *testing.T) {
	t.Parallel()
	cases := AuthHandlerLoginTestCases(t)
	for caseName, tCase := range cases {
		t.Run(caseName, func(t *testing.T) {
			authHandler := AuthHandler{
				NewRequest:  func() ILoginRequest { return tCase.mockRequest },
				UserService: tCase.mockUserService,
			}
			response := &httptest.ResponseRecorder{}
			request := &http.Request{}
			got := authHandler.Login(response, request)
			assert.Equal(t, tCase.expectedErr, got)
			assert.Equal(t, tCase.expectedRespCode, response.Code)
		})
	}
}

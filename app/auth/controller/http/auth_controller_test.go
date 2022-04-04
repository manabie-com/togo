package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ansidev/togo/domain/user"
	"github.com/ansidev/togo/errs"
	"github.com/ansidev/togo/test"
	"github.com/ansidev/togo/wire"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestAuthController(t *testing.T) {
	suite.Run(t, new(AuthControllerTestSuite))
}

type AuthControllerTestSuite struct {
	test.ControllerTestSuite
	tokenTTL time.Duration
}

func (s *AuthControllerTestSuite) SetupSuite() {
	s.ControllerTestSuite.SetupSuite()
	s.tokenTTL = 24 * time.Hour

	authService := wire.InitAuthService(s.GormDb, s.Rdb, 1*time.Minute)

	NewAuthController(s.Router, authService)
}

func (s *AuthControllerTestSuite) TestLogin_SubmitEmptyBodyRequest_ShouldReturn400() {
	w := httptest.NewRecorder()
	body := `{}`
	req, _ := http.NewRequest("POST", "/auth/v1/login", bytes.NewBufferString(body))
	s.Router.ServeHTTP(w, req)

	resp := w.Body.Bytes()

	require.NotNil(s.T(), resp)
	require.Equal(s.T(), http.StatusBadRequest, w.Code)

	var actualBody errs.ErrorResponse
	err := json.Unmarshal(resp, &actualBody)

	require.NoError(s.T(), err)
	require.Equal(s.T(), errs.ErrCodeValidationError, actualBody.Code)
	require.Equal(s.T(), errs.ErrValidationError, actualBody.Message)
	require.Equal(s.T(), 2, len(actualBody.Errors))
	require.Equal(s.T(), "username", actualBody.Errors[0].Field)
	require.Equal(s.T(), "Field 'username' is required", actualBody.Errors[0].Message)
	require.Equal(s.T(), "password", actualBody.Errors[1].Field)
	require.Equal(s.T(), "Field 'password' is required", actualBody.Errors[1].Message)
}

func (s *AuthControllerTestSuite) TestLogin_SubmitOnlyUsername_ShouldReturn400() {
	w := httptest.NewRecorder()
	body := `{"username": "test_user"}`
	req, _ := http.NewRequest("POST", "/auth/v1/login", bytes.NewBufferString(body))
	s.Router.ServeHTTP(w, req)

	resp := w.Body.Bytes()

	require.NotNil(s.T(), resp)
	require.Equal(s.T(), http.StatusBadRequest, w.Code)

	var actualBody errs.ErrorResponse
	err := json.Unmarshal(resp, &actualBody)

	require.NoError(s.T(), err)
	require.Equal(s.T(), errs.ErrCodeValidationError, actualBody.Code)
	require.Equal(s.T(), errs.ErrValidationError, actualBody.Message)
	require.Equal(s.T(), 1, len(actualBody.Errors))
	require.Equal(s.T(), "password", actualBody.Errors[0].Field)
	require.Equal(s.T(), "Field 'password' is required", actualBody.Errors[0].Message)
}

func (s *AuthControllerTestSuite) TestLogin_SubmitOnlyPassword_ShouldReturn400() {
	w := httptest.NewRecorder()
	body := `{"password": "test_password"}`
	req, _ := http.NewRequest("POST", "/auth/v1/login", bytes.NewBufferString(body))
	s.Router.ServeHTTP(w, req)

	resp := w.Body.Bytes()

	require.NotNil(s.T(), resp)
	require.Equal(s.T(), http.StatusBadRequest, w.Code)

	var actualBody errs.ErrorResponse
	err := json.Unmarshal(resp, &actualBody)

	require.NoError(s.T(), err)
	require.Equal(s.T(), errs.ErrCodeValidationError, actualBody.Code)
	require.Equal(s.T(), errs.ErrValidationError, actualBody.Message)
	require.Equal(s.T(), 1, len(actualBody.Errors))
	require.Equal(s.T(), "username", actualBody.Errors[0].Field)
	require.Equal(s.T(), "Field 'username' is required", actualBody.Errors[0].Message)
}

func (s *AuthControllerTestSuite) TestLogin_SubmitUnregisteredUsername_ShouldReturn401() {
	w := httptest.NewRecorder()
	body := `{"username": "test_user", "password": "test_password"}`
	req, _ := http.NewRequest("POST", "/auth/v1/login", bytes.NewBufferString(body))
	s.Router.ServeHTTP(w, req)

	resp := w.Body.Bytes()

	require.NotNil(s.T(), resp)
	require.Equal(s.T(), http.StatusUnauthorized, w.Code)

	var actualBody errs.ErrorResponse
	err := json.Unmarshal(resp, &actualBody)

	require.NoError(s.T(), err)
	require.NoError(s.T(), err)
	require.Equal(s.T(), errs.ErrCodeUnauthorized, actualBody.Code)
	require.Equal(s.T(), "Unauthorized", actualBody.Message)
	require.True(s.T(), len(actualBody.Error) > 0)
}

func (s *AuthControllerTestSuite) TestLogin_SubmitRegisteredUsername_ShouldReturn200() {
	username := "test_user"
	password := "test_password"
	s.initUserData(username, password)

	w := httptest.NewRecorder()
	body := fmt.Sprintf(`{"username": "%s", "password": "%s"}`, username, password)
	req, _ := http.NewRequest("POST", "/auth/v1/login", bytes.NewBufferString(body))
	s.Router.ServeHTTP(w, req)

	resp := w.Body.Bytes()

	require.NotNil(s.T(), resp)
	require.Equal(s.T(), http.StatusOK, w.Code)

	type tokenResponse struct {
		Token string `json:"token"`
	}

	var actualBody tokenResponse
	err := json.Unmarshal(resp, &actualBody)

	require.NoError(s.T(), err)
	require.NoError(s.T(), err)
	require.True(s.T(), len(actualBody.Token) > 0)
}

func (s *AuthControllerTestSuite) initUserData(username string, password string) {
	hashedPassword, err1 := bcrypt.GenerateFromPassword([]byte(password), 12)
	require.NoError(s.T(), err1)

	_, err2 := s.SqlDb.Exec(`INSERT INTO "user" (id, username, password, max_daily_task, created_at) VALUES ($1, $2, $3, $4, $5)`,
		1,
		username,
		hashedPassword,
		5,
		"2022-02-22 01:23:45")

	require.NoError(s.T(), err2)

	var userModel user.User

	result := s.GormDb.First(&userModel, 1)
	require.NoError(s.T(), result.Error)
}

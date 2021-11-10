package http_test

import (
	"github.com/labstack/echo/v4"
	"github.com/manabie-com/togo/internal/domain"
	auth_http "github.com/manabie-com/togo/internal/transport/http"
	"github.com/manabie-com/togo/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	sampleAccessToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjgxN2MxNWMtNDFkOC0xMWVjLTlkY2QtZTQ1NGU4MjUyZjUyIiwidXNlcm5hbWUiOiJmaXJzdFVzZXIiLCJpc3N1ZWRBdCI6IjIwMjEtMTEtMTBUMTA6NDY6NTMuNWQ4MjI1MDI5KzA3OjAwIiwiZXhwaXJlZEF0IjoiMjAyMS0xMS0xMFQxMTowMTo1My41ODIyNTAzMjcrMDc6MDAifQ.sOt7ty930uML4PRv6BQ2C6wMpRk7SA-sHSGWz1yl2kM"
)

func TestLoginSuccess(t *testing.T) {
	mockAuthUseCase := new(mocks.AuthUseCase)
	username := "firstUser"
	password := "example"
	mockAuthUseCase.On("SignIn", mock.Anything, username, password).
		Return(sampleAccessToken, nil)
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/login?username="+username+"&password="+password, nil)
	rec := httptest.NewRecorder()
	auth_http.NewAuthHandler(e, mockAuthUseCase)
	e.ServeHTTP(rec, req)
	require.Equal(t, rec.Code, http.StatusOK)
}

func TestLoginMissingUsername(t *testing.T) {
	mockAuthUseCase := new(mocks.AuthUseCase)
	username := "firstUser"
	password := "example"
	mockAuthUseCase.On("SignIn", mock.Anything, username, password).
		Return(sampleAccessToken, nil)
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/login?password="+password, nil)
	rec := httptest.NewRecorder()
	auth_http.NewAuthHandler(e, mockAuthUseCase)
	e.ServeHTTP(rec, req)
	require.Equal(t, rec.Code, http.StatusBadRequest)
}

func TestLoginMissingPassword(t *testing.T) {
	mockAuthUseCase := new(mocks.AuthUseCase)
	username := "firstUser"
	password := "example"
	mockAuthUseCase.On("SignIn", mock.Anything, username, password).
		Return(sampleAccessToken, nil)
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/login?username="+username, nil)
	rec := httptest.NewRecorder()
	auth_http.NewAuthHandler(e, mockAuthUseCase)
	e.ServeHTTP(rec, req)
	require.Equal(t, rec.Code, http.StatusBadRequest)
}

func TestLoginUseCaseReturnUnAuthorized(t *testing.T) {
	mockAuthUseCase := new(mocks.AuthUseCase)
	username := "firstUser"
	password := "example"
	mockAuthUseCase.On("SignIn", mock.Anything, username, password).
		Return("", domain.WrongPassword)
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/login?username="+username+"&password="+password, nil)
	rec := httptest.NewRecorder()
	auth_http.NewAuthHandler(e, mockAuthUseCase)
	e.ServeHTTP(rec, req)
	require.Equal(t, rec.Code, http.StatusUnauthorized)
}

func TestLoginUseCaseReturnUserNotFound(t *testing.T) {
	mockAuthUseCase := new(mocks.AuthUseCase)
	username := "firstUser"
	password := "example"
	mockAuthUseCase.On("SignIn", mock.Anything, username, password).
		Return("", domain.UserNotFound)
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/login?username="+username+"&password="+password, nil)
	rec := httptest.NewRecorder()
	auth_http.NewAuthHandler(e, mockAuthUseCase)
	e.ServeHTTP(rec, req)
	require.Equal(t, rec.Code, http.StatusNotFound)
}

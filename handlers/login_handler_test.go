package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/manabie-com/togo/domains"
	"github.com/manabie-com/togo/usecases"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginHandler_Success(t *testing.T) {
	req := createLoginRequest("user", "pass")
	expectedToken := "AAAAA"
	expectedUserId := int64(9999)

	mockRepo := new(usecases.DBMock)
	mockRepo.On("VerifyUser", req.Context(), &domains.LoginRequest{Username: "user", Password: "pass"}).
		Return(&domains.User{Id: expectedUserId}, nil)

	mockAuth := new(usecases.AuthMock)
	mockAuth.On("CreateToken", expectedUserId).
		Return(expectedToken, nil)

	uc := usecases.NewLoginUseCase(mockRepo, mockAuth)
	recorder := httptest.NewRecorder()

	handler := &LoginHandler{Uc: uc}
	handler.ServeHTTP(recorder, req)

	mockAuth.AssertExpectations(t)
	mockRepo.AssertExpectations(t)

	resp := recorder.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	data := &struct {
		Data struct{
			Token string `json:"token"`
		} `json:"data"`
	}{}
	err := json.NewDecoder(resp.Body).Decode(data)
	assert.Nil(t, err)
	assert.Equal(t, expectedToken, data.Data.Token)
}

func TestLoginHandler_BadRequestWithEmptyUsername(t *testing.T) {
	req := createLoginRequest("", "AAAA")
	expectedStatus := "error"

	mockRepo := new(usecases.DBMock)
	mockAuth := new(usecases.AuthMock)
	uc := usecases.NewLoginUseCase(mockRepo, mockAuth)
	recorder := httptest.NewRecorder()

	handler := &LoginHandler{Uc: uc}
	handler.ServeHTTP(recorder, req)

	resp := recorder.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	data := &struct {
		Status string `json:"status"`
	}{}
	err := json.NewDecoder(resp.Body).Decode(data)
	assert.Nil(t, err)
	assert.Equal(t, expectedStatus, data.Status)
}

func TestLoginHandler_BadRequestWithWrongUsernameAndPassword(t *testing.T) {
	req := createLoginRequest("user", "pass")
	expectedStatus := "error"

	mockRepo := new(usecases.DBMock)
	mockRepo.On("VerifyUser", req.Context(), &domains.LoginRequest{Username: "user", Password: "pass"}).
		Return(nil, domains.ErrorNotFound)
	mockAuth := new(usecases.AuthMock)

	uc := usecases.NewLoginUseCase(mockRepo, mockAuth)
	recorder := httptest.NewRecorder()

	handler := &LoginHandler{Uc: uc}
	handler.ServeHTTP(recorder, req)

	mockAuth.AssertExpectations(t)
	mockRepo.AssertExpectations(t)

	resp := recorder.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	data := &struct {
		Status string `json:"status"`
	}{}
	err := json.NewDecoder(resp.Body).Decode(data)
	assert.Nil(t, err)
	assert.Equal(t, expectedStatus, data.Status)
}

func createLoginRequest(username, password string) *http.Request {
	body, _ := json.Marshal(&usecases.LoginInput{Username: username, Password: password})
	req := httptest.NewRequest("POST", "localhost:5050/login", bytes.NewBuffer(body))
	return req
}

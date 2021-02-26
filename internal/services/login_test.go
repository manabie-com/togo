package services

import (
	"bytes"
	"encoding/json"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/storages/postgres"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	testJWTKey = "1234567890abcdxyz"
	testUser   = &loginParams{Username: "userid", Password: "password"}
)

func newLoginRequest(username, password string) *http.Request {
	body, _ := json.Marshal(&loginParams{Username: username, Password: password})
	req := httptest.NewRequest("POST", "localhost:5050/login", bytes.NewBuffer(body))
	return req
}

func mockCreateToken(t *testing.T, user *loginParams, err error) *http.Response {
	req := newLoginRequest(user.Username, user.Password)
	db := new(postgres.DatabaseMock)
	db.On("ValidateUser", req.Context(), user.Username, user.Password).Return(&storages.PgUser{}, err)

	s := NewToDoService(testJWTKey, ":6000", db)

	recorder := httptest.NewRecorder()
	s.createTokenHandler(recorder, req)
	db.AssertExpectations(t)

	return recorder.Result()
}

func TestLoginCorrectUsernamePassword(t *testing.T) {
	resp := mockCreateToken(t, testUser, nil)
	defer func() {
		_ = resp.Close
	}()

	requireTest := require.New(t)

	requireTest.Equal(http.StatusOK, resp.StatusCode)

	data := &struct {
		Data string `json:"data"`
	}{}
	err := json.NewDecoder(resp.Body).Decode(data)
	requireTest.NoError(err)
	requireTest.NotEmpty(data.Data)
}

func TestLoginWrongUsernamePassword(t *testing.T) {
	resp := mockCreateToken(t, testUser, postgres.ErrIncorrectUsernameOrPassword)
	defer func() {
		_ = resp.Close
	}()

	requireTest := require.New(t)

	requireTest.Equal(http.StatusBadRequest, resp.StatusCode)

	data := &struct {
		Err string `json:"error"`
	}{}
	err := json.NewDecoder(resp.Body).Decode(data)
	requireTest.NoError(err)
	requireTest.NotEmpty(data.Err)
}

func TestAuth(t *testing.T) {

}
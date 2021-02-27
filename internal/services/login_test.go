package services

import (
	"bytes"
	"encoding/json"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/storages/postgres"
	"github.com/stretchr/testify/require"
	"io/ioutil"
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
	db.On("ValidateUser", req.Context(), user.Username, user.Password).Return(&storages.User{}, err)

	s := NewToDoService(testJWTKey, ":6000", db)

	recorder := httptest.NewRecorder()
	s.createTokenHandler(recorder, req)
	db.AssertExpectations(t)

	return recorder.Result()
}

func TestLoginCorrectUsernamePassword(t *testing.T) {
	resp := mockCreateToken(t, testUser, nil)
	defer resp.Body.Close()

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
	defer resp.Body.Close()

	requireTest := require.New(t)
	requireTest.Equal(http.StatusBadRequest, resp.StatusCode)

	expectedErrResp := &ApiErrResp{Error: postgres.ErrIncorrectUsernameOrPassword.Error()}
	assertErrResp(t, expectedErrResp, resp)
}

func TestAuthSuccess(t *testing.T) {
	resp := mockGetAuthToken(t, storages.User{Id: 1, Username: "abc"}, true)
	defer resp.Body.Close()

	requireTest := require.New(t)
	requireTest.Equal(http.StatusOK, resp.StatusCode)
}

func TestAuthFailure(t *testing.T) {
	resp := mockGetAuthToken(t, storages.User{Id: 1, Username: "abc"}, false)
	defer resp.Body.Close()

	requireTest := require.New(t)
	requireTest.Equal(http.StatusUnauthorized, resp.StatusCode)
}

func mockGetAuthToken(t *testing.T, user storages.User, expectedValidToken bool) *http.Response {
	requireTest := require.New(t)
	req := httptest.NewRequest("GET", "localhost:5050/login", nil)

	db := new(postgres.DatabaseMock)
	s := NewToDoService(testJWTKey, ":6000", db)

	if expectedValidToken {
		token, err := s.createToken(user.Id)
		requireTest.NoError(err)
		req.Header.Set("Authorization", token)
	} else {
		req.Header.Set("Authorization", "invalid token")
	}

	recorder := httptest.NewRecorder()
	s.authHandler(func(writer http.ResponseWriter, request *http.Request) {})(recorder, req)
	return recorder.Result()
}

func assertDataResp(t *testing.T, expected *ApiDataResp, actualResp *http.Response) {
	defer actualResp.Body.Close()
	requireTest := require.New(t)

	actualRespBytes, err := ioutil.ReadAll(actualResp.Body)
	requireTest.NoError(err)

	expectedBytes, err := json.Marshal(expected)
	requireTest.NoError(err)
	requireTest.Equal(expectedBytes, bytes.TrimSpace(actualRespBytes))
}

func assertErrResp(t *testing.T, expected *ApiErrResp, actualResp *http.Response) {
	defer actualResp.Body.Close()
	requireTest := require.New(t)

	actualRespBytes, err := ioutil.ReadAll(actualResp.Body)
	requireTest.NoError(err)

	expectedBytes, err := json.Marshal(expected)
	requireTest.NoError(err)
	requireTest.Equal(expectedBytes, bytes.TrimSpace(actualRespBytes))
}
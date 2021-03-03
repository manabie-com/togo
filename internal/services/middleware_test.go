package services

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_authenticateMiddleware (t *testing.T){
	Test_loginHandler_Happy(t)

	reqURL := "/mock"
	reqToken := Token

	req := httptest.NewRequest("", reqURL, nil)
	req.Header.Add("Authorization", reqToken)
	respW := httptest.NewRecorder()

	var nextHndl http.HandlerFunc
	nextHndl = func(writer http.ResponseWriter, request *http.Request) {}
	s.authenticateMiddleware(nextHndl).ServeHTTP(respW, req)
	resp := respW.Result()

	require.NotEqual(t, http.StatusUnauthorized, resp.StatusCode)
}

func Test_authenticateMiddleware_Fail (t *testing.T){
	reqURL := "/mock"
	reqToken := "qweeasedada"

	req := httptest.NewRequest("", reqURL, nil)
	req.Header.Add("Authorization", reqToken)
	respW := httptest.NewRecorder()

	var nextHndl http.HandlerFunc
	nextHndl = func(writer http.ResponseWriter, request *http.Request) {}
	s.authenticateMiddleware(nextHndl).ServeHTTP(respW, req)
	resp := respW.Result()

	require.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}
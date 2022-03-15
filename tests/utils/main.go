package test_utils

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

type LoginRes struct {
	Token   string `json:"access_token"`
	Message string `json:"message"`
	Error   bool   `json:"error"`
}

func ApiRoute(method string, uri string, body io.Reader, token string) *http.Request {
	r1 := httptest.NewRequest(method, uri, body)
	r1.Header.Add("Content-Type", "application/json")
	r1.Header.Set("Accept", "application/json")
	if len(token) > 0 {
		var bearer = "Bearer " + token
		r1.Header.Set("Authorization", bearer)
	}
	return r1
}

func GetBody(resp *http.Response) (string, error) {
	defer resp.Body.Close()
	bodyR, err := ioutil.ReadAll(resp.Body)
	return string(bodyR), err
}

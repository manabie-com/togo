package task

import (
	"bytes"
	"encoding/json"
	"github.com/HoangVyDuong/togo/pkg/dtos/auth"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

var TaskURL *url.URL
var httpClient *http.Client

func init() {
	TaskURL, _ = url.Parse("http://localhost:8080/api/tasks")
	httpClient = &http.Client{
		Timeout: 10 * time.Second,
	}
}

func GetTrueToken() string{
	AuthURL, _ := url.Parse("http://localhost:8080/api/auth")

	jsonReq, _ := json.Marshal(auth.AuthUserRequest{
		Username: "firstUser",
		Password: "example",
	})

	resp, _ := httpClient.Do(&http.Request{
		Method: "POST",
		URL:    AuthURL,
		Header: map[string][]string{
			"Content-Type": {"application/json; charset=utf-8"},
		},
		Body: ioutil.NopCloser(bytes.NewBuffer(jsonReq)),
	})

	var authResp auth.AuthUserResponse
	_ = json.NewDecoder(resp.Body).Decode(&authResp)
	return authResp.Token
}

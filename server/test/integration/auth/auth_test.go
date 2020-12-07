//+build !unit

package auth

import (
	"bytes"
	"encoding/json"
	"github.com/HoangVyDuong/togo/pkg/define"
	"github.com/HoangVyDuong/togo/pkg/dtos"
	"github.com/HoangVyDuong/togo/pkg/dtos/auth"
	"github.com/HoangVyDuong/togo/pkg/logger"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

func Test_Auth(t *testing.T) {
	tests := []struct {
		name string
		request auth.AuthUserRequest
		statusCode int
		errResponse dtos.ErrorResponse
	}{
		{
			name: "Auth User Success",
			request: auth.AuthUserRequest{
				Username: "firstUser",
				Password: "example",
			},
			statusCode: 200,
		},
		{
			name: "Auth User Fail By Wrong Username",
			request: auth.AuthUserRequest{
				Username: "firstUserFalse",
				Password: "example",
			},
			statusCode: 404,
			errResponse: dtos.ErrorResponse{
				Message: define.AccountNotExist.Error(),
			},
		},
		{
			name: "Auth User Fail By Wrong Password",
			request: auth.AuthUserRequest{
				Username: "firstUser",
				Password: "exampleWrong",
			},
			statusCode: 401,
			errResponse: dtos.ErrorResponse{
				Message: define.AccountNotAuthorized.Error(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonReq, err := json.Marshal(tt.request)
			if err != nil {
				t.Errorf("Cannot marshal request to json: %v", err)
			}

			resp, err := httpClient.Do(&http.Request{
				Method: "POST",
				URL:    AuthURL,
				Header: map[string][]string{
					"Content-Type": {"application/json; charset=utf-8"},
				},
				Body: ioutil.NopCloser(bytes.NewBuffer(jsonReq)),
			})
			if err != nil {
				t.Errorf("this is the error: %v\n", err)
				return
			}
			defer resp.Body.Close()

			if tt.statusCode != resp.StatusCode {
				t.Errorf("Actual Code = %v, Expected Code %v", resp.StatusCode, tt.statusCode)
				return
			}
			if tt.statusCode != 200 {
				var errResponse dtos.ErrorResponse
				err = json.NewDecoder(resp.Body).Decode(&errResponse)
				if err != nil {
					t.Errorf("Cannot convert to json: %v", err)
					return
				}
				if !reflect.DeepEqual(errResponse, tt.errResponse) {
					t.Errorf("Expected: %v, Actual: %v", tt.errResponse, errResponse)
					return
				}
			} else {
				var authResp auth.AuthUserResponse
				err = json.NewDecoder(resp.Body).Decode(&authResp)
				if err != nil {
					t.Errorf("Cannot convert to json: %v", err)
					return
				}
				logger.Debugf("token response: %s", authResp.Token)
			}
		})
	}
}

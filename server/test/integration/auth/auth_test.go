//+build !unit

package auth

import (
	"encoding/json"
	"github.com/HoangVyDuong/togo/pkg/define"
	"github.com/HoangVyDuong/togo/pkg/dtos"
	"net/http"
	"reflect"
	"testing"
)

func Test_Auth(t *testing.T) {
	tests := []struct {
		name string
		statusCode int
		errResponse dtos.ErrorResponse
	}{
		{
			name: "Auth User Success",
			statusCode: 200,
		},
		{
			name: "Auth User Fail By Wrong Username",
			statusCode: 404,
			errResponse: dtos.ErrorResponse{
				Message: define.AccountNotExist,
			},
		},
		{
			name: "Auth User Fail By Wrong Password",
			statusCode: 401,
			errResponse: dtos.ErrorResponse{
				Message: define.AccountNotAuthorized,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := httpClient.Do(&http.Request{
				Method: "POST",
				URL:    AuthURL,
				Header: map[string][]string{
					"Content-Type": {"application/json; charset=utf-8"},
				},
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
			}
		})
	}
}

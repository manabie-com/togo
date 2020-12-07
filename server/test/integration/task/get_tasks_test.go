//+build !unit

package task

import (
	"encoding/json"
	"github.com/HoangVyDuong/togo/pkg/define"
	"github.com/HoangVyDuong/togo/pkg/dtos"
	"github.com/HoangVyDuong/togo/pkg/dtos/task"
	"github.com/HoangVyDuong/togo/pkg/logger"
	"net/http"
	"reflect"
	"testing"
)

func Test_GetTasks(t *testing.T) {
	TrueToken := GetTrueToken()
	ErrorToken := "error token"
	tests := []struct {
		name string
		authorization string
		statusCode int
		errResponse dtos.ErrorResponse
	}{
		{
			name: "Get Tasks Success",
			authorization: TrueToken,
			statusCode: http.StatusOK,
		},
		{
			name: "Get Task Failed By Invalid Token",
			authorization: ErrorToken,
			statusCode: http.StatusUnauthorized,
			errResponse: dtos.ErrorResponse{
				Message: define.Unauthenticated.Error(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := httpClient.Do(&http.Request{
				Method: "GET",
				URL:    TaskURL,
				Header: map[string][]string{
					"Authorization": {tt.authorization},
					"Content-Type": {"application/json; charset=utf-8"},
				},
			})
			if err != nil {
				t.Errorf("this is the error: %v\n", err)
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
				getTaskResponse := task.Tasks{}
				err = json.NewDecoder(resp.Body).Decode(&getTaskResponse)
				if err != nil {
					t.Errorf("Cannot convert to json: %v", err)
				}
				if len(getTaskResponse.Data) == 0 {
					t.Errorf("TaskID can not be empty")
				}
				logger.Debugf("createTaskResponse: %v", getTaskResponse)
			}

		})
	}

}

//+build !unit

package task

import (
	"bytes"
	"encoding/json"
	"github.com/HoangVyDuong/togo/pkg/define"
	"github.com/HoangVyDuong/togo/pkg/dtos"
	"github.com/HoangVyDuong/togo/pkg/dtos/task"
	"github.com/HoangVyDuong/togo/pkg/logger"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

func Test_CreateTask(t *testing.T) {
	TrueToken := GetTrueToken()
	ErrorToken := "error token"
	tests := []struct {
		name string
		request  task.CreateTaskRequest
		authorization string
		statusCode int
		errResponse dtos.ErrorResponse
	}{
		{
			name: "Create Task Success",
			authorization: TrueToken,
			request:  task.CreateTaskRequest{
				Content: "Third Content",
			},
			statusCode: http.StatusOK,
		},
		{
			name: "Create Task Failed By Empty Content",
			authorization: TrueToken,
			request:  task.CreateTaskRequest{
				Content: "",
			},
			errResponse: dtos.ErrorResponse{
				Message: define.FailedValidation.Error(),
			},
			statusCode: http.StatusBadRequest,
		},
		{
			name: "Create Task Failed By Invalid Token",
			authorization: ErrorToken,
			request:  task.CreateTaskRequest{
				Content: "Second Content",
			},
			errResponse: dtos.ErrorResponse{
				Message: define.Unauthenticated.Error(),
			},
			statusCode: http.StatusUnauthorized,
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
				URL:    TaskURL,
				Header: map[string][]string{
					"Authorization": {tt.authorization},
					"Content-Type": {"application/json; charset=utf-8"},
				},
				Body: ioutil.NopCloser(bytes.NewBuffer(jsonReq)),
			})
			if err != nil {
				t.Errorf("this is the error: %v\n", err)
			}
			defer resp.Body.Close()

			if tt.statusCode != resp.StatusCode {
				t.Errorf("Actual Code = %v, Expected Code %v", resp.StatusCode, tt.statusCode)
				return
			}

			if tt.statusCode != 201 {
				errResponse := dtos.ErrorResponse{}
				err = json.NewDecoder(resp.Body).Decode(&errResponse)
				if err != nil {
					t.Errorf("Cannot convert to json: %v", err)
				}
				if !reflect.DeepEqual(errResponse, tt.errResponse) {
					t.Errorf("Expected: %v, Actual: %v", tt.errResponse, errResponse)
				}
				logger.Debugf("ErrorResponse: %v", errResponse)
			} else {
				createTaskResponse := task.CreateTaskResponse{}
				err = json.NewDecoder(resp.Body).Decode(&createTaskResponse)
				if err != nil {
					t.Errorf("Cannot convert to json: %v", err)
				}
				if createTaskResponse.TaskID == "" {
					t.Errorf("TaskID can not be empty")
				}
				logger.Debugf("createTaskResponse: %v", createTaskResponse)
			}
		})
	}

}

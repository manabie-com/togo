//+build !unit

package task

import (
	"bytes"
	"encoding/json"
	"github.com/manabie-com/togo/pkg/define"
	"github.com/manabie-com/togo/pkg/dtos"
	"github.com/manabie-com/togo/pkg/dtos/task"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

func Test_CreateTask(t *testing.T) {
	tests := []struct {
		name string
		request  task.CreateTaskRequest
		authorization string
		statusCode int
		errResponse dtos.ErrorResponse
	}{
		{
			name: "Get Task Success",
			authorization: TrueToken,
			request:  task.CreateTaskRequest{
				Content: "First Content",
			},
			statusCode: 200,
		},
		{
			name: "Create Task Failed By Invalid Token",
			authorization: TrueToken,
			request:  task.CreateTaskRequest{
				Content: "",
			},
			errResponse: dtos.ErrorResponse{
				Message: define.Unauthenticated,
			},
			statusCode: 401,
		},
		{
			name: "Create Task Failed By Invalid Token",
			authorization: ErrorToken,
			request:  task.CreateTaskRequest{
				Content: "First Content",
			},
			errResponse: dtos.ErrorResponse{
				Message: define.Unauthenticated,
			},
			statusCode: 401,
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
					return
				}
				if !reflect.DeepEqual(errResponse, tt.errResponse) {
					t.Errorf("Expected: %v, Actual: %v", tt.errResponse, errResponse)
				}
			} else {
				createTaskResponse := task.CreateTaskResponse{}
				err = json.NewDecoder(resp.Body).Decode(&createTaskResponse)
				if err != nil {
					t.Errorf("Cannot convert to json: %v", err)
				}
				if createTaskResponse.TaskID == "" {
					t.Errorf("TaskID can not be empty")
				}
			}
		})
	}

}

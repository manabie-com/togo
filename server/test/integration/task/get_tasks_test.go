//+build !unit

package task

import (
	"encoding/json"
	"github.com/manabie-com/togo/pkg/define"
	"github.com/manabie-com/togo/pkg/dtos"
	"net/http"
	"reflect"
	"testing"
)

func Test_GetTasks(t *testing.T) {
	tests := []struct {
		name string
		authorization string
		statusCode int
		errResponse dtos.ErrorResponse
	}{
		{
			name: "Get Tasks Success",
			authorization: TrueToken,
			statusCode: 200,
		},
		{
			name: "Create Task Failed By Invalid UserID (not 0)",
			authorization: ErrorToken,
			statusCode: 401,
			errResponse: dtos.ErrorResponse{
				Message: define.Unauthenticated,
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
			}
		})
	}

}

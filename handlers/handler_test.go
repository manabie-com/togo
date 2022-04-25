package handlers

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestTodoTaskHandler(t *testing.T) {
	handler := New()
	recorder := httptest.NewRecorder()

	route := "/user/{userId}/todo"

	type ReqBody struct {
		Title      string `json:"title"`
		Detail     string `json:"detail"`
		RemindDate string `json:"remind_date"`
	}

	tests := []struct {
		name    string
		payload ReqBody
		method  string
		wantErr bool
		userId  string
	}{
		{
			name: "Success",
			payload: ReqBody{
				Title:      "Sample title",
				Detail:     "Sample detail",
				RemindDate: "2022-04-25",
			},
			wantErr: false,
			userId:  "1",
		},
		{
			name: "Success",
			payload: ReqBody{
				Title:      "Sample title",
				Detail:     "Sample detail",
				RemindDate: "2022-04-25",
			},
			wantErr: false,
			userId:  "1",
		},
		{
			name: "Success",
			payload: ReqBody{
				Title:      "Sample title",
				Detail:     "Sample detail",
				RemindDate: "2022-04-25",
			},
			wantErr: false,
			userId:  "1",
		},
		{
			name: "Title Empty",
			payload: ReqBody{
				Detail:     "Sample detail",
				RemindDate: "2022-04-25",
			},
			wantErr: true,
			userId:  "1",
		},
		{
			name: "Title Detail",
			payload: ReqBody{
				Title:      "Sample title",
				RemindDate: "2022-04-25",
			},
			wantErr: false,
			userId:  "1",
		},
		{
			name: "Remind Date Empty",
			payload: ReqBody{
				Title:  "Sample title",
				Detail: "Sample detail",
			},
			wantErr: false,
			userId:  "1",
		},
		{
			name: "User Id Not Exist",
			payload: ReqBody{
				Title:  "Sample title",
				Detail: "Sample detail",
			},
			wantErr: true,
			userId:  "4",
		},
		{
			name: "User Id Negative",
			payload: ReqBody{
				Title:      "Sample title",
				Detail:     "Sample detail",
				RemindDate: "2022-04-25",
			},
			wantErr: true,
			userId:  "-1",
		},
		{
			name: "Exceed Limit",
			payload: ReqBody{
				Title:      "Sample title",
				Detail:     "Sample detail",
				RemindDate: "2022-04-25",
			},
			wantErr: true,
			userId:  "1",
		},
	}

	for _, tt := range tests {
		var body bytes.Buffer
		err := json.NewEncoder(&body).Encode(tt.payload)
		if err != nil {
			t.Fatal(err)
		}

		req := httptest.NewRequest(tt.method, route, &body)
		vars := make(map[string]string)
		vars["userId"] = tt.userId
		req = mux.SetURLVars(req, vars)

		t.Run(tt.name, func(t *testing.T) {
			resp := handler.TodoTasksHanlder(recorder, req)
			if resp != nil {
				if (resp.Error() != "") != tt.wantErr {
					t.Fatalf("%q", resp)
				}
			}
		})
	}
}

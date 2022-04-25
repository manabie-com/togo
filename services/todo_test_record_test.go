package services

import (
	"testing"

	"github.com/SVincentTran/togo/forms"
)

func TestRecordTodoTasks(t *testing.T) {
	mockService := New()

	tests := []struct {
		name      string
		req       forms.TodoTaskRequest
		wantError bool
	}{
		{
			name: "Success",
			req: forms.TodoTaskRequest{
				UserId:     1,
				Title:      "Sample Title",
				Detail:     "Sample Detail",
				RemindDate: "2022-04-25",
			},
			wantError: false,
		},
		{
			name: "UserId Not Existed",
			req: forms.TodoTaskRequest{
				UserId:     4,
				Title:      "Sample Title",
				Detail:     "Sample Detail",
				RemindDate: "2022-04-25",
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := mockService.RecordTodoTasks(tt.req); (err == nil) == tt.wantError {
				t.Fatalf("Want error: %v, got %s", tt.wantError, err)
			}
		})
	}
}

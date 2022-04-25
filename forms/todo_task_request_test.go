package forms

import (
	"testing"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name      string
		req       TodoTaskRequest
		wantError bool
	}{
		{
			name: "Success",
			req: TodoTaskRequest{
				UserId:     1,
				Title:      "Sample Title",
				Detail:     "Sample Detail",
				RemindDate: "2022-04-25",
			},
			wantError: false,
		},
		{
			name: "UserId Empty",
			req: TodoTaskRequest{
				Title:      "Sample Title",
				Detail:     "Sample Detail",
				RemindDate: "2022-04-25",
			},
			wantError: true,
		},
		{
			name: "Title Empty",
			req: TodoTaskRequest{
				UserId:     1,
				Detail:     "Sample Detail",
				RemindDate: "2022-04-25",
			},
			wantError: true,
		},
		{
			name: "Detail Empty",
			req: TodoTaskRequest{
				UserId:     1,
				Title:      "Sample Title",
				RemindDate: "2022-04-25",
			},
			wantError: false,
		},
		{
			name: "Remind Date Empty",
			req: TodoTaskRequest{
				UserId: 1,
				Title:  "Sample Title",
				Detail: "Sample Detail",
			},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.req.Validate(); (err == nil) == tt.wantError {
				t.Fatalf("Want error: %v, got %s", tt.wantError, err)
			}
		})
	}
}

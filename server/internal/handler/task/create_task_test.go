package task

import (
	"context"
	"errors"
	taskEntity "github.com/HoangVyDuong/togo/internal/storages/task"
	taskMock "github.com/HoangVyDuong/togo/internal/usecase/task"
	"github.com/HoangVyDuong/togo/pkg/define"
	taskDTO "github.com/HoangVyDuong/togo/pkg/dtos/task"
	"reflect"
	"testing"
)

func Test_taskHandler_CreateTask(t *testing.T) {
	tests := []struct {
		name         string
		request taskDTO.CreateTaskRequest
		userID uint64
		createTasksError error
		wantErr      error
	}{
		{
			name: "Success Case",
			request: taskDTO.CreateTaskRequest{
				Content: "content",
			},
			userID: 111,
		},
		{
			name: "Empty Content",
			request: taskDTO.CreateTaskRequest{
				Content: "",
			},
			userID: 111,
			wantErr: define.FailedValidation,
		},
		{
			name: "Zero userID",
			request: taskDTO.CreateTaskRequest{
				Content: "",
			},
			userID: 0,
			wantErr: define.FailedValidation,
		},
		{
			name: "Failed Case By DB Error",
			request: taskDTO.CreateTaskRequest{
				Content: "content",
			},
			userID: 111,
			createTasksError: errors.New("error"),
			wantErr: define.Unknown,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			taskServiceMock := &taskMock.ServiceMock{}
			taskServiceMock.CreateTaskFunc = func(ctx context.Context, taskEntity taskEntity.Task) error {
				return tt.createTasksError
			}

			handler := &taskHandler{
				taskService: taskServiceMock,
			}

			ctx := context.WithValue(context.Background(), define.ContextKeyUserID, tt.userID)
			gotResponse, err := handler.CreateTask(ctx, tt.request)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("Auth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil && len(gotResponse.TaskID) == 0 {
				t.Errorf("Auth() gotResponse is empty")
			}
		})
	}
}
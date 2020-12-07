package task

import (
	"context"
	"errors"
	taskEntity "github.com/HoangVyDuong/togo/internal/storages/task"
	taskMock "github.com/HoangVyDuong/togo/internal/usecase/task"
	"github.com/HoangVyDuong/togo/pkg/define"
	"github.com/HoangVyDuong/togo/pkg/dtos"
	taskDTO "github.com/HoangVyDuong/togo/pkg/dtos/task"
	"reflect"
	"testing"
)

func Test_taskHandler_CreateTask(t *testing.T) {
	tests := []struct {
		name         string
		request dtos.EmptyRequest
		createTasksResp int64
		createTasksError error
		wantResponse taskDTO.CreateTaskResponse
		wantErr      error
	}{
		{
			name: "Success Case",
			createTasksResp: 1,
			wantResponse: taskDTO.CreateTaskResponse{
				TaskID: "1",
			},
		},
		{
			name: "Failed Case By Error",
			createTasksError: errors.New("error"),
			wantErr: define.Unknown,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			taskServiceMock := &taskMock.ServiceMock{}
			taskServiceMock.CreateTaskFunc = func(ctx context.Context, taskEntity taskEntity.Task) (int64, error) {
				return tt.createTasksResp, tt.createTasksError
			}

			handler := &taskHandler{
				taskService: taskServiceMock,
			}

			gotResponse, err := handler.CreateTask(context.Background(), taskDTO.CreateTaskRequest{})
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("Auth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("Auth() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}
package task

import (
	"context"
	"errors"
	taskEntity "github.com/HoangVyDuong/togo/internal/storages/task"
	taskMock "github.com/HoangVyDuong/togo/internal/usecase/task"
	"github.com/HoangVyDuong/togo/pkg/dtos"
	taskDTO "github.com/HoangVyDuong/togo/pkg/dtos/task"
	"reflect"
	"testing"
)

func Test_taskHandler_GetTasks(t *testing.T) {
	tests := []struct {
		name         string
		request dtos.EmptyRequest
		getTasksResp []taskEntity.Task
		getTasksError error
		wantResponse taskDTO.Tasks
		wantErr      error
	}{
		{
			name: "Success Case",
			getTasksResp: []taskEntity.Task{
				{
					ID: 1,
					Content: "content",
					UserID: 1,
				},
			},
			wantResponse: taskDTO.Tasks{
				Data: []taskDTO.Task{
					{
						Id: "1",
						Content: "content",
					},
				},
			},
		},
		{
			name: "Failed Case For Error",
			getTasksError: errors.New("error"),
			wantResponse: taskDTO.Tasks{},
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			taskServiceMock := &taskMock.ServiceMock{}
			taskServiceMock.GetTasksFunc = func(ctx context.Context, userId int64) ([]taskEntity.Task, error) {
				return tt.getTasksResp, tt.getTasksError
			}

			handler := &taskHandler{
				taskService: taskServiceMock,
			}

			gotResponse, err := handler.GetTasks(context.Background(), tt.request)
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
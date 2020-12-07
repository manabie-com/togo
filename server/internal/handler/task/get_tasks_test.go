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

func Test_taskHandler_GetTasks(t *testing.T) {
	tests := []struct {
		name         string
		request dtos.EmptyRequest
		userID uint64
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
			userID: 111,
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
			name: "Empty UserID",
			getTasksResp: []taskEntity.Task{
				{
					ID: 1,
					Content: "content",
					UserID: 1,
				},
			},
			userID: 0,
			wantErr: define.FailedValidation,
		},
		{
			name: "Failed Case For DB Error",
			getTasksResp: []taskEntity.Task{
				{
					ID: 1,
					Content: "content",
					UserID: 1,
				},
			},
			userID: 111,
			getTasksError: errors.New("error"),
			wantErr: define.Unknown,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			taskServiceMock := &taskMock.ServiceMock{}
			taskServiceMock.GetTasksFunc = func(ctx context.Context, userId uint64) ([]taskEntity.Task, error) {
				return tt.getTasksResp, tt.getTasksError
			}

			handler := &taskHandler{
				taskService: taskServiceMock,
			}

			ctx := context.WithValue(context.Background(), define.ContextKeyUserID, tt.userID)
			gotResponse, err := handler.GetTasks(ctx, tt.request)
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
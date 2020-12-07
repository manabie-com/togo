package task

import (
	"context"
	taskEntity "github.com/HoangVyDuong/togo/internal/storages/task"
	"github.com/HoangVyDuong/togo/pkg/define"
	"reflect"
	"testing"
	"time"
)

func Test_taskService_CreateTask(t *testing.T) {
	tests := []struct {
		name    string
		request taskEntity.Task
		addTaskErr error
		want    uint64
		wantErr error
	}{
		{
			name: "Success Case",
			request: taskEntity.Task{
				ID:      111,
				Content: "content",
				UserID:  1,
			},
			want: 111,
		},
		{
			name: "zero ID",
			request: taskEntity.Task{
				ID: 0,
				Content: "content",
				UserID: 1,
			},
			wantErr: define.FailedValidation,
		},
		{
			name: "empty content",
			request: taskEntity.Task{
				ID: 1,
				Content: "",
				UserID: 1,
			},
			wantErr: define.FailedValidation,
		},
		{
			name: "zero userID",
			request: taskEntity.Task{
				ID: 1,
				Content: "content",
				UserID: 0,
			},
			wantErr: define.FailedValidation,
		},
		{
			name: "Database Error",
			request: taskEntity.Task{
				ID:      111,
				Content: "content",
				UserID:  1,
			},
			addTaskErr: define.DatabaseError,
			wantErr: define.DatabaseError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			taskRepoMock := &RepositoryMock{}
			taskRepoMock.AddTaskFunc = func(ctx context.Context, taskEntity taskEntity.Task, createdAt time.Time) error {
				return tt.addTaskErr
			}
			ts := &taskService{
				repo: taskRepoMock,
			}
			err := ts.CreateTask(context.Background(), tt.request)
			if !reflect.DeepEqual(err, tt.wantErr){
				t.Errorf("CreateTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_taskService_Delete(t *testing.T) {
	tests := []struct {
		name    string
		taskID uint64
		deleteTaskErr error
		wantErr error
	}{
		{
			name: "Success Case",
			taskID: 111,
		},
		{
			name: "zero ID",
			taskID: 0,
			wantErr: define.FailedValidation,
		},
		{
			name: "Database Error",
			taskID: 111,
			deleteTaskErr: define.DatabaseError,
			wantErr: define.DatabaseError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			taskRepoMock := &RepositoryMock{}
			taskRepoMock.SoftDeleteTaskFunc = func(ctx context.Context, taskId uint64, deletedAt time.Time) error {
				return tt.deleteTaskErr
			}
			ts := &taskService{
				repo: taskRepoMock,
			}
			err := ts.DeleteTask(context.Background(), tt.taskID)
			if !reflect.DeepEqual(err, tt.wantErr){
				t.Errorf("CreateTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_taskService_GetTasks(t *testing.T) {
	tests := []struct {
		name    string
		userID uint64
		retrieveTasksResp []taskEntity.Task
		retrieveTasksErr error
		want []taskEntity.Task
		wantErr error
	}{
		{
			name: "Success Case",
			userID: 111,
			retrieveTasksResp: []taskEntity.Task{
				{
					ID: 111,
					Content: "content",
					UserID: 111,
				},
			},
			want: []taskEntity.Task{
				{
					ID: 111,
					Content: "content",
					UserID: 111,
				},
			},
		},
		{
			name: "Success Case But Empty Response",
			userID: 111,
			retrieveTasksResp: []taskEntity.Task{},
			want: []taskEntity.Task{},
		},
		{
			name: "empty userID",
			userID: 0,
			wantErr: define.FailedValidation,
		},
		{
			name: "Database Error",
			userID: 111,
			retrieveTasksErr: define.DatabaseError,
			wantErr: define.DatabaseError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			taskRepoMock := &RepositoryMock{}
			taskRepoMock.RetrieveTasksFunc = func(ctx context.Context, userId uint64) ([]taskEntity.Task, error) {
				return tt.retrieveTasksResp, tt.retrieveTasksErr
			}
			ts := &taskService{
				repo: taskRepoMock,
			}
			got, err := ts.GetTasks(context.Background(), tt.userID)
			if !reflect.DeepEqual(err, tt.wantErr){
				t.Errorf("CreateTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Auth() gotResponse = %v, want %v", got, tt.want)
			}
		})
	}
}
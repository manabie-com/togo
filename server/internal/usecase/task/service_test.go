package task

import (
	"context"
	taskEntity "github.com/HoangVyDuong/togo/internal/storages/task"
	"github.com/HoangVyDuong/togo/pkg/define"
	"reflect"
	"testing"
)

func Test_taskService_CreateTask(t *testing.T) {
	tests := []struct {
		name    string
		request taskEntity.Task
		addTaskResp int64
		addTaskErr error
		want    int64
		wantErr error
	}{
		{
			name: "Success Case",
			request: taskEntity.Task{
				ID:      111,
				Content: "content",
				UserID:  1,
			},
			addTaskResp: 111,
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
		{
			name: "Not Match TaskID Response",
			request: taskEntity.Task{
				ID:      111,
				Content: "content",
				UserID:  1,
			},
			addTaskResp: 11,
			wantErr: define.Unknown,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			taskRepoMock := &RepositoryMock{}
			taskRepoMock.AddTaskFunc = func(ctx context.Context, taskEntity taskEntity.Task) (int64, error) {
				return tt.addTaskResp, tt.addTaskErr
			}
			ts := &taskService{
				repo: taskRepoMock,
			}
			got, err := ts.CreateTask(context.Background(), tt.request)
			if !reflect.DeepEqual(err, tt.wantErr){
				t.Errorf("CreateTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateTask() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_taskService_Delete(t *testing.T) {
	tests := []struct {
		name    string
		taskID int64
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
			taskRepoMock.SoftDeleteTaskFunc = func(ctx context.Context, taskId int64) error {
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
		userID int64
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
			name: "userID not match",
			userID: 111,
			retrieveTasksResp: []taskEntity.Task{
				{
					ID: 111,
					Content: "content",
					UserID: 111,
				},
			},
			wantErr: define.DatabaseError,
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
			taskRepoMock.RetrieveTasksFunc = func(ctx context.Context, userId int64) ([]taskEntity.Task, error) {
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
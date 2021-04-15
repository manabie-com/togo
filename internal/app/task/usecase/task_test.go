package usecase

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/manabie-com/togo/internal/app/task/model"
	"github.com/manabie-com/togo/internal/app/task/usecase/mock"
	"github.com/manabie-com/togo/internal/util"
	"reflect"
	"testing"
)

func TestTaskService_RetrieveTasks(t1 *testing.T) {
	t1.Parallel()
	tests := []struct {
		name        string
		mockStorage func() TaskStorage
		mockContext func() context.Context
		want        []model.Task
		wantErr     bool
	}{
		{
			name: "unable to get user ID",
			mockStorage: func() TaskStorage {
				ts := mock.NewMockTaskStorage(gomock.NewController(t1))
				ts.EXPECT().RetrieveTasks(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("mock error")).AnyTimes()
				return ts
			},
			mockContext: func() context.Context {
				return context.Background()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "unable to get tasks",
			mockStorage: func() TaskStorage {
				ts := mock.NewMockTaskStorage(gomock.NewController(t1))
				ts.EXPECT().RetrieveTasks(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("mock error")).AnyTimes()
				return ts
			},
			mockContext: func() context.Context {
				return util.SetUserIDToContext(context.Background(), "123")
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success",
			mockStorage: func() TaskStorage {
				ts := mock.NewMockTaskStorage(gomock.NewController(t1))
				ts.EXPECT().RetrieveTasks(gomock.Any(), gomock.Any(), gomock.Any()).Return([]model.Task{
					{
						ID:          "ID",
						Content:     "Content",
						UserID:      "UserID",
						CreatedDate: "CreatedDate",
					},
				}, nil).AnyTimes()
				return ts
			},
			mockContext: func() context.Context {
				return util.SetUserIDToContext(context.Background(), "123")
			},
			want: []model.Task{
				{
					ID:          "ID",
					Content:     "Content",
					UserID:      "UserID",
					CreatedDate: "CreatedDate",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t1.Run(tt.name, func(t1 *testing.T) {
			t1.Parallel()
			t := TaskService{
				taskStorage: tt.mockStorage(),
			}
			got, err := t.RetrieveTasks(tt.mockContext(), "")
			if (err != nil) != tt.wantErr {
				t1.Errorf("RetrieveTasks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("RetrieveTasks() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskService_AddTask(t1 *testing.T) {
	t1.Parallel()
	tests := []struct {
		name        string
		mockStorage func() TaskStorage
		mockContext func() context.Context
		wantErr     bool
	}{
		{
			name: "unable to get user ID",
			mockStorage: func() TaskStorage {
				return nil
			},
			mockContext: func() context.Context {
				return context.Background()
			},
			wantErr: true,
		},
		{
			name: "unable to check the limit",
			mockStorage: func() TaskStorage {
				ts := mock.NewMockTaskStorage(gomock.NewController(t1))
				ts.EXPECT().LimitReached(gomock.Any(), gomock.Any(), gomock.Any()).Return(false, errors.New("mock error")).AnyTimes()
				return ts
			},
			mockContext: func() context.Context {
				return util.SetUserIDToContext(context.Background(), "123")
			},
			wantErr: true,
		},
		{
			name: "limit reached",
			mockStorage: func() TaskStorage {
				ts := mock.NewMockTaskStorage(gomock.NewController(t1))
				ts.EXPECT().LimitReached(gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil).AnyTimes()
				return ts
			},
			mockContext: func() context.Context {
				return util.SetUserIDToContext(context.Background(), "123")
			},
			wantErr: true,
		},
		{
			name: "unable to add task",
			mockStorage: func() TaskStorage {
				ts := mock.NewMockTaskStorage(gomock.NewController(t1))
				ts.EXPECT().LimitReached(gomock.Any(), gomock.Any(), gomock.Any()).Return(false, nil).AnyTimes()
				ts.EXPECT().AddTask(gomock.Any(), gomock.Any()).Return(errors.New("mock error")).AnyTimes()
				return ts
			},
			mockContext: func() context.Context {
				return util.SetUserIDToContext(context.Background(), "123")
			},
			wantErr: true,
		},
		{
			name: "success",
			mockStorage: func() TaskStorage {
				ts := mock.NewMockTaskStorage(gomock.NewController(t1))
				ts.EXPECT().LimitReached(gomock.Any(), gomock.Any(), gomock.Any()).Return(false, nil).AnyTimes()
				ts.EXPECT().AddTask(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
				return ts
			},
			mockContext: func() context.Context {
				return util.SetUserIDToContext(context.Background(), "123")
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t1.Run(tt.name, func(t1 *testing.T) {
			t1.Parallel()
			t := TaskService{
				taskStorage: tt.mockStorage(),
			}
			_, err := t.AddTask(tt.mockContext(), model.Task{})
			if (err != nil) != tt.wantErr {
				t1.Errorf("RetrieveTasks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

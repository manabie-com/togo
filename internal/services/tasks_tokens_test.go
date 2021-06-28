package services

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"

	"github.com/manabie-com/togo/internal/repositories"
	"github.com/manabie-com/togo/internal/tokens"

	repoMock "github.com/manabie-com/togo/internal/repositories/mocks"
	tokenMock "github.com/manabie-com/togo/internal/tokens/mocks"
)

func TestToDoServiceImpl_AddTask(t *testing.T) {
	userRepo := new(repoMock.UserRepo)
	taskRepo := new(repoMock.TaskRepo)
	cachingRepo := new(repoMock.CachingRepo)
	tokenManager := new(tokenMock.TokenManager)

	type fields struct {
		TaskRepo     repositories.TaskRepo
		TokenManager tokens.TokenManager
		CachingRepo  repositories.CachingRepo
		UserRepo     repositories.UserRepo
	}
	type args struct {
		userID string
		task   *repositories.Task
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *repositories.Task
		wantErr bool
	}{
		{
			name: "Test Add Task Success",
			fields: fields{
				TaskRepo:     taskRepo,
				TokenManager: tokenManager,
				CachingRepo:  cachingRepo,
				UserRepo:     userRepo,
			},
			args: args{
				userID: "1234",
				task: &repositories.Task{
					ID:      "1111",
					UserID:  "1234",
					Content: "example",
				},
			},
			want: &repositories.Task{
				ID:        "1111",
				UserID:    "1234",
				Content:   "example",
				CreatedAt: nil,
			},
			wantErr: false,
		},
		{
			name: "Test TaskRepo.AddTask return Error",
			fields: fields{
				TaskRepo:     taskRepo,
				TokenManager: tokenManager,
				CachingRepo:  cachingRepo,
				UserRepo:     userRepo,
			},
			args: args{
				userID: "1234",
				task: &repositories.Task{
					ID:      "1111",
					UserID:  "1233",
					Content: "example",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test Exceed Limit Times to AddTask Per Day",
			fields: fields{
				TaskRepo:     taskRepo,
				TokenManager: tokenManager,
				CachingRepo:  cachingRepo,
				UserRepo:     userRepo,
			},
			args: args{
				userID: "1235",
				task: &repositories.Task{
					ID:      "1111",
					UserID:  "1235",
					Content: "example",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	userRepo.On("GetMaxToDoOfUser", mock.Anything).Return(5, nil)
	cachingRepo.On("Get", buildCachedAddTaskTimesKey("1234")).Return("2", nil)
	cachingRepo.On("Get", buildCachedAddTaskTimesKey("1233")).Return("2", nil)
	cachingRepo.On("Get", buildCachedAddTaskTimesKey("1235")).Return("5", nil)
	cachingRepo.On("Increase", mock.Anything, mock.Anything).Return(nil)

	taskRepo.On("AddTask", &repositories.Task{
		ID:      "1111",
		UserID:  "1234",
		Content: "example",
	}).Return(&repositories.Task{
		ID:        "1111",
		UserID:    "1234",
		Content:   "example",
		CreatedAt: nil,
	}, nil)

	taskRepo.On("AddTask", &repositories.Task{
		ID:      "1111",
		UserID:  "1233",
		Content: "example",
	}).Return(nil, errors.New("add task failed"))

	taskRepo.On("AddTask", &repositories.Task{
		ID:      "1111",
		UserID:  "1235",
		Content: "example",
	}).Return(nil, errors.New("add task failed"))

	//for _, tt := range tests {
	tt0 := tests[0]
	t.Run(tt0.name, func(t *testing.T) {
		s := &ToDoServiceImpl{
			TaskRepo:     tt0.fields.TaskRepo,
			TokenManager: tt0.fields.TokenManager,
			CachingRepo:  tt0.fields.CachingRepo,
			UserRepo:     tt0.fields.UserRepo,
		}
		got, err := s.AddTask(tt0.args.userID, tt0.args.task)
		assert.Nil(t, err)
		if !reflect.DeepEqual(got, tt0.want) {
			t.Errorf("AddTask() got = %v, want %v", got, tt0.want)
		}
	})

	// Test TaskRepo.AddTask return Error
	tt1 := tests[1]
	t.Run(tt1.name, func(t *testing.T) {
		s := &ToDoServiceImpl{
			TaskRepo:     tt1.fields.TaskRepo,
			TokenManager: tt1.fields.TokenManager,
			CachingRepo:  tt1.fields.CachingRepo,
			UserRepo:     tt1.fields.UserRepo,
		}
		got, err := s.AddTask(tt1.args.userID, tt1.args.task)
		assert.Equal(t, err.Error(), "add task failed")
		assert.Nil(t, got)
	})

	// Test Exceed Limit Times to AddTask Per Day
	tt2 := tests[2]
	t.Run(tt2.name, func(t *testing.T) {
		s := &ToDoServiceImpl{
			TaskRepo:     tt2.fields.TaskRepo,
			TokenManager: tt2.fields.TokenManager,
			CachingRepo:  tt2.fields.CachingRepo,
			UserRepo:     tt2.fields.UserRepo,
		}
		got, err := s.AddTask(tt2.args.userID, tt2.args.task)
		assert.Equal(t, err.Error(), "exceed the limited times to add task per day")
		assert.Nil(t, got)
	})

}

func TestToDoServiceImpl_ListTasks(t *testing.T) {
	userRepo := new(repoMock.UserRepo)
	taskRepo := new(repoMock.TaskRepo)
	cachingRepo := new(repoMock.CachingRepo)
	tokenManager := new(tokenMock.TokenManager)

	type fields struct {
		TaskRepo     repositories.TaskRepo
		TokenManager tokens.TokenManager
		CachingRepo  repositories.CachingRepo
		UserRepo     repositories.UserRepo
	}
	type args struct {
		userID    string
		createdAt string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]repositories.Task
		wantErr bool
	}{
		{
			name: "Test ListTask Success",
			fields: fields{
				TaskRepo:     taskRepo,
				TokenManager: tokenManager,
				CachingRepo:  cachingRepo,
				UserRepo:     userRepo,
			},
			args: args{
				userID:    "12345",
				createdAt: "2021-06-28",
			},
			want: &[]repositories.Task{
				{
					ID:        "1111",
					Content:   "example",
					UserID:    "12345",
					CreatedAt: nil,
				},
			},
			wantErr: false,
		},
		{
			name: "Test ListTask Failed",
			fields: fields{
				TaskRepo:     taskRepo,
				TokenManager: tokenManager,
				CachingRepo:  cachingRepo,
				UserRepo:     userRepo,
			},
			args: args{
				userID:    "123456",
				createdAt: "2021-06-28",
			},
			want:    nil,
			wantErr: true,
		},
	}
	taskRepo.On("ListTasks", "12345", "2021-06-28").Return(&[]repositories.Task{
		{
			ID:        "1111",
			Content:   "example",
			UserID:    "12345",
			CreatedAt: nil,
		},
	}, nil)
	taskRepo.On("ListTasks", "123456", "2021-06-28").Return(nil, errors.New("list tasks failed"))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ToDoServiceImpl{
				TaskRepo:     tt.fields.TaskRepo,
				TokenManager: tt.fields.TokenManager,
				CachingRepo:  tt.fields.CachingRepo,
				UserRepo:     tt.fields.UserRepo,
			}
			got, err := s.ListTasks(tt.args.userID, tt.args.createdAt)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListTasks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListTasks() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildCachedAddTaskTimesKey(t *testing.T) {
	assertCorrectResult := func(t testing.TB, got, want string) {
		if got != want {
			t.Errorf("Got: %q. Want: %q", got, want)
		}
	}

	t.Run("Test build key success", func(t *testing.T) {
		userID := "1234"
		got := buildCachedAddTaskTimesKey(userID)
		want := fmt.Sprintf(cachedAddTaskTimesKeyPrefix, userID)
		assertCorrectResult(t, got, want)
	})
}

func TestToDoServiceImpl_GetAuthToken(t *testing.T) {
	userRepo := new(repoMock.UserRepo)
	taskRepo := new(repoMock.TaskRepo)
	cachingRepo := new(repoMock.CachingRepo)

	tokenManager := new(tokenMock.TokenManager)

	type fields struct {
		TaskRepo     repositories.TaskRepo
		TokenManager tokens.TokenManager
		CachingRepo  repositories.CachingRepo
		UserRepo     repositories.UserRepo
	}
	type args struct {
		userID   string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test gen token success",
			fields: fields{
				TaskRepo:     taskRepo,
				TokenManager: tokenManager,
				CachingRepo:  cachingRepo,
				UserRepo:     userRepo,
			},
			args: args{
				userID:   "12345",
				password: "12345",
			},
			want:    "123123123",
			wantErr: false,
		},
		{
			name: "Test gen token failed",
			fields: fields{
				TaskRepo:     taskRepo,
				TokenManager: tokenManager,
				CachingRepo:  cachingRepo,
				UserRepo:     userRepo,
			},
			args: args{
				userID:   "123456",
				password: "12345",
			},

			want:    "",
			wantErr: true,
		},
	}
	tokenManager.On("GetAuthToken", "12345", "12345").Return("123123123", nil)
	tokenManager.On("GetAuthToken", "123456", "12345").Return("", errors.New("test failed"))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ToDoServiceImpl{
				TaskRepo:     tt.fields.TaskRepo,
				TokenManager: tt.fields.TokenManager,
				CachingRepo:  tt.fields.CachingRepo,
				UserRepo:     tt.fields.UserRepo,
			}
			got, err := s.GetAuthToken(tt.args.userID, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAuthToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAuthToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToDoServiceImpl_ValidToken(t *testing.T) {
	userRepo := new(repoMock.UserRepo)
	taskRepo := new(repoMock.TaskRepo)
	cachingRepo := new(repoMock.CachingRepo)
	tokenManager := new(tokenMock.TokenManager)

	type fields struct {
		TaskRepo     repositories.TaskRepo
		TokenManager tokens.TokenManager
		CachingRepo  repositories.CachingRepo
		UserRepo     repositories.UserRepo
	}
	type args struct {
		token string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantUserID string
		wantValid  bool
	}{
		{
			name: "Test Token Valid",
			fields: fields{
				TaskRepo:     taskRepo,
				TokenManager: tokenManager,
				CachingRepo:  cachingRepo,
				UserRepo:     userRepo,
			},
			args: args{
				token: "12345",
			},
			wantUserID: "11111",
			wantValid:  true,
		},
		{
			name: "Test Token Invalid",
			fields: fields{
				TaskRepo:     taskRepo,
				TokenManager: tokenManager,
				CachingRepo:  cachingRepo,
				UserRepo:     userRepo,
			},
			args: args{
				token: "123456",
			},
			wantUserID: "",
			wantValid:  false,
		},
	}
	tokenManager.On("ValidToken", "12345").Return("11111", true)
	tokenManager.On("ValidToken", "123456").Return("", false)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ToDoServiceImpl{
				TaskRepo:     tt.fields.TaskRepo,
				TokenManager: tt.fields.TokenManager,
				CachingRepo:  tt.fields.CachingRepo,
				UserRepo:     tt.fields.UserRepo,
			}
			gotUserID, gotValid := s.ValidToken(tt.args.token)
			if gotUserID != tt.wantUserID {
				t.Errorf("ValidToken() gotUserID = %v, want %v", gotUserID, tt.wantUserID)
			}
			if gotValid != tt.wantValid {
				t.Errorf("ValidToken() gotValid = %v, want %v", gotValid, tt.wantValid)
			}
		})
	}
}

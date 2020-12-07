package user

import (
	"context"
	"github.com/HoangVyDuong/togo/pkg/define"
	"reflect"
	"testing"
	"time"
)

func Test_userService_IncreaseTaskTimesPerDuration(t *testing.T) {
	type args struct {
		userId   string
		duration time.Duration
	}
	tests := []struct {
		name    string
		args args
		increaseTaskResp int
		increaseTaskErr error
		want    int
		wantErr error
	}{
		{
			name: "Success Case",
			args: args{
				userId: "111",
				duration: time.Second * 86400,
			},
			increaseTaskResp: 1,
			want: 1,
		},
		{
			name: "Empty userId",
			args: args{
				userId: "",
			},
			wantErr: define.FailedValidation,
		},
		{
			name: "Cache error",
			args: args{
				userId: "111",
			},
			increaseTaskErr: define.CacheError,
			wantErr: define.CacheError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userCacheMock := &CacheMock{}
			userCacheMock.IncreaseTaskFunc = func(ctx context.Context, userId string, time time.Duration) (int, error) {
				return tt.increaseTaskResp, tt.increaseTaskErr
			}
			s := &userService{
				cache: userCacheMock,
			}
			got, err := s.IncreaseTaskTimesPerDuration(context.Background(), tt.args.userId, tt.args.duration)
			if !reflect.DeepEqual(err, tt.wantErr){
				t.Errorf("IncreaseTaskTimesPerDuration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IncreaseTaskTimesPerDuration() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_IsOverLimitTask(t *testing.T) {
	type args struct {
		userId   string
		limit int
	}
	tests := []struct {
		name    string
		args args
		checkLimitResp bool
		checkLimitErr error
		want    bool
		wantErr error
	}{
		{
			name: "Success Case",
			args: args{
				userId: "111",
				limit: 5,
			},
			checkLimitResp: true,
			want: true,
		},
		{
			name: "Empty userId",
			args: args{
				userId: "",
				limit: 5,
			},
			wantErr: define.FailedValidation,
		},
		{
			name: "zero limit",
			args: args{
				userId: "111",
				limit: 0,
			},
			wantErr: define.FailedValidation,
		},
		{
			name: "Cache error",
			args: args{
				userId: "111",
			},
			checkLimitErr: define.CacheError,
			wantErr: define.CacheError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userCacheMock := &CacheMock{}
			userCacheMock.CheckLimitFunc = func(ctx context.Context, userKey string, limit int) (bool, error) {
				return tt.checkLimitResp, tt.checkLimitErr
			}
			s := &userService{
				cache: userCacheMock,
			}
			got, err := s.IsOverLimitTask(context.Background(), tt.args.userId, tt.args.limit)
			if !reflect.DeepEqual(err, tt.wantErr){
				t.Errorf("IncreaseTaskTimesPerDuration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IncreaseTaskTimesPerDuration() got = %v, want %v", got, tt.want)
			}
		})
	}
}
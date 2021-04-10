package auth

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"

	"github.com/manabie-com/togo/internal/services/users"
)

func Test_genHash(t *testing.T) {
	type args struct {
		pwd []byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{pwd: []byte("123456789")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := genHash(tt.args.pwd)
			if (err != nil) != tt.wantErr {
				t.Errorf("genHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("genHash() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authService_login(t *testing.T) {
	type fields struct {
		userRepo users.UserRepo
	}
	type args struct {
		ctx      context.Context
		username string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "happy case",
			fields: fields{
				userRepo: func() users.UserRepo {
					ctrl := gomock.NewController(t)
					mockUserRepo := users.NewMockUserRepo(ctrl)
					mockUserRepo.EXPECT().GetByUsername(gomock.Any(), gomock.Any()).Return(&users.User{
						ID:           uuid.MustParse("456020ea-257c-4066-8b46-b5b186b2335d"),
						Username:     "zahj",
						Password:     "$2a$04$.jJLqmqcp0f0mHBWaZ0Tv.I10qbTmivct1yTmhdPauBIIiNW.Ti0K",
						CreatedAt:    time.Now(),
						UpdatedAt:    time.Now(),
						OldPasswords: []string{},
					}, nil).AnyTimes()
					return mockUserRepo
				}(),
			},
			args: args{
				ctx:      context.Background(),
				username: "zahj",
				password: "123456789",
			},
			want:    "",
			wantErr: false,
		},

		{
			name: "wrong username",
			fields: fields{
				userRepo: func() users.UserRepo {
					ctrl := gomock.NewController(t)
					mockUserRepo := users.NewMockUserRepo(ctrl)
					mockUserRepo.EXPECT().GetByUsername(gomock.Any(), gomock.Any()).Return(nil, nil).AnyTimes()
					return mockUserRepo
				}(),
			},
			args: args{
				ctx:      context.Background(),
				username: "zahjxxx",
				password: "123456789",
			},
			want:    "",
			wantErr: true,
		},

		{
			name: "wrong password",
			fields: fields{
				userRepo: func() users.UserRepo {
					ctrl := gomock.NewController(t)
					mockUserRepo := users.NewMockUserRepo(ctrl)
					mockUserRepo.EXPECT().GetByUsername(gomock.Any(), gomock.Any()).Return(&users.User{
						ID:           uuid.MustParse("456020ea-257c-4066-8b46-b5b186b2335d"),
						Username:     "zahj",
						Password:     "$2a$04$.jJLqmqcp0f0mHBWaZ0Tv.I10qbTmivct1yTmhdPauBIIiNW.Ti0K",
						CreatedAt:    time.Now(),
						UpdatedAt:    time.Now(),
						OldPasswords: []string{},
					}, nil).AnyTimes()
					return mockUserRepo
				}(),
			},
			args: args{
				ctx:      context.Background(),
				username: "zahj",
				password: "123456788",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				userRepo: tt.fields.userRepo,
			}
			got, err := s.login(tt.args.ctx, tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Login() got = %v, want %v", got, tt.want)
			}
		})
	}
}

package services

import (
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/manabie-com/togo/internal/storages/postgres/mocks"

	"github.com/manabie-com/togo/internal/storages/postgres"
	"github.com/manabie-com/togo/pkg/model"
)

func Test_userService_hashPassword(t *testing.T) {

	type args struct {
		rawPass string
		salt1   string
		salt2   string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "same salt",
			args: args{
				rawPass: "1234",
				salt2:   "1234",
				salt1:   "1234",
			},
		},
		{
			name: "same salt",
			args: args{
				rawPass: "1234",
				salt2:   "1234",
				salt1:   "1235",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &userService{}
			got1 := s.hashPassword(tt.args.rawPass, tt.args.salt1)
			got2 := s.hashPassword(tt.args.rawPass, tt.args.salt2)
			if tt.args.salt1 == tt.args.salt2 && got1 != got2 {
				t.Errorf("hash err, return 2 diff hash from same salt")
			}
			if tt.args.salt1 != tt.args.salt2 && got1 == got2 {
				t.Errorf("hash err, return 2 same hash with diff salt")
			}
		})
	}
}

func Test_userService_Authentication(t *testing.T) {
	type fields struct {
		repo   postgres.Repository
		jwtKey string
	}
	type args struct {
		user *model.User
	}
	tests := []struct {
		name   string
		fields fields
		args   args

		wantErr  bool
		mockRepo func(mock *mocks.Repository, service *userService)
	}{
		// TODO: Add test cases.
		{
			name: "no user",
			fields: fields{
				repo:   &mocks.Repository{},
				jwtKey: "123456",
			},
			args: args{
				user: &model.User{
					UserName: "testUserName",
					Password: "1234",
				},
			},

			wantErr: true,
			mockRepo: func(m *mocks.Repository, service *userService) {
				m.On("GetUser", "testUserName").Return(nil, gorm.ErrRecordNotFound)
			},
		},
		{
			name: "found user",
			fields: fields{
				repo:   &mocks.Repository{},
				jwtKey: "123456",
			},
			args: args{
				user: &model.User{
					UserName: "testUserName",
					Password: "1234",
				},
			},

			wantErr: false,
			mockRepo: func(m *mocks.Repository, service *userService) {
				pw := service.hashPassword("1234", "1234")
				m.On("GetUser", "testUserName").Return(&model.User{
					UserName: "testUserName",
					Password: pw,
					Salt:     "1234",
				}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			s := &userService{
				repo:   tt.fields.repo,
				jwtKey: tt.fields.jwtKey,
			}
			tt.mockRepo(tt.fields.repo.(*mocks.Repository), s)
			got, err := s.Authentication(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Authentication() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got == "" {
				t.Errorf("empty jwt")
			}
		})
	}
}

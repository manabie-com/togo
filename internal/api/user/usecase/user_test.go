package usecase

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/manabie-com/togo/internal/api/user/storages/mock"
	"testing"

	"github.com/manabie-com/togo/internal/api/user/storages"
)

func TestUser_IsValidate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	storage := mock.NewMockStore(ctrl)

	type fields struct {
		Store storages.Store
	}
	type args struct {
		ctx      context.Context
		userID   string
		password string
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		mockStoreGet func()
		want         bool
		wantErr      bool
	}{
		{
			name:   "invalid",
			fields: fields{Store: storage},
			args: args{
				ctx:      nil,
				userID:   "test1",
				password: "example",
			},
			mockStoreGet: func() {
				storage.EXPECT().Get(gomock.Any(), "test1").
					Return(&storages.User{
						ID:       "test1",
						Password: "$2a$14$BdgOuNVBU7sdGW9rIDIIv.MWXDdvTVKyTppb3qW03bmvz/6hhA1FO",
						MaxTodo:  5,
					}, nil)
			},
			want:    true,
			wantErr: false,
		},
		{
			name:   "invalid",
			fields: fields{Store: storage},
			args: args{
				ctx:      nil,
				userID:   "test1",
				password: "random",
			},
			mockStoreGet: func() {
				storage.EXPECT().Get(gomock.Any(), "test1").
					Return(&storages.User{
						ID:       "test1",
						Password: "$2a$14$BdgOuNVBU7sdGW9rIDIIv.MWXDdvTVKyTppb3qW03bmvz/6hhA1FO",
						MaxTodo:  5,
					}, nil)
			},
			want:    false,
			wantErr: false,
		},
		{
			name:   "storages error",
			fields: fields{Store: storage},
			args: args{
				ctx:      nil,
				userID:   "test1",
				password: "random",
			},
			mockStoreGet: func() {
				storage.EXPECT().Get(gomock.Any(), "test1").
					Return(nil, errors.New("storage got error"))
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockStoreGet()
			s := &User{
				Store: tt.fields.Store,
			}
			got, err := s.IsValidate(tt.args.ctx, tt.args.userID, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.IsValidate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("User.IsValidate() = %v, want %v", got, tt.want)
			}
		})
	}
}

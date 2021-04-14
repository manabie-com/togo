package user_tasks

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/looplab/eventhorizon/mocks"

	"github.com/golang/mock/gomock"

	"github.com/google/uuid"
	"github.com/looplab/eventhorizon"
	"github.com/manabie-com/togo/internal/services/users"
)

func Test_service_CreateTask(t *testing.T) {
	timeNow := time.Now()
	userID := uuid.MustParse("456020ea-257c-4066-8b46-b5b186b2335d")

	type fields struct {
		commandBus     eventhorizon.CommandHandler
		userConfigRepo users.UserConfigRepo
	}
	type args struct {
		ctx     context.Context
		userID  uuid.UUID
		content string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "happy case",
			fields: fields{
				commandBus: &mocks.CommandHandler{},
				userConfigRepo: func() users.UserConfigRepo {
					ctrl := gomock.NewController(t)
					mockUserConfigRepo := users.NewMockUserConfigRepo(ctrl)
					mockUserConfigRepo.EXPECT().GetByUserID(gomock.Any(), gomock.Any()).Return(&users.UserConfig{
						ID:        10,
						UserID:    userID,
						TaskLimit: 5,
						IsActive:  true,
						CreatedAt: timeNow,
						UpdatedAt: timeNow,
					}, nil).AnyTimes()
					return mockUserConfigRepo
				}(),
			},
			args: args{
				ctx:     context.Background(),
				userID:  userID,
				content: "hello",
			},
			wantErr: false,
		},

		{
			name: "user is not configed",
			fields: fields{
				commandBus: &mocks.CommandHandler{},
				userConfigRepo: func() users.UserConfigRepo {
					ctrl := gomock.NewController(t)
					mockUserConfigRepo := users.NewMockUserConfigRepo(ctrl)
					mockUserConfigRepo.EXPECT().GetByUserID(gomock.Any(), gomock.Any()).Return(nil, nil).AnyTimes()
					return mockUserConfigRepo
				}(),
			},
			args: args{
				ctx:     context.Background(),
				userID:  userID,
				content: "hello",
			},
			wantErr: true,
		},

		{
			name: "handle command failed by business error",
			fields: fields{
				commandBus: &mocks.CommandHandler{
					Err: eventhorizon.AggregateError{Err: errors.New("test")},
				},
				userConfigRepo: func() users.UserConfigRepo {
					ctrl := gomock.NewController(t)
					mockUserConfigRepo := users.NewMockUserConfigRepo(ctrl)
					mockUserConfigRepo.EXPECT().GetByUserID(gomock.Any(), gomock.Any()).Return(&users.UserConfig{
						ID:        10,
						UserID:    userID,
						TaskLimit: 5,
						IsActive:  true,
						CreatedAt: timeNow,
						UpdatedAt: timeNow,
					}, nil).AnyTimes()
					return mockUserConfigRepo
				}(),
			},
			args: args{
				ctx:     context.Background(),
				userID:  userID,
				content: "hello",
			},
			wantErr: true,
		},

		{
			name: "handle command failed by system error",
			fields: fields{
				commandBus: &mocks.CommandHandler{
					Err: errors.New("test"),
				},
				userConfigRepo: func() users.UserConfigRepo {
					ctrl := gomock.NewController(t)
					mockUserConfigRepo := users.NewMockUserConfigRepo(ctrl)
					mockUserConfigRepo.EXPECT().GetByUserID(gomock.Any(), gomock.Any()).Return(&users.UserConfig{
						ID:        10,
						UserID:    userID,
						TaskLimit: 5,
						IsActive:  true,
						CreatedAt: timeNow,
						UpdatedAt: timeNow,
					}, nil).AnyTimes()
					return mockUserConfigRepo
				}(),
			},
			args: args{
				ctx:     context.Background(),
				userID:  userID,
				content: "hello",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				commandBus:     tt.fields.commandBus,
				userConfigRepo: tt.fields.userConfigRepo,
			}
			if err := s.CreateTask(tt.args.ctx, tt.args.userID, tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("CreateTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

package usecase

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/manabie-com/togo/internal/api/task/storages/mock"
	mock2 "github.com/manabie-com/togo/internal/api/user/storages/mock"
	"github.com/manabie-com/togo/internal/api/utils"
	"reflect"
	"testing"

	"github.com/manabie-com/togo/internal/api/task/storages"
	userStorages "github.com/manabie-com/togo/internal/api/user/storages"
)

func TestTask_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	storage := mock.NewMockStore(ctrl)

	type fields struct {
		Store storages.Store
	}
	type args struct {
		ctx         context.Context
		userID      string
		createdDate string
		page        int
		limit       int
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		mockStoreList func()
		want          []*storages.Task
		wantErr       bool
	}{
		{
			name:   "empty userID",
			fields: fields{Store: storage},
			args: args{
				ctx:         nil,
				userID:      "",
				createdDate: "2021-07-12",
				page:        0,
				limit:       0,
			},
			mockStoreList: func() {
				storage.EXPECT().RetrieveTasks(gomock.Any(), "", "2021-07-12", 0, 0).
					Return(nil, errors.New("storage got error"))
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:   "empty createDate",
			fields: fields{Store: storage},
			args: args{
				ctx:         nil,
				userID:      "test1",
				createdDate: "",
				page:        0,
				limit:       0,
			},
			mockStoreList: func() {
				storage.EXPECT().RetrieveTasks(gomock.Any(), "test1", "", 0, 0).
					Return(nil, errors.New("storage got error"))
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:   "empty userID and createDate",
			fields: fields{Store: storage},
			args: args{
				ctx:         nil,
				userID:      "",
				createdDate: "",
				page:        0,
				limit:       0,
			},
			mockStoreList: func() {
				storage.EXPECT().RetrieveTasks(gomock.Any(), "", "", 0, 0).
					Return(nil, errors.New("storage got error"))
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:   "negative page and limit",
			fields: fields{Store: storage},
			args: args{
				ctx:         nil,
				userID:      "test1",
				createdDate: "2021-07-12",
				page:        -1,
				limit:       -1,
			},
			mockStoreList: func() {
				storage.EXPECT().RetrieveTasks(gomock.Any(), "test1", "2021-07-12", -1, -1).
					Return([]*storages.Task{{
						ID:          "1",
						Content:     "a",
						UserID:      "test1",
						CreatedDate: "2021-07-12",
					}, {
						ID:          "2",
						Content:     "b",
						UserID:      "test1",
						CreatedDate: "2021-07-12",
					}, {
						ID:          "3",
						Content:     "c",
						UserID:      "test1",
						CreatedDate: "2021-07-12",
					}}, nil)
			},
			want: []*storages.Task{{
				ID:          "1",
				Content:     "a",
				UserID:      "test1",
				CreatedDate: "2021-07-12",
			}, {
				ID:          "2",
				Content:     "b",
				UserID:      "test1",
				CreatedDate: "2021-07-12",
			}, {
				ID:          "3",
				Content:     "c",
				UserID:      "test1",
				CreatedDate: "2021-07-12",
			}},
			wantErr: false,
		},
		{
			name:   "normal case",
			fields: fields{Store: storage},
			args: args{
				ctx:         nil,
				userID:      "test1",
				createdDate: "2021-07-12",
				page:        0,
				limit:       0,
			},
			mockStoreList: func() {
				storage.EXPECT().RetrieveTasks(gomock.Any(), "test1", "2021-07-12", 0, 0).
					Return([]*storages.Task{{
						ID:          "1",
						Content:     "a",
						UserID:      "test1",
						CreatedDate: "2021-07-12",
					}, {
						ID:          "2",
						Content:     "b",
						UserID:      "test1",
						CreatedDate: "2021-07-12",
					}, {
						ID:          "3",
						Content:     "c",
						UserID:      "test1",
						CreatedDate: "2021-07-12",
					}}, nil)
			},
			want: []*storages.Task{{
				ID:          "1",
				Content:     "a",
				UserID:      "test1",
				CreatedDate: "2021-07-12",
			}, {
				ID:          "2",
				Content:     "b",
				UserID:      "test1",
				CreatedDate: "2021-07-12",
			}, {
				ID:          "3",
				Content:     "c",
				UserID:      "test1",
				CreatedDate: "2021-07-12",
			}},
			wantErr: false,
		},
		{
			name:   "set limit and page case",
			fields: fields{Store: storage},
			args: args{
				ctx:         nil,
				userID:      "test1",
				createdDate: "2021-07-12",
				page:        0,
				limit:       1,
			},
			mockStoreList: func() {
				storage.EXPECT().RetrieveTasks(gomock.Any(), "test1", "2021-07-12", 0, 1).
					Return([]*storages.Task{{
						ID:          "1",
						Content:     "a",
						UserID:      "test1",
						CreatedDate: "2021-07-12",
					}}, nil)
			},
			want: []*storages.Task{{
				ID:          "1",
				Content:     "a",
				UserID:      "test1",
				CreatedDate: "2021-07-12",
			}},
			wantErr: false,
		},
		{
			name:   "storage got error",
			fields: fields{Store: storage},
			args: args{
				ctx:         nil,
				userID:      "test1",
				createdDate: "2021-07-12",
				page:        0,
				limit:       0,
			},
			mockStoreList: func() {
				storage.EXPECT().RetrieveTasks(gomock.Any(), "test1", "2021-07-12", 0, 0).
					Return(nil, errors.New("storge got error"))
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockStoreList()
			s := &Task{
				Store: tt.fields.Store,
			}
			got, err := s.List(tt.args.ctx, tt.args.userID, tt.args.createdDate, tt.args.page, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("Task.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Task.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTask_Add(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	taskStorage := mock.NewMockStore(ctrl)
	userStorage := mock2.NewMockStore(ctrl)

	type fields struct {
		Store     storages.Store
		UserStore userStorages.Store
		GenFn     utils.GenerateNewUUIDFn
	}
	type args struct {
		ctx    context.Context
		userID string
		task   *storages.Task
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		mockStoreGetUser func()
		mockStoreList    func()
		mockStoreAdd     func()
		want             *storages.Task
		wantErr          bool
	}{
		{
			name: "user is empty",
			fields: fields{
				Store:     taskStorage,
				UserStore: userStorage,
				GenFn: func() string {
					return "id_test"
				},
			},
			args: args{
				ctx:    nil,
				userID: "",
				task:   &storages.Task{Content: "abc"},
			},
			mockStoreGetUser: func() {
				userStorage.EXPECT().Get(gomock.Any(), "").
					Return(nil, errors.New("storage got error"))
			},
			mockStoreList: nil,
			mockStoreAdd:  nil,
			want:          nil,
			wantErr:       true,
		},
		{
			name: "task is empty",
			fields: fields{
				Store:     taskStorage,
				UserStore: userStorage,
				GenFn: func() string {
					return "id_test"
				},
			},
			args: args{
				ctx:    nil,
				userID: "test1",
				task:   &storages.Task{},
			},
			mockStoreGetUser: func() {
				userStorage.EXPECT().Get(gomock.Any(), "test1").
					Return(&userStorages.User{
						ID:       "test1",
						Password: "$2a$14$BdgOuNVBU7sdGW9rIDIIv.MWXDdvTVKyTppb3qW03bmvz/6hhA1FO",
						MaxTodo:  5,
					}, nil)
			},
			mockStoreList: func() {
				taskStorage.EXPECT().RetrieveTasks(gomock.Any(), "test1", utils.GetTimeNowWithDefaultLayoutInString(), 0, 6).
					Return([]*storages.Task{{
						ID:          "1",
						Content:     "a",
						UserID:      "test1",
						CreatedDate: utils.GetTimeNowWithDefaultLayoutInString(),
					}, {
						ID:          "2",
						Content:     "b",
						UserID:      "test1",
						CreatedDate: utils.GetTimeNowWithDefaultLayoutInString(),
					}, {
						ID:          "3",
						Content:     "c",
						UserID:      "test1",
						CreatedDate: utils.GetTimeNowWithDefaultLayoutInString(),
					}}, nil)
			},
			mockStoreAdd: func() {
				taskStorage.EXPECT().AddTask(gomock.Any(), &storages.Task{
					ID:          "id_test",
					Content:     "",
					UserID:      "test1",
					CreatedDate: utils.GetTimeNowWithDefaultLayoutInString(),
				}).
					Return(errors.New("storage got error"))
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "task is empty",
			fields: fields{
				Store:     taskStorage,
				UserStore: userStorage,
				GenFn: func() string {
					return "id_test"
				},
			},
			args: args{
				ctx:    nil,
				userID: "test1",
				task:   &storages.Task{},
			},
			mockStoreGetUser: func() {
				userStorage.EXPECT().Get(gomock.Any(), "test1").
					Return(&userStorages.User{
						ID:       "test1",
						Password: "$2a$14$BdgOuNVBU7sdGW9rIDIIv.MWXDdvTVKyTppb3qW03bmvz/6hhA1FO",
						MaxTodo:  5,
					}, nil)
			},
			mockStoreList: func() {
				taskStorage.EXPECT().RetrieveTasks(gomock.Any(), "test1", utils.GetTimeNowWithDefaultLayoutInString(), 0, 6).
					Return([]*storages.Task{{
						ID:          "1",
						Content:     "a",
						UserID:      "test1",
						CreatedDate: utils.GetTimeNowWithDefaultLayoutInString(),
					}, {
						ID:          "2",
						Content:     "b",
						UserID:      "test1",
						CreatedDate: utils.GetTimeNowWithDefaultLayoutInString(),
					}, {
						ID:          "3",
						Content:     "c",
						UserID:      "test1",
						CreatedDate: utils.GetTimeNowWithDefaultLayoutInString(),
					}}, nil)
			},
			mockStoreAdd: func() {
				taskStorage.EXPECT().AddTask(gomock.Any(), &storages.Task{
					ID:          "id_test",
					Content:     "",
					UserID:      "test1",
					CreatedDate: utils.GetTimeNowWithDefaultLayoutInString(),
				}).
					Return(errors.New("storage got error"))
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "list task error",
			fields: fields{
				Store:     taskStorage,
				UserStore: userStorage,
				GenFn: func() string {
					return "id_test"
				},
			},
			args: args{
				ctx:    nil,
				userID: "test1",
				task:   &storages.Task{},
			},
			mockStoreGetUser: func() {
				userStorage.EXPECT().Get(gomock.Any(), "test1").
					Return(&userStorages.User{
						ID:       "test1",
						Password: "$2a$14$BdgOuNVBU7sdGW9rIDIIv.MWXDdvTVKyTppb3qW03bmvz/6hhA1FO",
						MaxTodo:  5,
					}, nil)
			},
			mockStoreList: func() {
				taskStorage.EXPECT().RetrieveTasks(gomock.Any(), "test1", utils.GetTimeNowWithDefaultLayoutInString(), 0, 6).
					Return(nil, errors.New("storage got error"))
			},
			mockStoreAdd: nil,
			want:         nil,
			wantErr:      true,
		},
		{
			name: "reach limit",
			fields: fields{
				Store:     taskStorage,
				UserStore: userStorage,
				GenFn: func() string {
					return "id_test"
				},
			},
			args: args{
				ctx:    nil,
				userID: "test1",
				task:   &storages.Task{},
			},
			mockStoreGetUser: func() {
				userStorage.EXPECT().Get(gomock.Any(), "test1").
					Return(&userStorages.User{
						ID:       "test1",
						Password: "$2a$14$BdgOuNVBU7sdGW9rIDIIv.MWXDdvTVKyTppb3qW03bmvz/6hhA1FO",
						MaxTodo:  5,
					}, nil)
			},
			mockStoreList: func() {
				taskStorage.EXPECT().RetrieveTasks(gomock.Any(), "test1", utils.GetTimeNowWithDefaultLayoutInString(), 0, 6).
					Return([]*storages.Task{{
						ID:          "1",
						Content:     "a",
						UserID:      "test1",
						CreatedDate: utils.GetTimeNowWithDefaultLayoutInString(),
					}, {
						ID:          "2",
						Content:     "b",
						UserID:      "test1",
						CreatedDate: utils.GetTimeNowWithDefaultLayoutInString(),
					}, {
						ID:          "3",
						Content:     "c",
						UserID:      "test1",
						CreatedDate: utils.GetTimeNowWithDefaultLayoutInString(),
					}, {
						ID:          "4",
						Content:     "c",
						UserID:      "test1",
						CreatedDate: utils.GetTimeNowWithDefaultLayoutInString(),
					}, {
						ID:          "5",
						Content:     "c",
						UserID:      "test1",
						CreatedDate: utils.GetTimeNowWithDefaultLayoutInString(),
					}}, nil)
			},
			mockStoreAdd: nil,
			want:         nil,
			wantErr:      true,
		},
		{
			name: "normal case",
			fields: fields{
				Store:     taskStorage,
				UserStore: userStorage,
				GenFn: func() string {
					return "id_test"
				},
			},
			args: args{
				ctx:    nil,
				userID: "test1",
				task:   &storages.Task{Content: "abc"},
			},
			mockStoreGetUser: func() {
				userStorage.EXPECT().Get(gomock.Any(), "test1").
					Return(&userStorages.User{
						ID:       "test1",
						Password: "$2a$14$BdgOuNVBU7sdGW9rIDIIv.MWXDdvTVKyTppb3qW03bmvz/6hhA1FO",
						MaxTodo:  5,
					}, nil)
			},
			mockStoreList: func() {
				taskStorage.EXPECT().RetrieveTasks(gomock.Any(), "test1", utils.GetTimeNowWithDefaultLayoutInString(), 0, 6).
					Return([]*storages.Task{{
						ID:          "1",
						Content:     "a",
						UserID:      "test1",
						CreatedDate: utils.GetTimeNowWithDefaultLayoutInString(),
					}, {
						ID:          "2",
						Content:     "b",
						UserID:      "test1",
						CreatedDate: utils.GetTimeNowWithDefaultLayoutInString(),
					}, {
						ID:          "3",
						Content:     "c",
						UserID:      "test1",
						CreatedDate: utils.GetTimeNowWithDefaultLayoutInString(),
					}, {
						ID:          "4",
						Content:     "c",
						UserID:      "test1",
						CreatedDate: utils.GetTimeNowWithDefaultLayoutInString(),
					}}, nil)
			},
			mockStoreAdd: func() {
				taskStorage.EXPECT().AddTask(gomock.Any(), &storages.Task{
					ID:          "id_test",
					Content:     "abc",
					UserID:      "test1",
					CreatedDate: utils.GetTimeNowWithDefaultLayoutInString(),
				}).
					Return(nil)
			},
			want: &storages.Task{
				ID:          "id_test",
				Content:     "abc",
				UserID:      "test1",
				CreatedDate: utils.GetTimeNowWithDefaultLayoutInString(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		if tt.mockStoreGetUser != nil {
			tt.mockStoreGetUser()
		}
		if tt.mockStoreList != nil {
			tt.mockStoreList()
		}
		if tt.mockStoreAdd != nil {
			tt.mockStoreAdd()
		}

		t.Run(tt.name, func(t *testing.T) {
			s := &Task{
				Store:           tt.fields.Store,
				UserStore:       tt.fields.UserStore,
				GeneratorUUIDFn: tt.fields.GenFn,
			}
			got, err := s.Add(tt.args.ctx, tt.args.userID, tt.args.task)
			if (err != nil) != tt.wantErr {
				t.Errorf("Task.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && tt.want != nil {
				if got.Content != tt.want.Content && got.CreatedDate != tt.want.CreatedDate && got.UserID != tt.want.UserID {
					t.Errorf("Task.Add() = %v, want %v", got, tt.want)
				}
			} else if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Task.Add() = %v, want %v", got, tt.want)
			}
		})

	}
}

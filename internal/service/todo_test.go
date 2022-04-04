package service

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vchitai/togo/configs"
	"github.com/vchitai/togo/internal/mapper"
	"github.com/vchitai/togo/internal/models"
	"github.com/vchitai/togo/internal/store"
	"github.com/vchitai/togo/internal/utils"
	mockStore "github.com/vchitai/togo/mocks/store"
	"github.com/vchitai/togo/pb"
)

var (
	someRandomDay   = utils.RoundDate(time.Now().Add(7 * time.Hour))
	someRandomError = fmt.Errorf("an error")
)

func TestServerImplAddToDoListRequestValidate(t *testing.T) {
	for _, tc := range []struct {
		name     string
		req      *pb.AddToDoListRequest
		expected error
	}{
		{
			name: "happy case",
			req: &pb.AddToDoListRequest{
				UserId: "user-id",
				Entry: []*pb.ToDoEntry{
					{
						Content: "content",
					},
				},
			},
			expected: nil,
		},
		{
			name: "userID empty",
			req: &pb.AddToDoListRequest{
				UserId: "",
				Entry: []*pb.ToDoEntry{
					{
						Content: "content",
					},
				},
			},
			expected: fmt.Errorf("invalid AddToDoListRequest.UserId: value length must be at least 1 runes"),
		},
		{
			name: "entry empty",
			req: &pb.AddToDoListRequest{
				UserId: "user-id",
				Entry:  nil,
			},
			expected: fmt.Errorf("invalid AddToDoListRequest.Entry: value must contain at least 1 item(s)"),
		},
		{
			name: "entry empty",
			req: &pb.AddToDoListRequest{
				UserId: "user-id",
				Entry: []*pb.ToDoEntry{
					{
						Content: "",
					},
				},
			},
			expected: fmt.Errorf("invalid AddToDoListRequest.Entry[0]: embedded message failed validation | caused by: invalid ToDoEntry.Content: value length must be at least 1 runes"),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if tc.expected != nil {
				assert.EqualError(t, tc.req.Validate(), tc.expected.Error())
			} else {
				assert.NoError(t, tc.req.Validate())
			}
		})
	}
}

func TestServerImplAddToDoListSuccessUseUserConfig(t *testing.T) {
	var (
		ctx    = context.TODO()
		userID = "user-id"
		req    = pb.AddToDoListRequest{
			UserId: userID,
			Entry: []*pb.ToDoEntry{
				{
					Content: "content",
				},
			},
		}
	)
	mockToDoStore := mockStore.ToDo{}
	mockToDoStore.Mock.
		On("GetConfig", ctx, userID).
		Return(&models.ToDoConfig{
			UserID:  userID,
			Limited: int64(2),
		}, nil)
	mockToDoStore.Mock.
		On("GetUsedCount", ctx, userID, someRandomDay).
		Return(int64(1), nil)
	mockToDoStore.Mock.
		On("IncreaseUsedCount", ctx, userID, someRandomDay, int64(len(req.GetEntry()))).
		Return(int64(2), nil)
	var recordList = mapper.Proto2ModelToDoEntryList(req.GetEntry())
	for _, record := range recordList {
		record.UserID = userID
	}
	mockToDoStore.Mock.
		On("Record", ctx, recordList).
		Return(nil)
	srv := serverImpl{
		cfg: &configs.Config{
			ToDoListAddLimitedPerDay: 1,
		},
		toDoStore: &mockToDoStore,
	}
	resp, err := srv.AddToDoList(ctx, &req)
	assert.NoError(t, err)
	assert.Equal(t, resp, &pb.AddToDoListResponse{Message: "ok"})
	mockToDoStore.AssertExpectations(t)
}

func TestServerImplAddToDoListSuccessUseDefaultConfig(t *testing.T) {
	var (
		ctx    = context.TODO()
		userID = "user-id"
		req    = pb.AddToDoListRequest{
			UserId: userID,
			Entry: []*pb.ToDoEntry{
				{
					Content: "content",
				},
			},
		}
	)
	mockToDoStore := mockStore.ToDo{}
	mockToDoStore.Mock.
		On("GetConfig", ctx, userID).
		Return(nil, store.ErrNotFound)
	mockToDoStore.Mock.
		On("GetUsedCount", ctx, userID, someRandomDay).
		Return(int64(1), nil)
	mockToDoStore.Mock.
		On("IncreaseUsedCount", ctx, userID, someRandomDay, int64(len(req.GetEntry()))).
		Return(int64(2), nil)
	var recordList = mapper.Proto2ModelToDoEntryList(req.GetEntry())
	for _, record := range recordList {
		record.UserID = userID
	}
	mockToDoStore.Mock.
		On("Record", ctx, recordList).
		Return(nil)
	srv := serverImpl{
		cfg: &configs.Config{
			ToDoListAddLimitedPerDay: 2,
		},
		toDoStore: &mockToDoStore,
	}
	resp, err := srv.AddToDoList(ctx, &req)
	assert.NoError(t, err)
	assert.Equal(t, resp, &pb.AddToDoListResponse{Message: "ok"})
	mockToDoStore.AssertExpectations(t)
}

func TestServerImplAddToDoListFailedGetUsedCount(t *testing.T) {
	var (
		ctx    = context.TODO()
		userID = "user-id"
		req    = pb.AddToDoListRequest{
			UserId: userID,
			Entry: []*pb.ToDoEntry{
				{
					Content: "content",
				},
			},
		}
	)
	mockToDoStore := mockStore.ToDo{}
	mockToDoStore.Mock.
		On("GetConfig", ctx, userID).
		Return(nil, store.ErrNotFound)
	mockToDoStore.Mock.
		On("GetUsedCount", ctx, userID, someRandomDay).
		Return(int64(0), someRandomError)
	srv := serverImpl{
		cfg: &configs.Config{
			ToDoListAddLimitedPerDay: 2,
		},
		toDoStore: &mockToDoStore,
	}
	resp, err := srv.AddToDoList(ctx, &req)
	assert.ErrorIs(t, err, errInternal)
	assert.Nil(t, resp)
	mockToDoStore.AssertExpectations(t)
}

func TestServerImplAddToDoListFailedIncreaseUsedCount(t *testing.T) {
	var (
		ctx    = context.TODO()
		userID = "user-id"
		req    = pb.AddToDoListRequest{
			UserId: userID,
			Entry: []*pb.ToDoEntry{
				{
					Content: "content",
				},
			},
		}
	)
	mockToDoStore := mockStore.ToDo{}
	mockToDoStore.Mock.
		On("GetConfig", ctx, userID).
		Return(nil, store.ErrNotFound)
	mockToDoStore.Mock.
		On("GetUsedCount", ctx, userID, someRandomDay).
		Return(int64(1), nil)
	mockToDoStore.Mock.
		On("IncreaseUsedCount", ctx, userID, someRandomDay, int64(len(req.GetEntry()))).
		Return(int64(0), someRandomError)
	srv := serverImpl{
		cfg: &configs.Config{
			ToDoListAddLimitedPerDay: 2,
		},
		toDoStore: &mockToDoStore,
	}
	resp, err := srv.AddToDoList(ctx, &req)
	assert.ErrorIs(t, err, errInternal)
	assert.Nil(t, resp)
	mockToDoStore.AssertExpectations(t)
}

func TestServerImplAddToDoListFailedRecord(t *testing.T) {
	var (
		ctx    = context.TODO()
		userID = "user-id"
		req    = pb.AddToDoListRequest{
			UserId: userID,
			Entry: []*pb.ToDoEntry{
				{
					Content: "content",
				},
			},
		}
	)
	mockToDoStore := mockStore.ToDo{}
	mockToDoStore.Mock.
		On("GetConfig", ctx, userID).
		Return(nil, store.ErrNotFound)
	mockToDoStore.Mock.
		On("GetUsedCount", ctx, userID, someRandomDay).
		Return(int64(1), nil)
	mockToDoStore.Mock.
		On("IncreaseUsedCount", ctx, userID, someRandomDay, int64(len(req.GetEntry()))).
		Return(int64(2), nil)
	var recordList = mapper.Proto2ModelToDoEntryList(req.GetEntry())
	for _, record := range recordList {
		record.UserID = userID
	}
	mockToDoStore.Mock.
		On("Record", ctx, recordList).
		Return(someRandomError)
	mockToDoStore.Mock.
		On("DecreaseUsedCount", ctx, userID, someRandomDay, int64(len(req.GetEntry()))).
		Return(int64(1), nil)
	srv := serverImpl{
		cfg: &configs.Config{
			ToDoListAddLimitedPerDay: 2,
		},
		toDoStore: &mockToDoStore,
	}
	resp, err := srv.AddToDoList(ctx, &req)
	assert.ErrorIs(t, err, errInternal)
	assert.Nil(t, resp)
	mockToDoStore.AssertExpectations(t)
}

func TestServerImplAddToDoListFailedReachedLimit(t *testing.T) {
	var (
		ctx    = context.TODO()
		userID = "user-id"
		req    = pb.AddToDoListRequest{
			UserId: userID,
			Entry: []*pb.ToDoEntry{
				{
					Content: "content",
				},
			},
		}
	)
	mockToDoStore := mockStore.ToDo{}
	mockToDoStore.Mock.
		On("GetConfig", ctx, userID).
		Return(nil, store.ErrNotFound)
	mockToDoStore.Mock.
		On("GetUsedCount", ctx, userID, someRandomDay).
		Return(int64(2), nil)
	srv := serverImpl{
		cfg: &configs.Config{
			ToDoListAddLimitedPerDay: 2,
		},
		toDoStore: &mockToDoStore,
	}
	resp, err := srv.AddToDoList(ctx, &req)
	assert.ErrorIs(t, err, errDailyQuotaReached)
	assert.Nil(t, resp)
	mockToDoStore.AssertExpectations(t)
}

func TestServerImplAddToDoListFailedExceedLimit(t *testing.T) {
	var (
		ctx    = context.TODO()
		userID = "user-id"
		req    = pb.AddToDoListRequest{
			UserId: userID,
			Entry: []*pb.ToDoEntry{
				{
					Content: "content",
				},
				{
					Content: "content",
				},
			},
		}
	)
	mockToDoStore := mockStore.ToDo{}
	mockToDoStore.Mock.
		On("GetConfig", ctx, userID).
		Return(nil, store.ErrNotFound)
	mockToDoStore.Mock.
		On("GetUsedCount", ctx, userID, someRandomDay).
		Return(int64(1), nil)
	srv := serverImpl{
		cfg: &configs.Config{
			ToDoListAddLimitedPerDay: 2,
		},
		toDoStore: &mockToDoStore,
	}
	resp, err := srv.AddToDoList(ctx, &req)
	assert.ErrorIs(t, err, errDailyQuotaExceed)
	assert.Nil(t, resp)
	mockToDoStore.AssertExpectations(t)
}

func TestServerImplAddToDoListFailedExceedLimit2(t *testing.T) {
	var (
		ctx    = context.TODO()
		userID = "user-id"
		req    = pb.AddToDoListRequest{
			UserId: userID,
			Entry: []*pb.ToDoEntry{
				{
					Content: "content",
				},
				{
					Content: "content",
				},
			},
		}
	)
	mockToDoStore := mockStore.ToDo{}
	mockToDoStore.Mock.
		On("GetConfig", ctx, userID).
		Return(nil, store.ErrNotFound)
	mockToDoStore.Mock.
		On("GetUsedCount", ctx, userID, someRandomDay).
		Return(int64(0), nil)
	mockToDoStore.Mock.
		On("IncreaseUsedCount", ctx, userID, someRandomDay, int64(len(req.GetEntry()))).
		Return(int64(3), nil)
	mockToDoStore.Mock.
		On("DecreaseUsedCount", ctx, userID, someRandomDay, int64(len(req.GetEntry()))).
		Return(int64(1), nil)
	srv := serverImpl{
		cfg: &configs.Config{
			ToDoListAddLimitedPerDay: 2,
		},
		toDoStore: &mockToDoStore,
	}
	resp, err := srv.AddToDoList(ctx, &req)
	assert.ErrorIs(t, err, errDailyQuotaExceed)
	assert.Nil(t, resp)
	mockToDoStore.AssertExpectations(t)
}

package usecase

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"togo.com/pkg/model"
	"togo.com/pkg/repository/mock"
)

func Test_taskUseCase_AddTask(t *testing.T) {
	now := time.Now()
	createDate := now.Format("2006-01-02")
	var cases = []struct {
		name                string
		input               model.AddTaskParams
		mockGetLimitPerUser int64
		mockCountTaskPerDay int64
		output              model.AddTaskResponse
		outputError         error
	}{
		{name: "add Task success", input: model.AddTaskParams{
			UserId:     "1",
			CreateDate: createDate,
			Content:    "content",
		},
			mockGetLimitPerUser: int64(5),
			mockCountTaskPerDay: int64(1),
			output: model.AddTaskResponse{
				UserId:     "1",
				Content:    "content",
				CreateDate: createDate,
			},
			outputError: nil,
		},
		{name: "add Task fail by  limit task per day", input: model.AddTaskParams{
			UserId:     "1",
			CreateDate: createDate,
			Content:    "content",
		},
			mockGetLimitPerUser: int64(5),
			mockCountTaskPerDay: int64(5),
			output: model.AddTaskResponse{
				UserId:     "",
				Content:    "",
				CreateDate: "",
			},
			outputError: errors.New("Limit reached for the day "),
		},
	}
	ctrl := gomock.NewController(t)
	mockRepo := mock.NewMockRepository(ctrl)
	defer ctrl.Finish()
	for _, s := range cases {
		t.Run(s.name, func(t *testing.T) {
			if s.mockCountTaskPerDay < s.mockGetLimitPerUser {
				mockRepo.EXPECT().AddTask(context.TODO(), model.AddTaskParams{
					UserId:     s.input.UserId,
					CreateDate: s.input.CreateDate,
					Content:    s.input.Content,
				}).Times(1).Return(nil)
			}
			mockRepo.EXPECT().GetLimitPerUser(context.Background(), s.input.UserId).Times(1).Return(s.mockGetLimitPerUser, nil)
			mockRepo.EXPECT().CountTaskPerDay(context.Background(), s.input.UserId, s.input.CreateDate).Times(1).Return(s.mockCountTaskPerDay, nil)
			service := taskUseCase{repo: mockRepo}
			resp, err := service.AddTask(context.TODO(), s.input.UserId, model.AddTaskRequest{Content: s.input.Content})
			assert.Equal(t, s.outputError, err)
			assert.Equal(t, s.output, resp)
		})
	}

}

func Test_taskUseCase_GetListTaskByDate(t *testing.T) {
	tests := []struct {
		name              string
		userId            string
		createDate        string
		mockRetrieveTasks []model.Task
		output            []model.Task
	}{
		// TODO: Add test cases.
		{name: "get list task by user ", userId: "1", createDate: "2022-03-15", output: []model.Task{
			{ID: "1", Content: "content", UserID: "user", CreatedDate: "2022-03-14"},
		},
			mockRetrieveTasks: []model.Task{
				{ID: "1", Content: "content", UserID: "user", CreatedDate: "2022-03-14"},
			},
		},
		{name: "get list task by user  fail", userId: "1", createDate: "2022-03-15", output: []model.Task{
			{ID: "", Content: "", UserID: "", CreatedDate: ""},
		},
			mockRetrieveTasks: []model.Task{
				{ID: "", Content: "", UserID: "", CreatedDate: ""},
			},
		},
	}
	ctrl := gomock.NewController(t)
	mockRepo := mock.NewMockRepository(ctrl)
	defer ctrl.Finish()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.EXPECT().RetrieveTasks(context.Background(), tt.userId, tt.createDate).Times(1).Return(tt.output, nil)
			service := taskUseCase{repo: mockRepo}
			resp, _ := service.GetListTaskByDate(context.TODO(), tt.userId, tt.createDate)
			assert.Equal(t, tt.output, resp)
		})
	}
}

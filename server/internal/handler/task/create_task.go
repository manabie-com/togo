package task

import (
	"context"
	"github.com/HoangVyDuong/togo/internal/handler"
	taskEntity "github.com/HoangVyDuong/togo/internal/storages/task"
	"github.com/HoangVyDuong/togo/pkg/define"
	taskDTO "github.com/HoangVyDuong/togo/pkg/dtos/task"
	"github.com/HoangVyDuong/togo/pkg/logger"
	"github.com/HoangVyDuong/togo/pkg/utils"
	"strconv"
)

func (th *taskHandler) CreateTask(ctx context.Context, request taskDTO.CreateTaskRequest) (response taskDTO.CreateTaskResponse, err error) {
	logger.Debug("[Handler][CreateTask] Start Create Task")

	if request.Content == "" {
		logger.Debug("[Handler][CreateTask] Invalid Param")
		return taskDTO.CreateTaskResponse{}, define.FailedValidation
	}

	userID, err := handler.UserIDFromContext(ctx)
	if err != nil {
		return taskDTO.CreateTaskResponse{}, err
	}

	taskID := utils.GenID()

	err = th.taskService.CreateTask(ctx, taskEntity.Task{
		ID: taskID,
		Content: request.Content,
		UserID: userID,
	})
	if err != nil {
		return taskDTO.CreateTaskResponse{}, define.Unknown
	}

	logger.Debug("[Handler][CreateTask] Create Task successfully")
	return taskDTO.CreateTaskResponse{
		TaskID: strconv.FormatUint(taskID, 10),
	}, nil
}

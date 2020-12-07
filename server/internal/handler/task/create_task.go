package task

import (
	"context"
	taskDTO "github.com/HoangVyDuong/togo/pkg/dtos/task"
)

func (th *taskHandler) CreateTask(ctx context.Context, request taskDTO.CreateTaskRequest) (response taskDTO.CreateTaskResponse, err error) {
	return taskDTO.CreateTaskResponse{}, nil
}

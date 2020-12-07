package task

import (
	"context"
	"github.com/HoangVyDuong/togo/pkg/dtos"
	taskDTO "github.com/HoangVyDuong/togo/pkg/dtos/task"
)


func (th *taskHandler) GetTasks(ctx context.Context, request dtos.EmptyRequest) (response taskDTO.Tasks, err error) {
	return taskDTO.Tasks{}, nil
}


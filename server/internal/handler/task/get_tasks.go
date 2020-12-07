package task

import (
	"context"
	"github.com/HoangVyDuong/togo/internal/handler"
	"github.com/HoangVyDuong/togo/pkg/define"
	"github.com/HoangVyDuong/togo/pkg/dtos"
	taskDTO "github.com/HoangVyDuong/togo/pkg/dtos/task"
	"github.com/HoangVyDuong/togo/pkg/logger"
	"strconv"
)

func (th *taskHandler) GetTasks(ctx context.Context, _ dtos.EmptyRequest) (response taskDTO.Tasks, err error) {
	logger.Debugf("Start GetTasks")

	userID, err := handler.UserIDFromContext(ctx)

	if err != nil {
		return taskDTO.Tasks{}, err
	}

	tasks, err := th.taskService.GetTasks(ctx, userID)
	if err != nil {
		return taskDTO.Tasks{}, define.Unknown
	}

	var taskDTOs = make([]taskDTO.Task, len(tasks))
	for idx, task := range tasks {
		taskDTOs[idx] = taskDTO.Task{
			Id: strconv.FormatUint(task.ID, 10),
			Content: task.Content,
		}
	}

	logger.Debugf("GetTasks With UserID: %d", userID)
	return taskDTO.Tasks{
		Data: taskDTOs,
	}, nil
}


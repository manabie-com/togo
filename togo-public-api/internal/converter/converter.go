package converter

import (
	"togo-public-api/internal/service/togo_internal_v1"
	v1 "togo-public-api/pkg/api/v1"
)

func ToTask(task *togo_internal_v1.Task) *v1.Task {
	return &v1.Task{
		Id:          task.GetId(),
		UserId:      task.GetUserId(),
		Title:       task.GetTitle(),
		Content:     task.GetContent(),
		CreatedTime: task.GetCreatedTime(),
	}
}

func ToTasks(tasks []*togo_internal_v1.Task) []*v1.Task {
	if tasks == nil {
		return nil
	}
	result := make([]*v1.Task, 0, len(tasks))
	for _, task := range tasks {
		result = append(result, ToTask(task))
	}
	return result
}

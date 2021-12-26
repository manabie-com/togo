package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"togo-internal-service/internal/model"
	v1 "togo-internal-service/pkg/api/v1"
)

func TaskToDTO(task *model.Task) *v1.Task {
	return &v1.Task{
		Id:          task.ID,
		UserId:      task.UserID,
		Content:     task.Content,
		Title:       task.Title,
		CreatedTime: timestamppb.New(task.CreatedTime),
	}
}

func TasksToDTO(tasks []*model.Task) []*v1.Task {
	result := make([]*v1.Task, 0, len(tasks))
	for _, task := range tasks {
		result = append(result, TaskToDTO(task))
	}
	return result
}

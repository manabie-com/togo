package formatter

import (
	"context"

	"github.com/valonekowd/togo/domain/entity"
	"github.com/valonekowd/togo/usecase/interfaces"
	"github.com/valonekowd/togo/usecase/response"
)

type taskFormatter struct{}

var _ interfaces.TaskPresenter = taskFormatter{}

func NewTaskFormatter() interfaces.TaskPresenter {
	return taskFormatter{}
}

func (f taskFormatter) mappingTaskData(ctx context.Context, t *entity.Task) *response.TaskData {
	return &response.TaskData{
		ID:          t.ID,
		Content:     t.Content,
		UserID:      t.UserID,
		CreatedDate: t.CreatedAt.Format("2006-01-02"),
	}
}

func (f taskFormatter) Fetch(ctx context.Context, ts []*entity.Task) *response.GetTasks {
	data := make([]*response.TaskData, 0, len(ts))

	for _, t := range ts {
		data = append(data, f.mappingTaskData(ctx, t))
	}

	return &response.GetTasks{Data: data}
}

func (f taskFormatter) Create(ctx context.Context, t *entity.Task) *response.CreateTask {
	return &response.CreateTask{Data: f.mappingTaskData(ctx, t)}
}

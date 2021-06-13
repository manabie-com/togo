package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/model"
	"github.com/manabie-com/togo/internal/storages/ent"
	"github.com/manabie-com/togo/internal/storages/ent/task"
	"github.com/manabie-com/togo/internal/storages/ent/user"
	"time"
)

type ToDoService interface {
	CreateTask(ctx context.Context, req model.TaskCreationRequest) (*model.Task, error)

	GetTaskByDate(ctx context.Context, queryDate time.Time) (*[]model.Task, error)
}

type toDoServiceImpl struct {
	client *ent.Client
}

func NewToDoService(client *ent.Client) ToDoService {
	return &toDoServiceImpl{client: client}
}

func (t *toDoServiceImpl) CreateTask(ctx context.Context, req model.TaskCreationRequest) (*model.Task, error) {
	userId := ctx.Value("userId").(string)

	owner, err := t.client.User.Query().Where(user.UserIDEQ(userId)).Only(ctx)
	if err != nil {
		return nil, err
	}

	foundTask, err := t.client.Task.Create().
		SetTaskID(uuid.NewString()).
		SetContent(req.Content).SetOwner(owner).Save(ctx)
	if err != nil {
		return nil, err
	}

	return &model.Task{
		TaskID:      foundTask.TaskID,
		Content:     foundTask.Content,
		CreatedDate: foundTask.CreatedDate,
		UserID:      userId,
	}, nil
}

func (t *toDoServiceImpl) GetTaskByDate(ctx context.Context, queryDate time.Time) (*[]model.Task, error) {
	userId := ctx.Value("userId").(string)
	nextDate := queryDate.Add(time.Hour * 24)

	tasks, err := t.client.Task.Query().Where(task.HasOwnerWith(user.UserIDEQ(userId)), task.CreatedDateGTE(queryDate), task.CreatedDateLT(nextDate)).All(ctx)
	if err != nil {
		return nil, err
	}

	allTask := []model.Task{}

	for _, e := range tasks {
		elems := model.Task{
			TaskID:      e.TaskID,
			Content:     e.Content,
			CreatedDate: e.CreatedDate,
			UserID:      userId,
		}
		allTask = append(allTask, elems)

	}

	return &allTask, nil
}

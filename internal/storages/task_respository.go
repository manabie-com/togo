package storages

import (
	"context"
	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/storages/ent"
	"github.com/manabie-com/togo/internal/storages/ent/task"
	"github.com/manabie-com/togo/internal/storages/ent/user"
	"time"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, content string, owner *ent.User) (*ent.Task, error)

	GetTaskByDate(ctx context.Context, userId string, gte time.Time, lt time.Time) ([]*ent.Task, error)
}

type taskRepositoryImpl struct {
	client *ent.Client
}

func NewTaskRepository(client *ent.Client) TaskRepository {
	return &taskRepositoryImpl{client: client}
}

func (t *taskRepositoryImpl) CreateTask(ctx context.Context, content string, owner *ent.User) (*ent.Task, error) {
	return t.client.Task.Create().
		SetTaskID(uuid.NewString()).
		SetContent(content).SetOwner(owner).Save(ctx)
}

func (t *taskRepositoryImpl) GetTaskByDate(ctx context.Context, userId string, gte time.Time, lt time.Time) ([]*ent.Task, error) {
	return t.client.Task.Query().Where(task.HasOwnerWith(user.UserIDEQ(userId)), task.CreatedDateGTE(gte), task.CreatedDateLT(lt)).All(ctx)
}

package postgresql

import (
	"context"
	"database/sql"
	"github.com/manabie-com/togo/internal/model"
	"upper.io/db.v3/lib/sqlbuilder"
)

const tasksTable string = "tasks"

type TaskStore struct {
	db sqlbuilder.Database
}

func NewTaskStore(db sqlbuilder.Database) TaskStore {
	return TaskStore{
		db: db,
	}
}

func (s TaskStore) RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*model.Task, error) {
	var ts []*model.Task
	err := s.db.WithContext(ctx).SelectFrom(tasksTable).Where("user_id = ? and createdDate = ?", userID, createdDate).All(&ts)
	if err != nil {
		return nil, model.NewError(model.ErrListTasks, err.Error())
	}

	return ts, err
}

func (s TaskStore) AddTask(ctx context.Context, userID string, t *model.Task) (*model.Task, error) {
	t.UserID = userID
	if err := t.IsValid(); err != nil {
		return nil, err
	}

	i := s.db.WithContext(ctx).InsertInto(tasksTable).Values(t).Returning("created_date").Iterator()
	defer i.Close()

	err := i.NextScan(&t.CreatedDate)
	if err != nil {
		return nil, model.NewError(model.ErrAddTasks, err.Error())
	}

	return t, err
}

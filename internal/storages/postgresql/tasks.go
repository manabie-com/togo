package postgresql

import (
	"context"
	"database/sql"
	"github.com/manabie-com/togo/internal/model"
	"time"
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

func (s TaskStore) RetrieveTasks(ctx context.Context, userID string, createdDate sql.NullString) ([]*model.Task, error) {
	var ts []*model.Task
	var err error
	err = s.db.WithContext(ctx).SelectFrom(tasksTable).Where(`user_id = ?`, userID).All(&ts)
	if createdDate.Valid && len(createdDate.String) != 0 {
		err = s.db.WithContext(ctx).SelectFrom(tasksTable).Where(`user_id = ? and created_date = ?`, userID, createdDate).All(&ts)
	}
	if err != nil {
		return nil, model.NewError(model.ErrListTasks, err.Error())
	}

	return ts, err
}

func (s TaskStore) AddTask(ctx context.Context, userID string, t *model.Task) (*model.Task, error) {
	t.UserID = userID
	t.CreatedDate = time.Now().UTC().Format("2006-01-02")
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

func (s TaskStore) CountTasksByUser(ctx context.Context, userID string) (int, error) {
	rows, err := s.db.Query(
		`select count(*) from tasks where tasks.user_id = ?`,
		userID,
	)
	if err != nil {
		return 0, model.NewError(model.ErrCountTasks, err.Error())
	}
	defer rows.Close()

	res := 0
	err = sqlbuilder.NewIterator(rows).ScanOne(&res)
	if err != nil {
		return 0, model.NewError(model.ErrCountTasks, err.Error())
	}

	return res, nil
}

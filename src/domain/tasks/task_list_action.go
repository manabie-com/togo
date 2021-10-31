package tasks

import (
	"github.com/go-pg/pg/v10"
	"github.com/pkg/errors"
)

type TaskListAction struct {
	Db *pg.DB
}

func (T TaskListAction) Execute(UserID int64) (tasks Tasks, err error) {
	err = T.Db.Model(&tasks).Where("user_id = ?", UserID).Select()
	if err != nil {
		return Tasks{}, errors.Wrapf(err, "List task error")
	}

	return tasks, nil
}

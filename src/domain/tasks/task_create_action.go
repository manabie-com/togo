package tasks

import (
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/pkg/errors"
	"github.com/quochungphp/go-test-assignment/src/pkgs/token"
)

type TaskCreateAction struct {
	Db *pg.DB
}

func (T TaskCreateAction) Execute(Content string) (task Task, err error) {
	t := time.Now()
	currentDate := t.Format("2006-01-02")
	sessionUser := token.AccessUser
	task = Task{Content: Content, UserID: sessionUser.UserID}

	count, err := T.Db.Model(new(Task)).Where("user_id = ?", sessionUser.UserID).Where("created_date::timestamp::date = ?", currentDate).Count()
	if err != nil {
		return Task{}, errors.Wrapf(err, "Create a task error")
	}

	if count == sessionUser.MaxTodo {
		return Task{}, errors.Errorf("Task over %d tasks, can not create", sessionUser.MaxTodo)
	}

	_, err = T.Db.Model(&task).Insert()
	if err != nil {
		return Task{}, errors.Wrapf(err, "Create a task error")
	}

	return task, nil
}

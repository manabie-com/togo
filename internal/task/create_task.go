package task

import (
	"errors"

	"github.com/jmramos02/akaru/internal/model"
)

func (t task) CreateTask(name string) (task model.Task, err error) {
	task = model.Task{
		Name:   name,
		UserID: t.userID,
	}

	if err = t.db.Save(&task).Error; err != nil {
		return task, errors.New("error saving task")
	}

	return task, nil
}

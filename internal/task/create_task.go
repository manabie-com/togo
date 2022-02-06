package task

import "github.com/jmramos02/akaru/internal/model"

func (t task) CreateTask(name string) {
	task := model.Task{
		Name:   name,
		UserID: t.userID,
	}

	t.db.Save(&task)
}

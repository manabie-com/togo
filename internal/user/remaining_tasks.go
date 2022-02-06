package user

import "github.com/jmramos02/akaru/internal/model"

func (u user) getRemainingTasksForTheDay(id int) int {
	var numberOfTasks int64
	err := u.db.Model(model.Task{}).Where("user_id = ? AND created_at >= DATE('now') AND created_at < DATE('now', '+1 day')", id).Count(&numberOfTasks).Error

	if err != nil {
		panic(err)
	}

	return int(numberOfTasks)
}

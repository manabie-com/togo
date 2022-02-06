package user

import (
	"github.com/jmramos02/akaru/internal/model"
)

func (u user) CanUserInsert() bool {
	//get the user first
	var userRecord model.User
	err := u.db.Where("username= ?", u.username).First(&userRecord).Error

	if err != nil {
		panic(err)
	}

	//get the user's remaining task
	numberOfTasks := u.getRemainingTasksForTheDay(int(userRecord.ID))

	//compare the numbers
	if numberOfTasks > userRecord.Limit {
		return false
	}

	return true
}

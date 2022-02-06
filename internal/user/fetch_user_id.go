package user

import (
	"github.com/jmramos02/akaru/internal/model"
)

func (u user) GetUserID() int {
	var userRecord model.User
	err := u.db.Where("username= ?", u.username).First(&userRecord).Error

	if err != nil {
		panic(err)
	}

	return int(userRecord.ID)
}

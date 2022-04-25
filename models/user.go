package models

import (
	"time"
)

type User struct {
	LimitTasks int    `json:"limitTasks"`
	Id         int    `json:"userid" gorm:"primary_key;auto_increment"`
	Tasks      []Todo `gorm:"foreignKey:Userid;references:Id; polymorphic:Owner"`
}

func (b *User) CountTasks() int {
	var countTask = 0

	for _, v := range b.Tasks {
		if v.Date.Truncate(24 * time.Hour).Equal(time.Now().Truncate(24 * time.Hour)) {
			countTask++
		}
	}

	return countTask
}

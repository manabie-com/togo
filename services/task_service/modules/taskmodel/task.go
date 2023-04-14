package taskmodel

import (
	"time"

	"github.com/phathdt/libs/go-sdk/sdkcm"
)

type Task struct {
	sdkcm.SQLModel `json:",inline"`
	Content        string    `json:"content" gorm:"column:content;"`
	UserId         int       `json:"user_id" gorm:"column:user_id"`
	CreatedDate    time.Time `json:"created_date" gorm:"column:created_date"`
	IsDone         bool      `json:"is_done" gorm:"column:is_done"`
}

func (Task) TableName() string {
	return "tasks"
}

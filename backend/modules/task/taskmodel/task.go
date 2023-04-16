package taskmodel

import (
	"github.com/golang-module/carbon/v2"
	"github.com/phathdt/libs/go-sdk/sdkcm"
)

type Task struct {
	sdkcm.SQLModel `json:",inline"`
	Content        string      `json:"content" gorm:"column:content;"`
	UserId         int         `json:"user_id" gorm:"column:user_id"`
	CreatedDate    carbon.Date `json:"created_date" gorm:"column:created_date"`
	IsDone         bool        `json:"is_done" gorm:"column:is_done"`
}

func (Task) TableName() string {
	return "tasks"
}

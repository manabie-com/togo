package model

import (
	"errors"
	"github.com/manabie-com/togo/model"
	apperror "github.com/manabie-com/togo/shared/app_error"
)

const (
	EntityName = "Task"
)

var (
	ErrYouAreLimitedReach = apperror.NewCustomError(errors.New("you are limited reach"),
		"you are limited reach", "ErrYouAreLimitedReach")
)

type Task struct {
	model.BaseModel
	UserId  int    `gorm:"column:user_id"`
	Content string `gorm:"column:content"`
}

func (Task) TableName() string {
	return "tasks"
}

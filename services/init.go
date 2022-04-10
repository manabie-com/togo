package services

import (
	"github.com/namnhatdoan/togo/models"
	"github.com/namnhatdoan/togo/settings"
	"time"
)

type ToGoServiceI interface {
	AddNewTask(email, task string) (*models.Tasks, error)
	SetUserConfig(email string, limit int8, date time.Time) (*models.UserConfigs, error)
}

var log = settings.GetLogger()

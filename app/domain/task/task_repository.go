package task

import (
	"github.com/ansidev/togo/domain/user"
	"time"
)

type ITaskRepository interface {
	Create(taskModel Task, userModel user.User) (Task, error)
	GetTotalTasksByUserAndDate(userModel user.User, date time.Time) (int64, error)
}

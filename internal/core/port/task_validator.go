package port

import (
	"github.com/manabie-com/togo/internal/core/domain"
)

type TaskValidator interface {
	ValidateBeforeRetrieveTasks(userId, createdDate string) error
	ValidateBeforeAddTask(task *domain.Task) error
	ValidateBeforeLogin(username, password string) error
}

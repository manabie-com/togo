package validator

import (
	"errors"

	"github.com/manabie-com/togo/internal/core/domain"
	"github.com/manabie-com/togo/internal/core/port"
)

func NewTaskValidator() port.TaskValidator {
	return new(taskValidator)
}

type taskValidator struct {
}

func (p *taskValidator) ValidateBeforeRetrieveTasks(userId, createdDate string) error {
	lenUserId := len(userId)
	if lenUserId == 0 || lenUserId != 36 {
		return errors.New("invalid user-id")
	}
	if len(createdDate) == 0 {
		return errors.New("missing query created_date")
	}
	return nil
}

func (p *taskValidator) ValidateBeforeAddTask(task *domain.Task) error {
	if len(task.Content) == 0 {
		return errors.New("task's content cannot be empty")
	}
	return nil
}

func (p *taskValidator) ValidateBeforeLogin(username, password string) error {
	if len(username) == 0 {
		return errors.New("invalid username")
	}
	if len(password) == 0 {
		return errors.New("invalid password")
	}
	return nil
}

package task

import (
	"github.com/manabie-com/backend/entity"
	"github.com/manabie-com/backend/repository"
	"github.com/manabie-com/backend/utils"
)

type I_TaskServiceValidate interface {
	Validate(task *entity.Task) *utils.ErrorRest
}

type taskServiceValidate struct{}

var (
	repo repository.I_Repository
)

func NewTaskServiceValidate(repository repository.I_Repository) I_TaskServiceValidate {
	repo = repository
	return &taskServiceValidate{}
}

func (*taskServiceValidate) Validate(task *entity.Task) *utils.ErrorRest {
	if task == nil {

		return utils.ErrBadRequest("The task cannot empty")
	}

	if task.Content == "" {

		return utils.ErrBadRequest("The task content cannot empty")

	}

	if string(task.UserID) == "" {

		return utils.ErrBadRequest("The userId cannot empty")

	}

	result, err := repo.FindTaskByContent(task.Content)
	if result != nil {
		return utils.ErrBadRequest("The task is already existed")
	}
	if err != nil {
		return err
	}

	return nil
}

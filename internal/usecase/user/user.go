package user

import (
	"errors"

	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/storages/task"
	"github.com/manabie-com/togo/internal/storages/user"
)

type userUsecase struct {
	userStorage user.UserStorageInterface
	taskStorage task.TaskStorageInterface
}

func NewUserUsecase(userStorage user.UserStorageInterface, taskStorage task.TaskStorageInterface) UserUsecaseInterface {
	return &userUsecase{
		userStorage: userStorage,
		taskStorage: taskStorage,
	}
}

func (u *userUsecase) ValidateUser(id, password string) error {
	return u.userStorage.GetUser(id, password)
}

func (u *userUsecase) CreateTask(task *storages.Task) error {
	user, err := u.userStorage.GetUsersTasks(task.UserID, task.CreatedDate)
	if err != nil {
		return err
	}

	if len(user.Tasks) >= user.MaxTodo {
		return errors.New("over limit per day")
	}

	if err = u.taskStorage.CreateTask(task); err != nil {
		return err
	}

	return nil
}

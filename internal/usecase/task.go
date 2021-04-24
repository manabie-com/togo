package usecase

import "github.com/manabie-com/togo/internal/domain"

type taskUseCase struct {
	taskStore domain.TaskStore
	userStore domain.UserStore
}

func NewTaskUseCase(taskStore domain.TaskStore, userStore domain.UserStore) taskUseCase {
	return taskUseCase{
		taskStore: taskStore,
		userStore: userStore,
	}
}

func (u taskUseCase) AddTask(task domain.Task) error {
	limit, err := u.userStore.GetUserTasksPerDay(task.UserID)
	if err != nil {
		return err
	}
	return u.taskStore.AddTaskWithLimitPerDay(task, limit)
}

func (u taskUseCase) GetTasksByUserID(userID string) ([]domain.Task, error) {
	return u.taskStore.GetTasksByUserID(userID)
}

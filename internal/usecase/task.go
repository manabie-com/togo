package usecase

import (
	e "lntvan166/togo/internal/entities"
	repo "lntvan166/togo/internal/repository"
)

func AddTask(t *e.Task) error {
	return repo.Repository.CreateTask(t)
}

func GetNumberOfTaskTodayByUserID(id int) (int, error) {
	return repo.Repository.GetNumberOfTaskTodayByUserID(id)
}

func GetAllTask() (*[]e.Task, error) {
	return repo.Repository.GetAllTask()
}

func GetTaskByID(id int) (*e.Task, error) {
	return repo.Repository.GetTaskByID(id)
}

func GetTasksByUsername(username string) (*[]e.Task, error) {
	userID, err := GetUserIDByUsername(username)
	if err != nil {
		return nil, err
	}
	return repo.Repository.GetTasksByUserID(userID)
}

func CheckLimitTaskToday(id int) (bool, error) {
	maxTask, err := GetMaxTaskByUserID(id)
	if err != nil {
		return false, err
	}
	numberTask, err := repo.Repository.GetNumberOfTaskTodayByUserID(id)
	if err != nil {
		return false, err
	}
	if numberTask >= maxTask {
		return true, nil
	}
	return false, nil
}

func UpdateTask(t *e.Task) error {
	return repo.Repository.UpdateTask(t)
}

func DeleteTask(id int) error {
	return repo.Repository.DeleteTask(id)
}

package usecase

import (
	"errors"
	e "lntvan166/togo/internal/entities"
	repo "lntvan166/togo/internal/repository"
	"net/http"
)

func AddUser(u *e.User) error {
	return repo.Repository.AddUser(u)
}

func GetUserByName(username string) (*e.User, error) {
	u, err := repo.Repository.GetUserByName(username)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func GetUserIDByUsername(username string) (int, error) {
	u, err := repo.Repository.GetUserByName(username)
	if err != nil {
		return 0, err
	}
	return u.ID, nil
}

func GetUserIDByTaskID(id int) (int, error) {
	t, err := repo.Repository.GetTaskByID(id)
	if err != nil {
		return 0, err
	}
	return t.UserID, nil
}

func GetMaxTaskByUserID(id int) (int, error) {
	u, err := repo.Repository.GetUserByID(id)
	if err != nil {
		return 0, err
	}
	return int(u.MaxTodo), nil
}
func UpdateUser(u *e.User) error {
	return repo.Repository.UpdateUser(u)
}
func CompleteTask(id int) error {
	return repo.Repository.CompleteTask(id)
}

func CheckUserExist(username string) (bool, error) {
	u, err := repo.Repository.GetUserByName(username)
	if err != nil {
		return false, err
	}
	if u.ID == 0 {
		return false, nil
	}
	return true, nil
}

func CheckAccessPermission(w http.ResponseWriter, username string, taskUserID int) error {
	userID, err := GetUserIDByUsername(username)
	if err != nil {
		return err
	}

	if userID != taskUserID {
		return errors.New("you are not allowed to access this task")
	}

	return nil
}

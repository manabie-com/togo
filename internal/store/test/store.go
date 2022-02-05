package test

import (
	"github.com/kozloz/togo"
)

type Store struct {
}

func (s *Store) GetUser(userID int64) (*togo.User, error) {
	return &togo.User{
		ID:         1,
		DailyLimit: 5,
	}, nil
}

func (s *Store) CreateTask(userID int64, task string) (*togo.Task, error) {
	return &togo.Task{
		ID:     1,
		UserID: userID,
		Name:   "TEST",
	}, nil
}

func (s *Store) CreateUser() (*togo.User, error) {
	return &togo.User{
		ID:         1,
		DailyLimit: 5,
	}, nil
}

func (s *Store) UpdateUser(user *togo.User) (*togo.User, error) {
	return user, nil
}

package test

import (
	"github.com/kozloz/togo"
)

type Store struct {
}

func (s *Store) Get(userID int64) (*togo.User, error) {
	return &togo.User{
		ID:         1,
		Username:   "test",
		DailyLimit: 5,
		DailyCount: 3,
	}, nil
}

func (s *Store) Create(userID int64, task string) (*togo.Task, error) {
	return &togo.Task{
		ID:     1,
		UserID: userID,
		Name:   "TEST",
	}, nil
}

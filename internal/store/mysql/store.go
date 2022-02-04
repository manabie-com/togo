package mysql

import "github.com/kozloz/togo"

type Store struct {
}

func (s *Store) Get(userID int64) (*togo.User, error) {
	return nil, nil
}

func (s *Store) Create(userID int64, task string) (*togo.Task, error) {
	return nil, nil
}

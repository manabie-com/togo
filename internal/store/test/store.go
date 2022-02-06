package test

import (
	"log"
	"time"

	"github.com/kozloz/togo"
)

// Store is a test implementation of the app's storage functions. Primarily used for unit tests.
type Store struct {
}

var user1 *togo.User = &togo.User{
	ID:         1,
	DailyLimit: 2,
}

var user2 *togo.User = &togo.User{
	ID:         2,
	DailyLimit: 2,
	DailyCounter: &togo.DailyCounter{
		DailyCount:  3,
		LastUpdated: time.Now().Add(-24 * time.Hour),
	},
}

func (s *Store) GetUser(userID int64) (*togo.User, error) {
	if userID == 1 {
		return user1, nil
	}
	if userID == 2 {
		return user2, nil
	}

	return nil, nil
}

func (s *Store) CreateTask(userID int64, task string) (*togo.Task, error) {
	log.Println(userID, task)
	return &togo.Task{
		ID:     1,
		UserID: userID,
		Name:   task,
	}, nil
}

func (s *Store) CreateUser(userID int64) (*togo.User, error) {
	return &togo.User{
		ID:         userID,
		DailyLimit: 5,
	}, nil
}

func (s *Store) UpdateUser(user *togo.User) (*togo.User, error) {
	log.Printf("Saving user %v", user)
	return user, nil
}

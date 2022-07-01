package services

import (
	"errors"
	"math/rand"

	"github.com/lawtrann/togo"
)

type TodoService struct {
	DB togo.TodoDB
}

func NewTodoService(db togo.TodoDB) togo.TodoService {
	return &TodoService{
		DB: db,
	}
}

// Adding new todo task to user if not exceed a limited per day
func (ts *TodoService) AddTodoByUser(userName string, t *togo.Todo) (*togo.Todo, error) {

	// uFlag for checking existed user, isExceed for checking limited per day
	uFlag := false

	// Get user by name
	u, err := ts.DB.GetUserByName(userName)
	if err != nil {
		return &togo.Todo{}, errors.New(err.Error())
	}

	if (togo.User{}) != *u {
		t.UserID = u.ID
		// flag existed user
		uFlag = true
		// checkif exceed a limited per day
		isExceedPerDay, err := ts.DB.IsExceedPerDay(*u)
		if isExceedPerDay || err != nil {
			return &togo.Todo{}, errors.New("you have reached the limit of adding todo task per day")
		}
	} else {
		u.UserName = userName
		u.LimitedPerDay = rand.Intn(9) + 1
	}

	// Add todo with transaction
	ts.DB.AddTodoByUser(u, t, uFlag)

	return t, nil
}

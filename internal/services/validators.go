package services

import (
	"database/sql"
	"errors"
	"net/http"
)

func (s *ToDoService) validateAddTask(req *http.Request, userID, createdDate string) error {
	numOfTasks, err := s.Store.CountTasks(
		req.Context(),
		sql.NullString{
			String: userID,
			Valid:  true,
		},
		sql.NullString{
			String: createdDate,
			Valid:  true,
		},
	)
	if err != nil {
		return err
	}

	maxTodo, err := s.Store.GetMaxToDo(
		req.Context(),
		sql.NullString{
			String: userID,
			Valid:  true,
		},
	)
	if numOfTasks >= maxTodo {
		return errors.New("You reached the limit of tasks per day")
	}

	return nil
}

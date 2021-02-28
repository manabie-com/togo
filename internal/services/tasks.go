package services

import (
	"database/sql"
	"encoding/json"
	"github.com/manabie-com/togo/internal/services/helper"
	"net/http"
	"time"

	"github.com/google/uuid"
	taskmodel "github.com/manabie-com/togo/internal/storages/task/model"
)

const (
	keyTodosMax     = "user_%s_max_todo"
	keyTodosCurrent = "user_%s_current"
)

func (s *ToDoService) ListTasks(resp http.ResponseWriter, req *http.Request) {
	id, _ := userIDFromCtx(req.Context())
	tasks, err := s.taskstore.RetrieveTasks(
		req.Context(),
		sql.NullString{
			String: id,
			Valid:  true,
		},
		value(req, "created_date"),
	)

	resp.Header().Set("Content-Type", "application/json")

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string][]*taskmodel.Task{
		"data": tasks,
	})
}

func (s *ToDoService) AddTask(resp http.ResponseWriter, req *http.Request) {
	t := &taskmodel.Task{}
	err := json.NewDecoder(req.Body).Decode(t)
	defer req.Body.Close()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	now := time.Now()
	userID, _ := userIDFromCtx(req.Context())
	t.ID = uuid.New().String()
	t.UserID = userID
	t.CreatedDate = now.Format("2006-01-02")

	{
		todos, ok := s.mapUserAndTodos[userID]
		if !ok {
			user, err := s.userstore.FindByID(req.Context(),
				sql.NullString{String: userID, Valid: true})
			if err != nil {
				resp.WriteHeader(http.StatusInternalServerError)
				return
			}

			numOfTasks, err := s.taskstore.CountByUserID(req.Context(),
				sql.NullString{String: userID, Valid: true},
				sql.NullString{String: t.CreatedDate, Valid: true})
			if err != nil {
				resp.WriteHeader(http.StatusInternalServerError)
				return
			}

			todos = helper.NewToDos(user.MaxTodo, numOfTasks)
			s.mapUserAndTodos[userID] = todos
		}

		if !todos.CanAddNewTodo() {
			resp.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(resp).Encode(map[string]string{
				"error": "the number of todos daily limit is reached",
			})
			return
		}
	}

	resp.Header().Set("Content-Type", "application/json")

	err = s.taskstore.AddTask(req.Context(), t)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string]*taskmodel.Task{
		"data": t,
	})
}

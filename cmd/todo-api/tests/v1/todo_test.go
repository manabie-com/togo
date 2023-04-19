package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/manabie-com/togo/cmd/todo-api/handlers"
	"github.com/manabie-com/togo/cmd/todo-api/handlers/v1/todogrp"
	"github.com/manabie-com/togo/internal/data/dbtest"
	createtodo "github.com/manabie-com/togo/internal/features/create_todo"
	"github.com/manabie-com/togo/platform/validate"
	"github.com/stretchr/testify/require"
)

// TodoTests holds methods for each todo subtest. This type allows passing
// dependencies for tests while still providing a convenient syntax when
// subtests are registered.
type TodoTests struct {
	app       http.Handler
	userToken string
}

// Test_Todos is the entry point for testing todo management functions.
func Test_Todos(t *testing.T) {
	t.Parallel()

	test := dbtest.NewIntegration(t, c)
	defer func() {
		if r := recover(); r != nil {
			t.Error(r)
		}
		test.Teardown()
	}()

	tests := TodoTests{
		app: handlers.NewApp(handlers.AppConfig{
			Log:  test.Log,
			DB:   test.DB,
			Auth: test.Auth,
		}),
		userToken: test.Token("user@example.com", "gophers"),
	}

	t.Run("postTodo400", tests.postTodo400)
	t.Run("postTodo409", tests.postTodo409)
	// more subtests ...
}

// postTodo400 validates a todo can't be created with the endpoint
// unless a valid todo document is submitted.
func (tt *TodoTests) postTodo400(t *testing.T) {
	require := require.New(t)

	newTodo := todogrp.AppNewTodo{
		Title:   "",
		Content: "",
	}

	body, err := json.Marshal(newTodo)
	require.NoError(err)

	r := httptest.NewRequest(http.MethodPost, "/v1/todos", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	r.Header.Set("Authorization", "Bearer "+tt.userToken)
	r.Header.Set("Content-Type", "application/json")
	tt.app.ServeHTTP(w, r)

	require.Equal(http.StatusBadRequest, w.Code)

	var got validate.FieldErrors
	require.NoError(json.NewDecoder(w.Body).Decode(&got))
	exp := validate.FieldErrors{
		{Field: "title", Err: "title is a required field"},
	}
	require.Equal(exp, got)

}

// postTodo409 validates a todo can't be created with the endpoint
// if the number of todos already exceeds maximum.
func (tt *TodoTests) postTodo409(t *testing.T) {
	require := require.New(t)

	// This user has maximum daily todo of 1, so the first call should succeed.
	{
		newTodo := todogrp.AppNewTodo{
			Title:   "title 1",
			Content: "content 1",
		}

		body, err := json.Marshal(newTodo)
		require.NoError(err)

		r := httptest.NewRequest(http.MethodPost, "/v1/todos", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		r.Header.Set("Authorization", "Bearer "+tt.userToken)
		r.Header.Set("Content-Type", "application/json")
		tt.app.ServeHTTP(w, r)

		require.Equal(http.StatusCreated, w.Code)

		var respBody createtodo.Todo
		require.NoError(json.NewDecoder(w.Body).Decode(&respBody))

		got := todogrp.AppNewTodo{
			Title:   respBody.Title,
			Content: respBody.Content,
		}
		exp := newTodo
		require.Equal(exp, got)
		require.NotEqual(respBody.ID, uuid.Nil)
	}

	// This call should fail with 409 status code
	{
		newTodo := todogrp.AppNewTodo{
			Title:   "title 2",
			Content: "content 2",
		}

		body, err := json.Marshal(newTodo)
		require.NoError(err)

		r := httptest.NewRequest(http.MethodPost, "/v1/todos", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		r.Header.Set("Authorization", "Bearer "+tt.userToken)
		r.Header.Set("Content-Type", "application/json")
		tt.app.ServeHTTP(w, r)

		require.Equal(http.StatusConflict, w.Code)

		var respBody echo.HTTPError
		require.NoError(json.NewDecoder(w.Body).Decode(&respBody))

		got := respBody.Message
		exp := createtodo.ErrExceededDailyMaximumTodos.Error()
		require.Equal(exp, got)
	}
}

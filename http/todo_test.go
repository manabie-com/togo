package http

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/lawtrann/togo"
	"github.com/lawtrann/togo/mocks"
)

func TestHandlerTodoAdd_NewUser_NewTodo(t *testing.T) {
	// Mock TodoService
	svc := mocks.TodoService{}
	svc.AddFn = func(ctx context.Context, t *togo.Todo, username string) (*togo.Todo, error) {
		return &togo.Todo{ID: 1, Description: "Todo something"}, nil
	}
	// New Test Server
	s := NewServer()
	s.TodoService = &svc

	// New request
	r := httptest.NewRequest(http.MethodPost, "/api/lawtrann/todos", strings.NewReader(`{"description":"Todo something"}`))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// s.HandlerTodoAdd(w, r)
	s.Router.ServeHTTP(w, r)

	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)

	got := strings.TrimRight(string(data), "\n")
	expected := `{"Status":201,"Message":"Successfully adding new todo task","Data":{"id":1,"description":"Todo something"}}`

	if got != expected {
		t.Log(got, err, expected)
		t.Error("Error while running TestHandlerTodoAdd_NewUser_NewTodo")
	}
}

func TestHandlerTodoAdd_ExistedUser_NewTodo(t *testing.T) {
	// Mock TodoService
	svc := mocks.TodoService{}
	svc.AddFn = func(ctx context.Context, t *togo.Todo, username string) (*togo.Todo, error) {
		return &togo.Todo{ID: 2, Description: "Todo something"}, nil
	}
	// New Test Server
	s := NewServer()
	s.TodoService = &svc

	// New request
	r := httptest.NewRequest(http.MethodPost, "/api/lawtrann/todos", strings.NewReader(`{"description":"Todo something"}`))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// s.HandlerTodoAdd(w, r)
	s.Router.ServeHTTP(w, r)

	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)

	got := strings.TrimRight(string(data), "\n")
	expected := `{"Status":201,"Message":"Successfully adding new todo task","Data":{"id":2,"description":"Todo something"}}`

	if got != expected {
		t.Log(got, err, expected)
		t.Error("Error while running TestHandlerTodoAdd_ExistedUser_NewTodo")
	}
}

func TestHandlerTodoAdd_ExistedUser_IsExceed(t *testing.T) {
	// Mock TodoService
	svc := mocks.TodoService{}
	svc.AddFn = func(ctx context.Context, t *togo.Todo, username string) (*togo.Todo, error) {
		return &togo.Todo{}, togo.ErrIsExceedLimitedPerDay
	}
	// New Test Server
	s := NewServer()
	s.TodoService = &svc

	// New request
	r := httptest.NewRequest(http.MethodPost, "/api/lawtrann/todos", strings.NewReader(`{"description":"Todo something"}`))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// s.HandlerTodoAdd(w, r)
	s.Router.ServeHTTP(w, r)

	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)

	got := strings.TrimRight(string(data), "\n")
	expected := `{"Status":200,"Message":"You have reached the limit of adding todo task per day","Data":{"id":0,"description":""}}`

	if got != expected {
		t.Log(got, err, expected)
		t.Error("Error while running TestHandlerTodoAdd_ExistedUser_IsExceed")
	}
}

func TestHandlerTodoAdd_PageNotFound(t *testing.T) {
	// Mock TodoService
	svc := mocks.TodoService{}
	svc.AddFn = func(ctx context.Context, t *togo.Todo, username string) (*togo.Todo, error) {
		return &togo.Todo{ID: 1, Description: "Todo something"}, nil
	}
	// New Test Server
	s := NewServer()
	s.TodoService = &svc

	// New request
	r := httptest.NewRequest(http.MethodPost, "/api/lawtrann/todosss", strings.NewReader(`{"description":"Todo something"}`))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Set username to context
	ctx := context.WithValue(r.Context(), userName{}, "lawtrann")
	r = r.WithContext(ctx)

	// s.HandleNotFound(w, r)
	s.Router.ServeHTTP(w, r)

	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)

	got := strings.TrimRight(string(data), "\n")
	expected := `{"Status":404,"Message":"Your page cannot be found","Data":{"id":0,"description":""}}`

	if got != expected {
		t.Log(got, err, expected)
		t.Error("Error while running TestHandlerTodoAdd_PageNotFound")
	}
}

func TestHandlerTodoAdd_HTTPMethodNotAllowed(t *testing.T) {
	// Mock TodoService
	svc := mocks.TodoService{}
	svc.AddFn = func(ctx context.Context, t *togo.Todo, username string) (*togo.Todo, error) {
		return &togo.Todo{ID: 1, Description: "Todo something"}, nil
	}
	// New Test Server
	s := NewServer()
	s.TodoService = &svc

	// New request
	r := httptest.NewRequest(http.MethodGet, "/api/lawtrann/todos", strings.NewReader(`{"description":"Todo something"}`))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// s.HandlerTodoAdd(w, r)
	s.Router.ServeHTTP(w, r)

	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)

	got := strings.TrimRight(string(data), "\n")
	expected := `{"Status":405,"Message":"HTTP Method is not allowed","Data":{"id":0,"description":""}}`

	if got != expected {
		t.Log(got, err, expected)
		t.Error("Error while running TestHandlerTodoAdd_HTTPMethodNotAllowed")
	}
}

func TestHandlerTodoAdd_InvalidBody(t *testing.T) {
	// Mock TodoService
	svc := mocks.TodoService{}
	svc.AddFn = func(ctx context.Context, t *togo.Todo, username string) (*togo.Todo, error) {
		return &togo.Todo{ID: 1, Description: "Todo something"}, nil
	}
	// New Test Server
	s := NewServer()
	s.TodoService = &svc

	// New request
	r := httptest.NewRequest(http.MethodPost, "/api/lawtrann/todos", strings.NewReader(`{"test":""}`))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// s.HandlerTodoAdd(w, r)
	s.Router.ServeHTTP(w, r)

	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)

	got := strings.TrimRight(string(data), "\n")
	expected := `{"Status":422,"Message":"Could not parse Todo object","Data":{"id":0,"description":""}}`

	if got != expected {
		t.Log(got, err, expected)
		t.Error("Error while running TestHandlerTodoAdd_InvalidBody")
	}
}

func TestHandlerTodoAdd_EmptyDescription(t *testing.T) {
	// Mock TodoService
	svc := mocks.TodoService{}
	svc.AddFn = func(ctx context.Context, t *togo.Todo, username string) (*togo.Todo, error) {
		return &togo.Todo{ID: 1, Description: "Todo something"}, nil
	}
	// New Test Server
	s := NewServer()
	s.TodoService = &svc

	// New request
	r := httptest.NewRequest(http.MethodPost, "/api/lawtrann/todos", strings.NewReader(`{"description":""}`))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// s.HandlerTodoAdd(w, r)
	s.Router.ServeHTTP(w, r)

	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)

	got := strings.TrimRight(string(data), "\n")
	expected := `{"Status":422,"Message":"Could not parse Todo object","Data":{"id":0,"description":""}}`

	if got != expected {
		t.Log(got, err, expected)
		t.Error("Error while running TestHandlerTodoAdd_EmptyDescription")
	}
}

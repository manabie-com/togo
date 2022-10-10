package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"togo/domain/controllers"
	"togo/domain/models"
	"togo/utils"
)

func NewTodoHandler(todoController controllers.TodoController) func(r chi.Router) {
	th := &todoHandler{todoController: todoController}

	return func(r chi.Router) {
		r.Post("/todo", th.CreateTodo)
	}
}

type todoHandler struct {
	todoController controllers.TodoController
}

type TodoHandler interface {
	CreateTodo(w http.ResponseWriter, r *http.Request)
}

func (th *todoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var newTodo *models.NewTodo

	err := json.NewDecoder(r.Body).Decode(newTodo)
	if err != nil {
		http.Error(w, fmt.Sprintf("error todoHandler.CreateTodo: %v", err), http.StatusBadRequest)
		return
	}

	todoCreated, err := th.todoController.CreateTodo(r.Context(), newTodo)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
		return
	}

	utils.ResponseWithJSON(w, http.StatusOK, todoCreated)
}

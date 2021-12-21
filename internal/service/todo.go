package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/manabie-com/togo/api/model"
	todorepo "github.com/manabie-com/togo/internal/repo/todo"
)

var (
	ErrUnableToAssignID = errors.New("unable to assign id to new todo")
	ErrUnableToAddTodo  = errors.New("unable to add new todo")
)

type TodoService interface {
	Add(model.Todo) (string, error)
	Delete(int) (bool, error)
	Get([]string) ([]model.Todo, error)
	GetOne(string) (model.Todo, error)
	Update(model.Todo) (bool, error)
}

type DefaultTodo struct {
	Repo todorepo.TodoCrud
}

func (d *DefaultTodo) Add(m model.Todo) (string, error) {
	u, err := uuid.NewUUID()
	t := time.Now()
	if err != nil {
		return "", ErrUnableToAssignID
	}
	m.ID = u.String()
	m.CreatedDate = t.Format("12-12-2006")
	a, err := d.Repo.Add(m)
	if err != nil {
		return "", ErrUnableToAddTodo
	}
	return a.ID, nil
}

func (d *DefaultTodo) Delete(int) (bool, error) {
	return false, nil
}

func (d *DefaultTodo) Get(ids []string) ([]model.Todo, error) {
	r, err := d.Repo.Get(ids)
	if err != nil {
		return []model.Todo{}, errors.New("unable get todos")
	}
	return r, nil
}

func (d *DefaultTodo) GetOne(string) (model.Todo, error) {
	return model.Todo{}, nil
}

func (d *DefaultTodo) Update(model.Todo) (bool, error) {
	return false, nil
}

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

const dateFormat = "2006-01-02"

type TodoService interface {
	Add(model.Todo) (model.Todo, error)
	Delete(int) (bool, error)
	Get(string) ([]model.Todo, error)
	GetOne(string) (model.Todo, error)
	Update(model.Todo) (bool, error)
}

type DefaultTodo struct {
	Repo todorepo.TodoCrud
}

func (d *DefaultTodo) Add(m model.Todo) (model.Todo, error) {
	u, err := uuid.NewUUID()
	t := time.Now()
	if err != nil {
		return model.Todo{}, ErrUnableToAssignID
	}
	m.ID = u.String()
	m.CreatedDate = t.Format(dateFormat)

	a, err := d.Repo.Add(m)
	if err != nil {
		return model.Todo{}, ErrUnableToAddTodo
	}
	return a, nil
}

func (d *DefaultTodo) Delete(int) (bool, error) {
	return false, nil
}

func (d *DefaultTodo) Get(uid string) ([]model.Todo, error) {
	r, err := d.Repo.Get(uid)
	if err != nil {
		return []model.Todo{}, errors.New("unable get todos for user")
	}
	return r, nil
}

func (d *DefaultTodo) GetOne(string) (model.Todo, error) {
	return model.Todo{}, nil
}

func (d *DefaultTodo) Update(model.Todo) (bool, error) {
	return false, nil
}

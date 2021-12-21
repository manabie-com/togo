package service

import (
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/manabie-com/togo/api/model"
	todorepo "github.com/manabie-com/togo/internal/repo/todo"
)

var (
	ErrUnableToAssignID    = errors.New("unable to assign id to new todo")
	ErrUnableToAddTodo     = errors.New("unable to add new todo")
	ErrUserExceedDailyTodo = errors.New("maximum number of todo reached for today")
)

const (
	dateFormat = "2006-01-02"
	capPerUser = 5
)

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

// Add implementation of TODO service that creates new todo item with the designated user
func (d *DefaultTodo) Add(m model.Todo) (model.Todo, error) {

	//validate user todo for the day
	t := time.Now()
	m.CreatedDate = t.Format(dateFormat)
	tdt, err := d.Repo.GetByUserAndDate(m.UserID, m.CreatedDate)
	if err != nil {
		log.Print(err)
		return model.Todo{}, ErrUnableToAssignID
	}
	if len(tdt) >= 5 {
		return model.Todo{}, ErrUserExceedDailyTodo
	}

	// Generate todo id
	u, err := uuid.NewUUID()
	if err != nil {
		return model.Todo{}, ErrUnableToAssignID
	}
	m.ID = u.String()

	a, err := d.Repo.Add(m)
	if err != nil {
		return model.Todo{}, ErrUnableToAddTodo
	}
	return a, nil
}

func (d *DefaultTodo) Delete(int) (bool, error) {
	return false, nil
}

// Get todo by userid
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

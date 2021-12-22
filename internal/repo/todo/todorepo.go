package todorepo

import (
	"github.com/manabie-com/togo/api/model"
)

type TodoCrud interface {
	Add(model.Todo) (model.Todo, error)
	Update(model.Todo) (int, error)
	Delete(string) error
	GetOne(string) (model.Todo, error)
	GetByUserAndDate(ID, date string) ([]model.Todo, error)
	Get(string) ([]model.Todo, error)
}

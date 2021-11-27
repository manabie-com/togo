package store

import (
	"main/config"
	"main/internal/model"
)

type StorageType = string

const (
	InMemory = StorageType("InMemory")
)

type Store interface {
	CreateTodo(todo model.Todo) (model.Todo, error)
}

type InMemoryStore struct {
	idCounter int
	todoTable map[int]model.Todo
}

func (s InMemoryStore) CreateTodo(todo model.Todo) (model.Todo, error) {
	s.todoTable[s.idCounter] = todo
	return s.todoTable[todo.Id], nil
}

func NewStorage(cfg config.Config) (Store, error) {
	switch cfg.StorageType {
	case InMemory:
		return newInMemoryStorage()
	default:
		return newInMemoryStorage()
	}
}

func newInMemoryStorage() (Store, error) {
	table := make(map[int]model.Todo, 0)
	return InMemoryStore{todoTable: table, idCounter: 1}, nil
}

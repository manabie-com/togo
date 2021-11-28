package store

import (
	"errors"
	"main/config"
	"main/internal/model"
)

type StorageType = string

const (
	InMemory = StorageType("InMemory")
)

type Store interface {
	CreateTodo(todo model.Todo) (model.Todo, error)
	GetTodo(id uint) (model.Todo, error)
	GetAllTodo() ([]model.Todo, error)
}

var idCounter = uint(1)

type InMemoryStore struct {
	todoTable map[uint]model.Todo
	idCounter uint
}

func (s *InMemoryStore) CreateTodo(todo model.Todo) (model.Todo, error) {
	todo.Id = s.idCounter
	s.todoTable[idCounter] = todo
	s.idCounter = idCounter + 1
	return s.todoTable[todo.Id], nil
}

func (s *InMemoryStore) GetTodo(id uint) (model.Todo, error) {
	todo, ok := s.todoTable[id]
	if !ok {
		return model.Todo{}, errors.New("not found")
	}

	return todo, nil
}

func (s *InMemoryStore) GetAllTodo() ([]model.Todo, error) {
	var todos []model.Todo
	for _, todo := range s.todoTable {
		todos = append(todos, todo)
	}

	return todos, nil
}

func NewStorage(cfg config.Config) (Store, error) {
	switch cfg.StorageType {
	case InMemory:
		return newInMemoryStorage()
	default:
		return newInMemoryStorage()
	}
}

func newInMemoryStorage() (*InMemoryStore, error) {
	table := make(map[uint]model.Todo, 0)
	return &InMemoryStore{todoTable: table, idCounter: 1}, nil
}

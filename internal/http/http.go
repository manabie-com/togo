package http

import (
	"github.com/manabie/project/internal/usecase"
	"github.com/manabie/project/model"
)

type http struct {
	usecase usecase.Usecase
}

type Http interface {
	Login(user model.User) (string , error)
	SignUp(user model.User) error
	CreateTask(task model.Task, idUser int) error
	UpdateTask(id int, task model.Task) error
	DeleteTask(id int) error
	TaskAll() ([]model.Task, error)
	TaskById(id int) (model.Task, error)
}

func NewHttp(usecase usecase.Usecase) Http {
	return &http{
		usecase:usecase,
	}
}

func(h *http) Login(user model.User) (string , error) {
	token, err := h.usecase.Login(user)
	if err != nil {
		return "", err
	}
	return token, nil
}

func(h *http) SignUp(user model.User) error {
	if err := h.usecase.SignUp(user); err != nil {
		return err
	}
	return nil
}

func(h *http) CreateTask(task model.Task, idUser int) error {
	if err := h.usecase.CreateTask(task, idUser); err != nil {
		return err
	}
	return nil
}

func(h *http) UpdateTask(id int, task model.Task) error {
	if err := h.usecase.UpdateTask(id, task); err != nil {
		return err
	}
	return nil
}

func(h *http) DeleteTask(id int) error {
	if err := h.usecase.DeleteTask(id); err != nil {
		return err
	}
	return nil
}

func(h *http) TaskAll() ([]model.Task, error) {
	tasks, err := h.usecase.TaskAll()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func(h *http) TaskById(id int) (model.Task, error) {
	task, err := h.usecase.TaskById(id)
	if err != nil {
		return model.Task{}, err
	}
	return task, nil
}
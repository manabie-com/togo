package test

import (
	"main/config"
	"main/internal/logger"
	"main/internal/model"
	"main/internal/service"
	"main/internal/store"
	"reflect"
	"testing"
)

func TestCreateTodoService(t *testing.T) {
	cfg := config.Config{
		Base: config.Base{
			HTTPAddress: 9000,
		},
		Storage: config.Storage{
			StorageType: "InMem",
		},
	}

	log := logger.New()

	storage, err := store.NewStorage(cfg)
	if err != nil {
		t.Fatal("connect to store fail:", logger.Object("error", err))
	}

	svc, err := service.NewTogoService(cfg, storage, log)
	if err != nil {
		t.Fatal("create service fail:", logger.Object("error", err))
	}

	req := &service.CreateTodoRequest{
		Title:  "make breakfast",
		UserId: 1,
	}

	expected := &service.CreateTodoResponse{Todo: model.Todo{
		Id:     1,
		Title:  "make breakfast",
		UserId: 1,
	}}

	resp, err := svc.CreateTodo(req)
	if err != nil {
		t.Fatal("send request to service fail:", logger.Object("error", err))
	}

	compare(t, resp, expected)
}

func compare(t *testing.T, a interface{}, b interface{}) {
	if !reflect.DeepEqual(a, b) {
		t.Fatalf("found: %+v, expected: %+v\n", a, b)
	}
}

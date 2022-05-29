package inmemory_test

import (
	"context"
	"testing"
	"togo/domain/model"
	"togo/infrastructure/inmemory"
)

func TestNewInMemoryUserRepo(t *testing.T) {
	c := inmemory.NewInMemoryUserRepo()
	t.Logf("Inmemory repository: %#v", c)
}

func TestInmemoryUserRepo_Create(t *testing.T) {
	repo := inmemory.NewInMemoryUserRepo()
	user := model.User{
		Username: "admin",
		Password: "admin",
		Limit:    10,
	}
	_ = repo.Delete(context.Background(), "admin")
	err := repo.Create(context.Background(), user)
	if err != nil {
		t.Errorf("%s", err.Error())
		return
	}
	t.Log("Create user success")
}

func TestInmemoryUserRepo_Get(t *testing.T) {
	repo := inmemory.NewInMemoryUserRepo()
	user := model.User{
		Username: "admin",
		Password: "admin",
		Limit:    10,
	}
	err := repo.Create(context.Background(), user)
	if err != nil {
		t.Errorf("%s", err.Error())
		return
	}
	u2, err := repo.Get(context.Background(), user.Username)
	if err != nil {
		t.Errorf("%s", err.Error())
		return
	}
	t.Logf("User: %#v", u2)
}

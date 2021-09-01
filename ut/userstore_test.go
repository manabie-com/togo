package ut

import (
	"context"
	"log"
	"reflect"
	"testing"

	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/storages/repo"
	"github.com/manabie-com/togo/utils"
)

func InitUserStore() *repo.UserStore {
	db, err := utils.InitDB()
	if err != nil {
		log.Fatal("error opening db", err)
	}
	return &repo.UserStore{
		DB: db,
	}
}

func TestRetrieveUser(t *testing.T) {
	u := InitUserStore()
	user := &storages.User{
		ID:       "testUser",
		Password: "$2a$10$fYyWHHPGc3XhZjVYPiS7Y.f8LSfcJB3PgIiH9GuuSEOXubJ1y34Su",
		MaxTodo:  5,
	}
	ctx := context.Background()
	su, err := u.RetrieveUser(ctx, user.ID)
	if err != nil {
		t.Errorf("RetrieveUser returned err %v", err)
		return
	}

	if !reflect.DeepEqual(su, user) {
		t.Errorf("RetrieveUser returned wrong: got %v want %v", su, user)
	}
}

func TestAddUser(t *testing.T) {
	u := InitUserStore()
	user_id := RandStringRunes(32)
	password := "example"

	ctx := context.Background()
	err := u.AddUser(ctx, user_id, password)
	if err != nil {
		t.Errorf("AddUser returned err %v", err)
		return
	}
}

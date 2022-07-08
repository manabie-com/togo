package db

import (
	"context"
	"testing"
	"todo/be/utils"

	"go.mongodb.org/mongo-driver/bson"
)

func TestTodo_add(t *testing.T) {
	successDb := InitDb()
	if !successDb {
		t.Errorf("Output expect true instead of false")
	}
	result := Todo_add(context.TODO(), Todo{UserId: "TestDbUserId"})
	if utils.IsError(result) {
		t.Errorf("Output expect to not error")
	}
}

func TestTodo_count(t *testing.T) {
	successDb := InitDb()
	if !successDb {
		t.Errorf("Output expect true instead of false")
	}
	result := Todo_count(context.TODO(), bson.M{})
	if result < 0 {
		t.Errorf("Output expect >= 0")
	}
}

func TestTodo_delete(t *testing.T) {
	successDb := InitDb()
	if !successDb {
		t.Errorf("Output expect true instead of false")
	}
	result := Todo_delete(context.TODO(), bson.M{Todo_UserId: "TestDbUserId"})
	if utils.IsError(result) {
		t.Errorf("Output expect to not error")
	}
}

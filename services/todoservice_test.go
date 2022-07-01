package services

import (
	"testing"

	"github.com/lawtrann/togo"
	"github.com/lawtrann/togo/mocks"
)

func TestAddTodoByUser_WithNewUser(t *testing.T) {
	// Create DB mocks
	db := mocks.TodoDB{}

	// New User
	db.GetUserByNameResp = togo.User{}
	db.GetUserByNameErr = nil
	// Add todo with transaction
	db.AddTodoByUserErr = nil

	// mocks DB to TodoService
	svc := NewTodoService(db)

	// Input data
	username := "test1"
	todo := &togo.Todo{
		TodoID:      1,
		UserID:      1,
		Description: "Description 1",
	}
	// Got result
	got, err := svc.AddTodoByUser(username, todo)
	// Expected result
	expected := togo.Todo{
		TodoID:      1,
		UserID:      1,
		Description: "Description 1",
	}

	if *got != expected && err != nil {
		t.Log(*got)
		t.Error("Error while inserted new Todo!")
	}
}

func TestAddTodoByUser_WithExistedUser_Exceed(t *testing.T) {
	// Create DB mocks
	db := mocks.TodoDB{}

	// Existed User
	db.GetUserByNameResp = togo.User{
		ID:            1,
		UserName:      "test1",
		LimitedPerDay: 3,
	}
	db.GetUserByNameErr = nil
	// IsExceed = true
	db.IsExceedPerDayResp = true // check point
	db.IsExceedPerDayErr = nil
	// Add todo with transaction
	db.AddTodoByUserErr = nil

	// mocks DB to TodoService
	svc := NewTodoService(db)

	// Input data
	username := "test1"
	todo := &togo.Todo{
		TodoID:      1,
		UserID:      1,
		Description: "Description 1",
	}
	// Got result
	got, err := svc.AddTodoByUser(username, todo)
	// Expected result
	expected := togo.Todo{}

	if *got != expected && err != nil {
		t.Log(*got, expected)
		t.Error("The number of tasks exceeds a limited per day")
	}
}

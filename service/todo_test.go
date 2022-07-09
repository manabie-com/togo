package service

import (
	"context"
	"errors"
	"testing"

	"github.com/lawtrann/togo"
	"github.com/lawtrann/togo/mocks"
)

func TestAdd_NewUser_NewTodo(t *testing.T) {
	// Input data
	ctx := context.Background()
	todo := togo.Todo{
		Description: "Todo something",
	}
	userName := "lawtrann"

	// Mock UserRepo
	userRepo := mocks.UserRepo{}
	userRepo.GetUserByNameFn = func(ctx context.Context, username string) (*togo.User, error) {
		return &togo.User{}, nil
	}
	// Create UserService
	userService := NewUserService(&userRepo)

	// Mock TodoRepo
	todoRepo := mocks.TodoRepo{}
	todoRepo.AddWithNewUserFn = func(ctx context.Context, t *togo.Todo, u *togo.User) (*togo.Todo, error) {
		return &togo.Todo{ID: 1, Description: "Todo something"}, nil
	}

	// Create TodoService
	svc := NewTodoService(&todoRepo)
	svc.UserService = userService

	// Got result
	got, err := svc.Add(ctx, &todo, userName)
	// Expect result
	expect := togo.Todo{ID: 1, Description: "Todo something"}
	if *got != expect || err != nil {
		t.Log(*got, err, expect)
		t.Error("Error while running TestAdd_NewUser_NewTodo")
	}
}

func TestAdd_ExistedUser_NewTodo_NotExceed(t *testing.T) {
	// Input data
	ctx := context.Background()
	todo := togo.Todo{
		Description: "Todo something",
	}
	userName := "lawtrann"

	// Mock UserRepo
	userRepo := mocks.UserRepo{}
	userRepo.GetUserByNameFn = func(ctx context.Context, username string) (*togo.User, error) {
		return &togo.User{ID: 1, Username: "lawtrann", LimitedPerDay: 5}, nil
	}
	userRepo.IsExceedPerDayFn = func(ctx context.Context, u *togo.User) (bool, error) {
		return false, nil
	}
	// Create UserService
	userService := NewUserService(&userRepo)

	// Mock TodoRepo
	todoRepo := mocks.TodoRepo{}
	todoRepo.AddFn = func(ctx context.Context, t *togo.Todo, u *togo.User) (*togo.Todo, error) {
		return &togo.Todo{ID: 2, Description: "Todo something"}, nil
	}

	// Create TodoService
	svc := NewTodoService(&todoRepo)
	svc.UserService = userService

	// Got result
	got, err := svc.Add(ctx, &todo, userName)
	// Expect result
	expect := togo.Todo{ID: 2, Description: "Todo something"}
	if *got != expect || err != nil {
		t.Log(*got, err, expect)
		t.Error("Error while running TestAdd_ExistedUser_NewTodo_NotExceed")
	}
}

func TestAdd_ExistedUser_NewTodo_WithExceed(t *testing.T) {
	// Input data
	ctx := context.Background()
	todo := togo.Todo{
		Description: "Todo something",
	}
	userName := "lawtrann"

	// Mock UserRepo
	userRepo := mocks.UserRepo{}
	userRepo.GetUserByNameFn = func(ctx context.Context, username string) (*togo.User, error) {
		return &togo.User{ID: 1, Username: "lawtrann", LimitedPerDay: 5}, nil
	}
	userRepo.IsExceedPerDayFn = func(ctx context.Context, u *togo.User) (bool, error) {
		return true, nil
	}
	// Create UserService
	userService := NewUserService(&userRepo)

	// Mock TodoRepo
	todoRepo := mocks.TodoRepo{}
	todoRepo.AddFn = func(ctx context.Context, t *togo.Todo, u *togo.User) (*togo.Todo, error) {
		return &togo.Todo{ID: 2, Description: "Todo something"}, nil
	}

	// Create TodoService
	svc := NewTodoService(&todoRepo)
	svc.UserService = userService

	// Got result
	got, err := svc.Add(ctx, &todo, userName)
	// Expect result
	expect := togo.Todo{}
	if *got != expect && err != nil {
		t.Log(*got, err, expect)
		t.Error("Error while running TestAdd_ExistedUser_NewTodo_WithExceed")
	}
}

func TestAdd_ExistedUser_NewTodo_WithGetUserByNameFnError(t *testing.T) {
	// Input data
	ctx := context.Background()
	todo := togo.Todo{
		Description: "Todo something",
	}
	userName := "lawtrann"

	// Mock UserRepo
	userRepo := mocks.UserRepo{}
	userRepo.GetUserByNameFn = func(ctx context.Context, username string) (*togo.User, error) {
		return &togo.User{ID: 1, Username: "lawtrann", LimitedPerDay: 5}, errors.New("Error occurs")
	}
	userRepo.IsExceedPerDayFn = func(ctx context.Context, u *togo.User) (bool, error) {
		return false, nil
	}
	// Create UserService
	userService := NewUserService(&userRepo)

	// Mock TodoRepo
	todoRepo := mocks.TodoRepo{}
	todoRepo.AddFn = func(ctx context.Context, t *togo.Todo, u *togo.User) (*togo.Todo, error) {
		return &togo.Todo{ID: 2, Description: "Todo something"}, nil
	}

	// Create TodoService
	svc := NewTodoService(&todoRepo)
	svc.UserService = userService

	// Got result
	got, err := svc.Add(ctx, &todo, userName)
	// Expect result
	expect := togo.Todo{}
	if *got != expect && err != nil {
		t.Log(*got, err, expect)
		t.Error("Error while running TestAdd_ExistedUser_NewTodo_WithGetUserByNameFnError")
	}
}

func TestAdd_ExistedUser_NewTodo_WithIsExceedPerDayFnError(t *testing.T) {
	// Input data
	ctx := context.Background()
	todo := togo.Todo{
		Description: "Todo something",
	}
	userName := "lawtrann"

	// Mock UserRepo
	userRepo := mocks.UserRepo{}
	userRepo.GetUserByNameFn = func(ctx context.Context, username string) (*togo.User, error) {
		return &togo.User{ID: 1, Username: "lawtrann", LimitedPerDay: 5}, nil
	}
	userRepo.IsExceedPerDayFn = func(ctx context.Context, u *togo.User) (bool, error) {
		return false, errors.New("Error occurs")
	}
	// Create UserService
	userService := NewUserService(&userRepo)

	// Mock TodoRepo
	todoRepo := mocks.TodoRepo{}
	todoRepo.AddFn = func(ctx context.Context, t *togo.Todo, u *togo.User) (*togo.Todo, error) {
		return &togo.Todo{ID: 2, Description: "Todo something"}, nil
	}

	// Create TodoService
	svc := NewTodoService(&todoRepo)
	svc.UserService = userService

	// Got result
	got, err := svc.Add(ctx, &todo, userName)
	// Expect result
	expect := togo.Todo{}
	if *got != expect && err != nil {
		t.Log(*got, err, expect)
		t.Error("Error while running TestAdd_ExistedUser_NewTodo_WithIsExceedPerDayFnError")
	}
}

func TestAdd_ExistedUser_NewTodo_WithAddFnError(t *testing.T) {
	// Input data
	ctx := context.Background()
	todo := togo.Todo{
		Description: "Todo something",
	}
	userName := "lawtrann"

	// Mock UserRepo
	userRepo := mocks.UserRepo{}
	userRepo.GetUserByNameFn = func(ctx context.Context, username string) (*togo.User, error) {
		return &togo.User{ID: 1, Username: "lawtrann", LimitedPerDay: 5}, nil
	}
	userRepo.IsExceedPerDayFn = func(ctx context.Context, u *togo.User) (bool, error) {
		return false, nil
	}
	// Create UserService
	userService := NewUserService(&userRepo)

	// Mock TodoRepo
	todoRepo := mocks.TodoRepo{}
	todoRepo.AddFn = func(ctx context.Context, t *togo.Todo, u *togo.User) (*togo.Todo, error) {
		return &togo.Todo{ID: 2, Description: "Todo something"}, errors.New("Error occurs")
	}

	// Create TodoService
	svc := NewTodoService(&todoRepo)
	svc.UserService = userService

	// Got result
	got, err := svc.Add(ctx, &todo, userName)
	// Expect result
	expect := togo.Todo{}
	if *got != expect && err != nil {
		t.Log(*got, err, expect)
		t.Error("Error while running TestAdd_ExistedUser_NewTodo_WithAddFnError")
	}
}

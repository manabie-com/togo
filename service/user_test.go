package service

import (
	"context"
	"errors"
	"testing"

	"github.com/lawtrann/togo"
	"github.com/lawtrann/togo/mocks"
)

func TestGetUserByName_UserNotExisted(t *testing.T) {
	// Input data
	ctx := context.Background()
	userName := "lawtrann"

	// Mock UserRepo
	userRepo := mocks.UserRepo{}
	userRepo.GetUserByNameFn = func(ctx context.Context, username string) (*togo.User, error) {
		return &togo.User{}, nil
	}

	// Create UserService
	svc := NewUserService(&userRepo)

	// Got result
	got, err := svc.GetUserByName(ctx, userName)
	// Expect result
	expect := togo.User{}
	if *got != expect || err != nil {
		t.Log(*got, err, expect)
		t.Error("Error while running TestGetUserByName_UserNotExisteds")
	}
}

func TestGetUserByName_UserExisted(t *testing.T) {
	// Input data
	ctx := context.Background()
	userName := "lawtrann"

	// Mock UserRepo
	userRepo := mocks.UserRepo{}
	userRepo.GetUserByNameFn = func(ctx context.Context, username string) (*togo.User, error) {
		return &togo.User{ID: 1, Username: "lawtrann", LimitedPerDay: 5}, nil
	}

	// Create UserService
	svc := NewUserService(&userRepo)

	// Got result
	got, err := svc.GetUserByName(ctx, userName)
	// Expect result
	expect := togo.User{ID: 1, Username: "lawtrann", LimitedPerDay: 5}
	if *got != expect || err != nil {
		t.Log(*got, err, expect)
		t.Error("Error while running TestGetUserByName_UserExisted")
	}
}

func TestGetUserByName_WithError(t *testing.T) {
	// Input data
	ctx := context.Background()
	userName := "lawtrann"

	// Mock UserRepo
	userRepo := mocks.UserRepo{}
	userRepo.GetUserByNameFn = func(ctx context.Context, username string) (*togo.User, error) {
		return &togo.User{}, errors.New("Error occurs")
	}

	// Create UserService
	svc := NewUserService(&userRepo)

	// Got result
	got, err := svc.GetUserByName(ctx, userName)
	// Expect result
	expect := togo.User{}
	if *got != expect && err != nil {
		t.Log(*got, err, expect)
		t.Error("Error while running TestGetUserByName_WithError")
	}
}

func TestIsExceedPerDay_Exceed(t *testing.T) {
	// Input data
	ctx := context.Background()
	user := togo.User{ID: 1, Username: "lawtrann", LimitedPerDay: 5}

	// Mock UserRepo
	userRepo := mocks.UserRepo{}
	userRepo.IsExceedPerDayFn = func(ctx context.Context, u *togo.User) (bool, error) {
		return true, nil
	}

	// Create UserService
	svc := NewUserService(&userRepo)

	// Got result
	got, err := svc.IsExceedPerDay(ctx, &user)
	// Expect result
	expect := true
	if got != expect || err != nil {
		t.Log(got, err, expect)
		t.Error("Error while running TestIsExceedPerDay_Exceed")
	}
}

func TestIsExceedPerDay_NotExceed(t *testing.T) {
	// Input data
	ctx := context.Background()
	user := togo.User{ID: 1, Username: "lawtrann", LimitedPerDay: 5}

	// Mock UserRepo
	userRepo := mocks.UserRepo{}
	userRepo.IsExceedPerDayFn = func(ctx context.Context, u *togo.User) (bool, error) {
		return false, nil
	}

	// Create UserService
	svc := NewUserService(&userRepo)

	// Got result
	got, err := svc.IsExceedPerDay(ctx, &user)
	// Expect result
	expect := false
	if got != expect || err != nil {
		t.Log(got, err, expect)
		t.Error("Error while running TestIsExceedPerDay_NotExceed")
	}
}

func TestIsExceedPerDay_WithError(t *testing.T) {
	// Input data
	ctx := context.Background()
	user := togo.User{ID: 1, Username: "lawtrann", LimitedPerDay: 5}

	// Mock UserRepo
	userRepo := mocks.UserRepo{}
	userRepo.IsExceedPerDayFn = func(ctx context.Context, u *togo.User) (bool, error) {
		return false, errors.New("Error occurs")
	}

	// Create UserService
	svc := NewUserService(&userRepo)

	// Got result
	got, err := svc.IsExceedPerDay(ctx, &user)
	// Expect result
	expect := false
	if got != expect && err != nil {
		t.Log(got, err, expect)
		t.Error("Error while running TestIsExceedPerDay_NotExceed")
	}
}

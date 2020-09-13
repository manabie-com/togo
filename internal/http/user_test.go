package http

import (
  "context"
  "github.com/manabie-com/togo/internal/core"
  "testing"
)

type dummyUserRepo struct{}

var dummyUser = core.User{
  ID:      "dummy",
  Hash:    "password",
}

func (repo *dummyUserRepo) ById(ctx context.Context, id string) (*core.User, error) {
  if id == "dummy" {
    return &dummyUser, nil
  }
  return nil, core.ErrUserNotFound
}

func (repo *dummyUserRepo) Validate(ctx context.Context, id string, password string) (*core.User, bool) {
  if id == "dummy" && password == "password" {
    return &dummyUser, true
  }
  return nil, false
}

var userHandler UserHandler

func TestUserHandler_login(t *testing.T) {
  userHandler = UserHandler{
    userRepo:            &dummyUserRepo{},
    generateToken: func(user *core.User) (s string, err error) {
      return "token", nil
    },
  }

  t.Run("Invalid user", func(t *testing.T) {
    token, err := userHandler.login(context.Background(), "admin", "password")
    if token != "" {
      t.Error("Invalid user must generate empty token")
    }
    if err != ErrWrongIdPassword {
      t.Errorf("Wrong error: %v", err)
    }
  })
  t.Run("Wrong password", func(t *testing.T) {
    token, err := userHandler.login(context.Background(), "dummy", "wrongpassword")
    if token != "" {
      t.Error("Wrong password must generate empty token")
    }
    if err != ErrWrongIdPassword {
      t.Errorf("Wrong error: %v", err)
    }
  })
  t.Run("Valid user and correct password", func(t *testing.T) {
    token, err := userHandler.login(context.Background(), "dummy", "password")
    if token != "token" {
      t.Errorf("Expect: token = token, receive: token = %s", token)
    }
    if err != nil {
      t.Error("Error must be nil")
    }
  })
}

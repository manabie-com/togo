package postgresql

import (
  "context"
  "github.com/manabie-com/togo/internal/core"
  "testing"
)

func TestUserRepo_Create(t *testing.T) {
  t.Run("Create with new user", func(t *testing.T) {
    reset()
    err := userRepo.Create(context.Background(), &firstUser, "password")
    if err != nil {
      t.Errorf("Expect no error, got (%v)", err)
    }
  })
  t.Run("Create with existing user", func(t *testing.T) {
    reset()
    err := userRepo.Create(context.Background(), &firstUser, "password")
    if err != nil {
      t.Errorf("Expect no error, got (%v)", err)
    }
    err = userRepo.Create(context.Background(), &firstUser, "password")
    if err != core.ErrUserAlreadyExists {
      t.Errorf("Expect (%v), got (%v)", core.ErrUserAlreadyExists, err)
    }
  })
  t.Run("Create then read", func(t *testing.T) {
    reset()
    err := userRepo.Create(context.Background(), &firstUser, "password")
    if err != nil {
      t.Errorf("Expect no error, got (%v)", err)
    }
    user, err := userRepo.ById(context.Background(), firstUser.ID)
    if err != nil {
      t.Errorf("Expect no error, got (%v)", err)
    }
    if user.MaxTodo != firstUser.MaxTodo {
      t.Errorf("Error max todo - Expect %v, got %v", firstUser.MaxTodo, user.MaxTodo)
    }
    if !userRepo.Compare("password", user.Hash) {
      t.Error("Password and hash mismatch")
    }
  })
  t.Run("Create then validate", func(t *testing.T) {
    reset()
    err := userRepo.Create(context.Background(), &firstUser, "password")
    if err != nil {
      t.Errorf("Expect no error, got (%v)", err)
    }
    user, err :=userRepo.Validate(context.Background(), firstUser.ID, "password")
    if err != nil {
      t.Errorf("Expect no error, got (%v)", err)
    }
    if user.MaxTodo != firstUser.MaxTodo {
      t.Errorf("Error max todo - Expect %v, got %v", firstUser.MaxTodo, user.MaxTodo)
    }
    if !userRepo.Compare("password", user.Hash) {
      t.Error("Password and hash mismatch")
    }
  })
}


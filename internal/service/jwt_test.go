package service

import (
  "context"
  "github.com/manabie-com/togo/internal/core"
  "github.com/manabie-com/togo/internal/memory"
  "testing"
)

var jwtUserService JwtUserService

const fixedToken = "token"
type dummyJwtGenerator struct{}

func (jwtGen *dummyJwtGenerator) Token(user *core.User) (string, error) {
  return fixedToken, nil
}

var noHashUser = core.User{
  ID:      "dummy",
  Hash:    "password",
  MaxTodo: 5,
}

func resetJwtUserService() {
  userRepo := memory.NewUserRepo(
    func(password string) (s string, err error) {
      return password, nil
    },
    func(password, hash string) bool {
      return password == hash
    },
  )
  _ = userRepo.Create(context.Background(), &noHashUser, "password")
  jwtUserService = JwtUserService{
    UserRepo:   userRepo,
    JwtGenerator: &dummyJwtGenerator{},
  }
}

func TestJwtUserService_Login(t *testing.T) {
  resetJwtUserService()

  t.Run("Invalid user", func(t *testing.T) {
    resetJwtUserService()
    token, err := jwtUserService.Login(context.Background(), "admin", "password")
    if token != "" {
      t.Error("Invalid user must generate empty token")
    }
    if err != core.ErrWrongIdPassword {
      t.Errorf("Wrong error - expect %v, got %v", core.ErrWrongIdPassword, err)
    }
  })
  t.Run("Wrong password", func(t *testing.T) {
    resetJwtUserService()
    token, err := jwtUserService.Login(context.Background(), "dummy", "wrongpassword")
    if token != "" {
      t.Error("Wrong password must generate empty token")
    }
    if err != core.ErrWrongIdPassword {
      t.Errorf("Wrong error - expect %v, got %v", core.ErrWrongIdPassword, err)
    }
  })
  t.Run("Valid user and correct password", func(t *testing.T) {
    resetJwtUserService()
    token, err := jwtUserService.Login(context.Background(), "dummy", "password")
    if token != fixedToken {
      t.Errorf("Error: expect %v, got %v", fixedToken, token)
    }
    if err != nil {
      t.Errorf("Error must be nil, got %v", err)
    }
  })
}

func TestJwtUserService_Signup(t *testing.T) {
  resetJwtUserService()

  t.Run("Signup with existed user id", func(t *testing.T) {
    resetJwtUserService()
    token, err := jwtUserService.Signup(context.Background(), &noHashUser, "password")
    if token != "" {
      t.Error("Existed user must generate empty token")
    }
    if err != core.ErrUserAlreadyExists {
      t.Errorf("Wrong error - expect %v, got %v", core.ErrUserAlreadyExists, err)
    }
  })
  t.Run("Signup with new user id", func(t *testing.T) {
    resetJwtUserService()
    token, err := jwtUserService.Signup(context.Background(), &core.User{ID: "newUser", MaxTodo: 1}, "password")
    if token != fixedToken {
      t.Errorf("Error: expect %v, got %v", fixedToken, token)
    }
    if err != nil {
      t.Errorf("Error must be nil, got %v", err)
    }
  })
  t.Run("Signup with empty id", func(t *testing.T) {
    resetJwtUserService()
    token, err := jwtUserService.Signup(context.Background(), &core.User{ID: ""}, "password")
    if token != "" {
      t.Error("Empty id must generate empty token")
    }
    if err != core.ErrEmptyId {
      t.Errorf("Wrong error - expect %v, got %v", core.ErrEmptyId, err)
    }
  })
  t.Run("Signup with empty password", func(t *testing.T) {
    resetJwtUserService()
    token, err := jwtUserService.Signup(context.Background(), &core.User{ID: "user"}, "")
    if token != "" {
      t.Error("Empty password must generate empty token")
    }
    if err != core.ErrEmptyPassword {
      t.Errorf("Wrong error - expect %v, got %v", core.ErrEmptyPassword, err)
    }
  })
  t.Run("Signup with negative max todo", func(t *testing.T) {
    resetJwtUserService()
    token, err := jwtUserService.Signup(context.Background(), &core.User{ID: "user", MaxTodo: -1}, "password")
    if token != "" {
      t.Error("Negative max todo must generate empty token")
    }
    if err != core.ErrInvalidMaxTodo {
      t.Errorf("Wrong error - expect %v, got %v", core.ErrInvalidMaxTodo, err)
    }
  })
  t.Run("Signup with zero max todo", func(t *testing.T) {
    resetJwtUserService()
    token, err := jwtUserService.Signup(context.Background(), &core.User{ID: "user", MaxTodo: 0}, "password")
    if token != "" {
      t.Error("Zero max todo must generate empty token")
    }
    if err != core.ErrInvalidMaxTodo {
      t.Errorf("Wrong error - expect %v, got %v", core.ErrInvalidMaxTodo, err)
    }
  })
  t.Run("Signup with positive max todo", func(t *testing.T) {
    resetJwtUserService()
    token, err := jwtUserService.Signup(context.Background(), &core.User{ID: "user", MaxTodo: 5}, "password")
    if token != fixedToken {
      t.Errorf("Error: expect %v, got %v", fixedToken, token)
    }
    if err != nil {
      t.Errorf("Error must be nil, got %v", err)
    }
  })

  t.Run("Signup then Login", func(t *testing.T) {
    resetJwtUserService()
    // Signup
    token, err := jwtUserService.Signup(context.Background(), &core.User{ID: "user", MaxTodo: 5}, "password")
    if token != fixedToken {
      t.Errorf("Error: expect %v, got %v", fixedToken, token)
    }
    if err != nil {
      t.Errorf("Error must be nil, got %v", err)
    }

    // Login
    token, err = jwtUserService.Login(context.Background(), "user", "password")
    if token != fixedToken {
      t.Errorf("Error: expect %v, got %v", fixedToken, token)
    }
    if err != nil {
      t.Errorf("Error must be nil, got %v", err)
    }
  })
}

package http

import (
  "context"
  "github.com/manabie-com/togo/internal/core"
  "golang.org/x/crypto/bcrypt"
  "testing"
)

type inMemUserRepo struct{
  hash func(password string) (string, error)
  compare func(password, hash string) bool
  users map[string]*core.User
}

var noHashUser = core.User{
  ID:   "dummy",
  Hash: "password",
  MaxTodo: 2,
}

func (repo *inMemUserRepo) Hash(password string) (string, error) {
  return repo.hash(password)
}

func (repo *inMemUserRepo) Compare(password, hash string) bool {
  return repo.compare(password, hash)
}

func (repo *inMemUserRepo) Create(ctx context.Context, user *core.User, password string) (err error) {
  if _, ok := repo.users[user.ID]; ok {
    return core.ErrUserAlreadyExists
  }
  user.Hash, err = repo.hash(password)
  repo.users[user.ID] = user
  return nil
}

func (repo *inMemUserRepo) ById(ctx context.Context, id string) (*core.User, error) {
  if _, ok := repo.users[id]; ok {
    return repo.users[id], nil
  }
  return nil, core.ErrUserNotFound
}

func (repo *inMemUserRepo) Validate(ctx context.Context, id string, password string) (*core.User, error) {
  if u, ok := repo.users[id]; ok && repo.compare(password, u.Hash) {
    return u, nil
  }
  return nil, core.ErrWrongIdPassword
}

var userHandler UserHandler

func resetUserHandlerNoHash() {
  userHandler = UserHandler{
    userRepo: &inMemUserRepo{users: map[string]*core.User{
      noHashUser.ID: &noHashUser,
    }, hash: func(password string) (s string, err error) {
      return password, nil
    }, compare: func(password, hash string) bool {
      return password == hash
    }},
    generateToken: func(user *core.User) (s string, err error) {
      return "token", nil
    },
  }
}

func TestUserHandler_login_nohash(t *testing.T) {
  resetUserHandlerNoHash()

  t.Run("Invalid user", func(t *testing.T) {
    resetUserHandlerNoHash()
    token, err := userHandler.login(context.Background(), "admin", "password")
    if token != "" {
      t.Error("Invalid user must generate empty token")
    }
    if err != core.ErrWrongIdPassword {
      t.Errorf("Wrong error - expect %v, got %v", core.ErrWrongIdPassword, err)
    }
  })
  t.Run("Wrong password", func(t *testing.T) {
    resetUserHandlerNoHash()
    token, err := userHandler.login(context.Background(), "dummy", "wrongpassword")
    if token != "" {
      t.Error("Wrong password must generate empty token")
    }
    if err != core.ErrWrongIdPassword {
      t.Errorf("Wrong error - expect %v, got %v", core.ErrWrongIdPassword, err)
    }
  })
  t.Run("Valid user and correct password", func(t *testing.T) {
    resetUserHandlerNoHash()
    token, err := userHandler.login(context.Background(), "dummy", "password")
    if token != "token" {
      t.Errorf("Expect: token = token, receive: token = %s", token)
    }
    if err != nil {
      t.Errorf("Error must be nil, got %v", err)
    }
  })
}


func TestUserHandler_signup_nohash(t *testing.T) {
  resetUserHandlerNoHash()

  t.Run("Signup with existed user id", func(t *testing.T) {
    resetUserHandlerNoHash()
    token, err := userHandler.signup(context.Background(), &noHashUser, "password")
    if token != "" {
      t.Error("Existed user must generate empty token")
    }
    if err != core.ErrUserAlreadyExists {
      t.Errorf("Wrong error - expect %v, got %v", core.ErrUserAlreadyExists, err)
    }
  })
  t.Run("Signup with new user id", func(t *testing.T) {
    resetUserHandlerNoHash()
    token, err := userHandler.signup(context.Background(), &core.User{ID:"newUser", MaxTodo: 1}, "password")
    if token != "token" {
      t.Errorf("Expect: token = token, receive: token = %s", token)
    }
    if err != nil {
      t.Errorf("Error must be nil, got %v", err)
    }
  })
  t.Run("Signup with empty id", func(t *testing.T) {
    resetUserHandlerNoHash()
    token, err := userHandler.signup(context.Background(), &core.User{ID:""}, "password")
    if token != "" {
      t.Error("Empty id must generate empty token")
    }
    if err != core.ErrEmptyId {
      t.Errorf("Wrong error - expect %v, got %v", core.ErrEmptyId, err)
    }
  })
  t.Run("Signup with empty password", func(t *testing.T) {
    resetUserHandlerNoHash()
    token, err := userHandler.signup(context.Background(), &core.User{ID:"user"}, "")
    if token != "" {
      t.Error("Empty password must generate empty token")
    }
    if err != core.ErrEmptyPassword {
      t.Errorf("Wrong error - expect %v, got %v", core.ErrEmptyPassword, err)
    }
  })
  t.Run("Signup with negative max todo", func(t *testing.T) {
    resetUserHandlerNoHash()
    token, err := userHandler.signup(context.Background(), &core.User{ID:"user", MaxTodo: -1}, "password")
    if token != "" {
      t.Error("Negative max todo must generate empty token")
    }
    if err != core.ErrInvalidMaxTodo {
      t.Errorf("Wrong error - expect %v, got %v", core.ErrInvalidMaxTodo, err)
    }
  })
  t.Run("Signup with zero max todo", func(t *testing.T) {
    resetUserHandlerNoHash()
    token, err := userHandler.signup(context.Background(), &core.User{ID:"user", MaxTodo: 0}, "password")
    if token != "" {
      t.Error("Zero max todo must generate empty token")
    }
    if err != core.ErrInvalidMaxTodo {
      t.Errorf("Wrong error - expect %v, got %v", core.ErrInvalidMaxTodo, err)
    }
  })
  t.Run("Signup with positive max todo", func(t *testing.T) {
    resetUserHandlerNoHash()
    token, err := userHandler.signup(context.Background(), &core.User{ID:"user", MaxTodo: 5}, "password")
    if token != "token" {
      t.Errorf("Expect: token = token, receive: token = %s", token)
    }
    if err != nil {
      t.Errorf("Error must be nil, got %v", err)
    }
  })

  t.Run("Signup then login", func(t *testing.T) {
    resetUserHandlerNoHash()
    // signup
    token, err := userHandler.signup(context.Background(), &core.User{ID:"user", MaxTodo: 5}, "password")
    if token != "token" {
     t.Errorf("Signup - Expect: token = token, receive: token = %s", token)
    }
    if err != nil {
      t.Errorf("Error must be nil, got %v", err)
    }

    // login
    token, err = userHandler.login(context.Background(), "dummy", "password")
    if token != "token" {
      t.Errorf("Login - Expect: token = token, receive: token = %s", token)
    }
    if err != nil {
      t.Errorf("Error must be nil, got %v", err)
    }
  })
}

const bcryptCost = 10
var bcryptPassword, _  = bcrypt.GenerateFromPassword([]byte("password"), bcryptCost)
var bcryptHashUser = core.User{
  ID:   "dummy",
  Hash: string(bcryptPassword),
  MaxTodo: 2,
}

func resetUserHandlerBcrypt() {
  userHandler = UserHandler{
    userRepo: &inMemUserRepo{users: map[string]*core.User{
      bcryptHashUser.ID: &bcryptHashUser,
    }, hash: func(password string) (string, error) {
      hash, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
      return string(hash), err
    }, compare: func(password, hash string) bool {
      return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
    }},
    generateToken: func(user *core.User) (s string, err error) {
      return "token", nil
    },
  }
}

func TestUserHandler_login_bcrypt(t *testing.T) {
  resetUserHandlerBcrypt()

  t.Run("Invalid user", func(t *testing.T) {
    resetUserHandlerBcrypt()
    token, err := userHandler.login(context.Background(), "admin", "password")
    if token != "" {
      t.Error("Invalid user must generate empty token")
    }
    if err != core.ErrWrongIdPassword {
      t.Errorf("Wrong error - expect %v, got %v", core.ErrWrongIdPassword, err)
    }
  })
  t.Run("Wrong password", func(t *testing.T) {
    resetUserHandlerBcrypt()
    token, err := userHandler.login(context.Background(), "dummy", "wrongpassword")
    if token != "" {
      t.Error("Wrong password must generate empty token")
    }
    if err != core.ErrWrongIdPassword {
      t.Errorf("Wrong error - expect %v, got %v", core.ErrWrongIdPassword, err)
    }
  })
  t.Run("Valid user and correct password", func(t *testing.T) {
    resetUserHandlerBcrypt()
    token, err := userHandler.login(context.Background(), "dummy", "password")
    if token != "token" {
      t.Errorf("Expect: token = token, receive: token = %s", token)
    }
    if err != nil {
      t.Errorf("Error must be nil, got %v", err)
    }
  })
}


func TestUserHandler_signup_bcrypt(t *testing.T) {
  resetUserHandlerBcrypt()

  t.Run("Signup with existed user id", func(t *testing.T) {
    resetUserHandlerBcrypt()
    token, err := userHandler.signup(context.Background(), &noHashUser, "password")
    if token != "" {
      t.Error("Existed user must generate empty token")
    }
    if err != core.ErrUserAlreadyExists {
      t.Errorf("Wrong error - expect %v, got %v", core.ErrUserAlreadyExists, err)
    }
  })
  t.Run("Signup with new user id", func(t *testing.T) {
    resetUserHandlerBcrypt()
    token, err := userHandler.signup(context.Background(), &core.User{ID:"newUser", MaxTodo: 1}, "password")
    if token != "token" {
      t.Errorf("Expect: token = token, receive: token = %s", token)
    }
    if err != nil {
      t.Errorf("Error must be nil, got %v", err)
    }
  })
  t.Run("Signup with empty id", func(t *testing.T) {
    resetUserHandlerBcrypt()
    token, err := userHandler.signup(context.Background(), &core.User{ID:""}, "password")
    if token != "" {
      t.Error("Empty id must generate empty token")
    }
    if err != core.ErrEmptyId {
      t.Errorf("Wrong error - expect %v, got %v", core.ErrEmptyId, err)
    }
  })
  t.Run("Signup with empty password", func(t *testing.T) {
    resetUserHandlerBcrypt()
    token, err := userHandler.signup(context.Background(), &core.User{ID:"user"}, "")
    if token != "" {
      t.Error("Empty password must generate empty token")
    }
    if err != core.ErrEmptyPassword {
      t.Errorf("Wrong error - expect %v, got %v", core.ErrEmptyPassword, err)
    }
  })
  t.Run("Signup with negative max todo", func(t *testing.T) {
    resetUserHandlerBcrypt()
    token, err := userHandler.signup(context.Background(), &core.User{ID:"user", MaxTodo: -1}, "password")
    if token != "" {
      t.Error("Negative max todo must generate empty token")
    }
    if err != core.ErrInvalidMaxTodo {
      t.Errorf("Wrong error - expect %v, got %v", core.ErrInvalidMaxTodo, err)
    }
  })
  t.Run("Signup with zero max todo", func(t *testing.T) {
    resetUserHandlerBcrypt()
    token, err := userHandler.signup(context.Background(), &core.User{ID:"user", MaxTodo: 0}, "password")
    if token != "" {
      t.Error("Zero max todo must generate empty token")
    }
    if err != core.ErrInvalidMaxTodo {
      t.Errorf("Wrong error - expect %v, got %v", core.ErrInvalidMaxTodo, err)
    }
  })
  t.Run("Signup with positive max todo", func(t *testing.T) {
    resetUserHandlerBcrypt()
    token, err := userHandler.signup(context.Background(), &core.User{ID:"user", MaxTodo: 5}, "password")
    if token != "token" {
      t.Errorf("Expect: token = token, receive: token = %s", token)
    }
    if err != nil {
      t.Errorf("Error must be nil, got %v", err)
    }
  })

  t.Run("Signup then login", func(t *testing.T) {
    resetUserHandlerBcrypt()
    // signup
    token, err := userHandler.signup(context.Background(), &core.User{ID:"user", MaxTodo: 5}, "password")
    if token != "token" {
      t.Errorf("Signup - Expect: token = token, receive: token = %s", token)
    }
    if err != nil {
      t.Errorf("Error must be nil, got %v", err)
    }

    // login
    token, err = userHandler.login(context.Background(), "dummy", "password")
    if token != "token" {
      t.Errorf("Login - Expect: token = token, receive: token = %s", token)
    }
    if err != nil {
      t.Errorf("Error must be nil, got %v", err)
    }
  })
}

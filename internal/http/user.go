package http

import (
  "errors"
  "github.com/manabie-com/togo/internal/context"
  "github.com/manabie-com/togo/internal/core"
  "net/http"
)

var (
  ErrParseIdPassword       = errors.New("cannot parse user id and password")
  ErrParseUserInfoPassword = errors.New("cannot parse user info and password")
)

type UserHandler struct {
  userRepo core.UserRepo

  parseUserInfoPassword func(r *http.Request) (*core.User, string, error)
  parseIdPassword       func(r *http.Request) (string, string, error)
  generateToken         func(user *core.User) (string, error)

  respondLoginSuccess func(w http.ResponseWriter, r *http.Request, token string)
  respondLoginError   func(w http.ResponseWriter, r *http.Request, err error)

  respondSignupSuccess func(w http.ResponseWriter, r *http.Request, token string)
  respondSignupError   func(w http.ResponseWriter, r *http.Request, err error)
}

func (handler *UserHandler) login(ctx context.Context, id, password string) (string, error) {
  // validate user
  user, err := handler.userRepo.Validate(ctx, id, password)
  if err != nil {
    return "", err
  }
  // generate token
  return handler.generateToken(user)
}

func (handler *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
  // parse id and password
  id, password, err := handler.parseIdPassword(r)
  if err != nil {
    handler.respondLoginError(w, r, err)
    return
  }

  token, err := handler.login(r.Context(), id, password)
  if err != nil {
    handler.respondLoginError(w, r, err)
    return
  }
  handler.respondLoginSuccess(w, r, token)
}

func (handler *UserHandler) signup(ctx context.Context, user *core.User, password string) (string, error) {
  if user.ID == "" {
    return "", core.ErrEmptyId
  }
  if password == "" {
    return "", core.ErrEmptyPassword
  }
  if user.MaxTodo <= 0 {
    return "", core.ErrInvalidMaxTodo
  }
  err := handler.userRepo.Create(ctx, user, password)
  if err != nil {
    return "", err
  }
  return handler.generateToken(user)
}

func (handler *UserHandler) Signup(w http.ResponseWriter, r *http.Request) {
  // parse user & password
  user, password, err := handler.parseUserInfoPassword(r)
  if err != nil {
    handler.respondSignupError(w, r, err)
    return
  }

  token, err := handler.signup(r.Context(), user, password)
  if err != nil {
    handler.respondSignupError(w, r, err)
    return
  }
  handler.respondSignupSuccess(w, r, token)
}

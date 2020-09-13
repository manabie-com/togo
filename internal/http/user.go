package http

import (
  "errors"
  "github.com/manabie-com/togo/internal/context"
  "github.com/manabie-com/togo/internal/core"
  "net/http"
)

var ErrWrongIdPassword = errors.New("user id or password incorrect")

type UserHandler struct {
  userRepo core.UserRepo

  parseIdPassword func(r *http.Request) (id, password string)
  generateToken   func(user *core.User) (string, error)

  respondLoginSuccess func(w http.ResponseWriter, r *http.Request, token string)
  respondLoginError   func(w http.ResponseWriter, r *http.Request, err error)
}

func (handler *UserHandler) login(ctx context.Context, id, password string) (string, error) {
  // validate user
  user, ok := handler.userRepo.Validate(ctx, id, password)
  if !ok {
    return "", ErrWrongIdPassword
  }
  // generate token
  return handler.generateToken(user)
}

func (handler *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
  // parse id and password
  id, password := handler.parseIdPassword(r)

  token, err := handler.login(r.Context(), id, password)
  if err != nil {
    handler.respondLoginError(w, r, err)
    return
  }
  handler.respondLoginSuccess(w, r, token)
}

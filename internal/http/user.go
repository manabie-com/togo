package http

import (
  "errors"
  "github.com/manabie-com/togo/internal/context"
  "github.com/manabie-com/togo/internal/core"
  "log"
  "net/http"
)

var (
  ErrParseIdPassword       = errors.New("cannot parse user id and password")
  ErrParseUserInfoPassword = errors.New("cannot parse user info and password")
)

type userParser interface {
  parseUserInfoPassword(r *http.Request) (*core.User, string, error)
  parseIdPassword(r *http.Request) (string, string, error)
}

type userResponder interface {
  respondLoginSuccess(w http.ResponseWriter, r *http.Request, token string)
  respondLoginError(w http.ResponseWriter, r *http.Request, err error)
  respondSignupSuccess(w http.ResponseWriter, r *http.Request, token string)
  respondSignupError(w http.ResponseWriter, r *http.Request, err error)
  respondLogout(w http.ResponseWriter, r *http.Request)
}

type UserHandler struct {
  parser    userParser
  service   core.UserService
  responder userResponder
}

func (handler *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
  // parse id and password
  id, password, err := handler.parser.parseIdPassword(r)
  if err != nil {
    handler.responder.respondLoginError(w, r, err)
    return
  }

  token, err := handler.service.Login(r.Context(), id, password)
  if err != nil {
    handler.responder.respondLoginError(w, r, err)
    return
  }
  handler.responder.respondLoginSuccess(w, r, token)
}

func (handler *UserHandler) Signup(w http.ResponseWriter, r *http.Request) {
  // parse user & password
  user, password, err := handler.parser.parseUserInfoPassword(r)
  if err != nil {
    handler.responder.respondSignupError(w, r, err)
    return
  }

  token, err := handler.service.Signup(r.Context(), user, password)
  if err != nil {
    handler.responder.respondSignupError(w, r, err)
    return
  }
  handler.responder.respondSignupSuccess(w, r, token)
}

func (handler *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
  user := context.User(r.Context())
  if user == nil {
    w.WriteHeader(http.StatusInternalServerError)
    log.Fatalf("[http::UserHandler::Logout - user not set]\n")
    return
  }

  handler.responder.respondLogout(w, r)
}

type UserHandlerOption func(handler *UserHandler)

func WithUserParser(parser userParser) UserHandlerOption {
  return func(handler *UserHandler) {
    handler.parser = parser
  }
}

func WithUserService(service core.UserService) UserHandlerOption {
  return func(handler *UserHandler) {
    handler.service = service
  }
}

func WithUserResponder(responder userResponder) UserHandlerOption {
  return func(handler *UserHandler) {
    handler.responder = responder
  }
}

func NewUserHandler(opts ...UserHandlerOption) *UserHandler {
  handler := UserHandler{}
  for _, opt := range opts {
    opt(&handler)
  }
  return &handler
}

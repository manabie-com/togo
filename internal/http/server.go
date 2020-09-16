package http

import (
  "github.com/gorilla/mux"
  "github.com/manabie-com/togo/internal/http/middleware"
  "net/http"
)

type Server struct {
  authMw      AuthMw
  userHandler *UserHandler
  taskHandler *TaskHandler
  router      *mux.Router
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  s.router.ServeHTTP(w, r)
}

func (s *Server) routes() {
  s.router.Use(func(handler http.Handler) http.Handler {
    return middleware.Logger{}.MethodAndPath(handler)
  })
  s.router.Use(func(handler http.Handler) http.Handler {
    return middleware.CORS{}.All(handler)
  })
  s.router.HandleFunc("/signup", s.userHandler.Signup).Methods(http.MethodPost)
  s.router.HandleFunc("/login", s.userHandler.Login).Methods(http.MethodPost)
  s.router.Handle("/logout", middleware.ApplyFunc(s.userHandler.Logout, s.authMw.SetUser,
    s.authMw.RequireUser)).Methods(http.MethodPost)

  s.router.Handle("/tasks", middleware.ApplyFunc(s.taskHandler.Index, s.authMw.SetUser,
    s.authMw.RequireUser)).Methods(http.MethodGet)
  s.router.Handle("/tasks", middleware.ApplyFunc(s.taskHandler.Create, s.authMw.SetUser,
    s.authMw.RequireUser)).Methods(http.MethodPost)
  s.router.Handle("/tasks", middleware.ApplyFunc(s.taskHandler.Update, s.authMw.SetUser,
    s.authMw.RequireUser)).Methods(http.MethodPut)
  s.router.Handle("/tasks", middleware.ApplyFunc(s.taskHandler.Delete, s.authMw.SetUser,
    s.authMw.RequireUser)).Methods(http.MethodDelete)
}

type ServerOption func(*Server)

func WithAuthMiddleware(mw AuthMw) ServerOption {
  return func(server *Server) {
    server.authMw = mw
  }
}

func WithUserHandler(handler *UserHandler) ServerOption {
  return func(server *Server) {
    server.userHandler = handler
  }
}

func WithTaskHandler(handler *TaskHandler) ServerOption {
  return func(server *Server) {
    server.taskHandler = handler
  }
}

func NewServer(opts ...ServerOption) *Server {
  s := Server{}
  for _, opt := range opts {
    opt(&s)
  }
  s.router = mux.NewRouter()
  s.routes()
  return &s
}

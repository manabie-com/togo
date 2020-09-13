package http

import (
  "github.com/gorilla/mux"
  "github.com/manabie-com/togo/internal/core"
  "log"
  "net/http"
  "time"
)

type Server struct {
  authMw      AuthMw
  userHandler *UserHandler
  taskHandler *TaskHandler
  router      *mux.Router
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  log.Printf("[http::Server::ServeHTTP - %s %s]\n", r.Method, r.URL.Path)
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Access-Control-Allow-Headers", "*")
  w.Header().Set("Access-Control-Allow-Methods", "*")

  if r.Method == http.MethodOptions {
    w.WriteHeader(http.StatusOK)
    return
  }
  s.router.ServeHTTP(w, r)
}

func (s *Server) routes() {
  s.router.HandleFunc("/login", s.userHandler.Login).Methods(http.MethodGet)

  s.router.Handle("/tasks", ApplyFunc(s.taskHandler.Index, s.authMw.SetUser,
    s.authMw.RequireUser)).Methods(http.MethodGet)
  s.router.Handle("/tasks", ApplyFunc(s.taskHandler.Create, s.authMw.SetUser, s.authMw.RequireUser)).Methods(http.MethodPost)
}

func JSONServer(jwtKey string, userRepo core.UserRepo, taskRepo core.TaskRepo) *Server {
  server := Server{
    authMw: &jsonAuthMw{
      jwtKey: jwtKey,
      userRepo: userRepo,
    },
    userHandler: jsonUserHandler(userRepo, jwtKey, time.Minute*15),
    taskHandler: jsonTaskHandler(taskRepo),
    router: mux.NewRouter(),
  }
  server.routes()
  return &server
}

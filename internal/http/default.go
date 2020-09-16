package http

import (
  "encoding/json"
  "github.com/dgrijalva/jwt-go"
  "github.com/google/uuid"
  "github.com/gorilla/mux"
  "github.com/manabie-com/togo/internal/config"
  "github.com/manabie-com/togo/internal/core"
  "github.com/manabie-com/togo/internal/http/middleware"
  "github.com/manabie-com/togo/internal/service"
  "net/http"
  "time"
)

type jwtManager struct {
  Key        string
  Method     jwt.SigningMethod
  ExpireTime time.Duration
}

func (manager *jwtManager) Token(user *core.User) (string, error) {
  claims := jwt.MapClaims{}
  claims["user_id"] = user.ID
  claims["exp"] = time.Now().Add(manager.ExpireTime).Unix()
  token := jwt.NewWithClaims(manager.Method, claims)
  signed, err := token.SignedString([]byte(manager.Key))
  if err != nil {
    return "", err
  }
  return signed, nil
}

func DefaultServerWithSPA(userRepo core.UserRepo, taskRepo core.TaskRepo) http.Handler {
  apiServer := DefaultServer(userRepo, taskRepo)
  spaServer := &SpaHandler{
    staticPath: "frontend",
    indexPath: "index.html",
  }
  m := mux.NewRouter()
  m.PathPrefix("/api").Handler(http.StripPrefix("/api", apiServer))
  m.PathPrefix("/").Handler(spaServer)
  return m
}

func DefaultServer(userRepo core.UserRepo, taskRepo core.TaskRepo) *Server {
   return NewServer(
    WithAuthMiddleware(&middleware.JWT{
      JwtKey:   config.JwtKey,
      UserRepo: userRepo,
      UnauthorizedRespond: func(w http.ResponseWriter, message string) {
        _ = json.NewEncoder(w).Encode(map[string]string{
          "error": message,
        })
      },
    }),
    WithUserHandler(NewUserHandler(
      WithUserParser(&JsonUserParser{}),
      WithUserService(&service.JwtUserService{
        UserRepo:   userRepo,
        JwtGenerator: &jwtManager{
          Key:        config.JwtKey,
          Method:     jwt.SigningMethodHS256,
          ExpireTime: 15 * time.Minute,
        },
      }),
      WithUserResponder(&JsonUserResponder{}),
    )),
    WithTaskHandler(NewTaskHandler(
      WithTaskParser(&JsonTaskParser{}),
      WithTaskService(&service.BasicTaskService{
        TaskRepo: taskRepo,
        Today: time.Now,
        NewID: func() string {
          return uuid.New().String()
        },
      }),
      WithTaskResponder(&JsonTaskResponder{}),
    )),
  )
}

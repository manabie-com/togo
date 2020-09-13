package http

import (
  "encoding/json"
  "github.com/dgrijalva/jwt-go"
  "github.com/google/uuid"
  "github.com/manabie-com/togo/internal/context"
  "github.com/manabie-com/togo/internal/core"
  "log"
  "net/http"
  "strings"
  "time"
)

const timeLayout = "2006-01-02"

func respondError(w http.ResponseWriter, status int, message string) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(status)
  _ = json.NewEncoder(w).Encode(map[string]string{
    "error": message,
  })
}

func respondUnknownError(w http.ResponseWriter) {
  respondError(w, http.StatusInternalServerError, "Unknown error")
}

func respondData(w http.ResponseWriter, status int, data interface{}) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(status)
  _ = json.NewEncoder(w).Encode(map[string]interface{}{
    "data": data,
  })
}

func jsonUserHandler(userRepo core.UserRepo, jwtKey string, expireTime time.Duration) *UserHandler {
  return &UserHandler{
    userRepo: userRepo,
    parseIdPassword: func(r *http.Request) (id, password string) {
      id = r.FormValue("user_id")
      password = r.FormValue("password")
      return
    },
    generateToken: func(user *core.User) (s string, err error) {
      claims := jwt.MapClaims{}
      claims["user_id"] = user.ID
      claims["exp"] = time.Now().Add(expireTime).Unix()
      token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
      signed, err := token.SignedString([]byte(jwtKey))
      if err != nil {
        return "", err
      }
      return signed, nil
    },
    respondLoginSuccess: func(w http.ResponseWriter, r *http.Request, token string) {
      respondData(w, http.StatusOK, token)
    },
    respondLoginError: func(w http.ResponseWriter, r *http.Request, err error) {
      switch err {
      case ErrWrongIdPassword:
        respondError(w, http.StatusUnauthorized, "User ID or password incorrect")
      default:
        respondUnknownError(w)
      }
    },
  }
}

type jsonTask struct {
  ID          string `json:"id"`
  Content     string `json:"content"`
  CreatedDate string `json:"created_date"`
}

func (jsTask *jsonTask) read(task *core.Task) {
  jsTask.ID = task.ID
  jsTask.Content = task.Content
  jsTask.CreatedDate = task.CreatedDate.Format(timeLayout)
}

func taskSliceToJson(tasks []*core.Task) []*jsonTask {
  out := []*jsonTask{}
  for _, task := range tasks {
    var jsTask jsonTask
    jsTask.read(task)
    out = append(out, &jsTask)
  }
  return out
}

func jsonTaskHandler(taskRepo core.TaskRepo) *TaskHandler {
  return &TaskHandler{
    taskRepo: taskRepo,
    parseTask: func(r *http.Request) (*core.Task, error) {
      var req struct {
        Content string `json:"content"`
      }
      err := json.NewDecoder(r.Body).Decode(&req)
      if err != nil {
        return nil, ErrCannotParseTask
      }
      return &core.Task{Content: req.Content}, nil
    },
    generateTaskId: func() string {
      return uuid.New().String()
    },
    getToday: func() time.Time {
      return time.Now()
    },
    respondCreateSuccess: func(w http.ResponseWriter, r *http.Request, task *core.Task) {
      var jsTask jsonTask
      jsTask.read(task)
      respondData(w, http.StatusOK, jsTask)
    },
    respondCreateError: func(w http.ResponseWriter, r *http.Request, err error) {
      switch err {
      case ErrCannotParseTask:
        respondError(w, http.StatusBadRequest, "cannot parse task")
      case ErrMaxTodoReached:
        respondError(w, http.StatusBadRequest, "maximum number of todo reached")
      default:
        respondUnknownError(w)
      }
    },
    respondIndexSuccess: func(w http.ResponseWriter, r *http.Request, tasks []*core.Task) {
      respondData(w, http.StatusOK, taskSliceToJson(tasks))
    },
    respondIndexError: func(w http.ResponseWriter, r *http.Request, err error) {
      respondUnknownError(w)
    },
  }
}

type jsonAuthMw struct {
  jwtKey   string
  userRepo core.UserRepo
}

func (mw *jsonAuthMw) SetUser(next http.Handler) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    token := strings.TrimSpace(r.Header.Get("Authorization"))
    claims := make(jwt.MapClaims)
    t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
      return []byte(mw.jwtKey), nil
    })
    if err != nil {
      log.Printf("[http::jsonAuthMw::SetUser - parse token error: %v]\n", err)
      respondError(w, http.StatusUnauthorized, "token invalid")
      return
    }
    if !t.Valid {
      log.Printf("[http::jsonAuthMw::SetUser - invalid token: %v]\n", token)
      respondError(w, http.StatusUnauthorized, "token invalid")
      return
    }

    id, ok := claims["user_id"].(string)
    if !ok {
      log.Printf("[http::jsonAuthMw::SetUser - user_id not a string: %v]\n", id)
      respondError(w, http.StatusUnauthorized, "token invalid")
      return
    }

    user, err := mw.userRepo.ById(r.Context(), id)
    if err != nil {
      return
    }
    r = r.WithContext(context.WithUser(r.Context(), user))
    next.ServeHTTP(w, r)
  }
}

func (mw *jsonAuthMw) RequireUser(next http.Handler) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    user := context.User(r.Context())
    if user == nil {
      respondError(w, http.StatusUnauthorized, "unauthorized access")
      return
    }
    next.ServeHTTP(w, r)
  }
}

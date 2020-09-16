package http

import (
  "encoding/json"
  "fmt"
  "github.com/manabie-com/togo/internal/core"
  "net/http"
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

func setCookie(w http.ResponseWriter, name, value string) {
  http.SetCookie(w, &http.Cookie{
    Name:     name,
    Value:    value,
    HttpOnly: true,
  })
}

func removeCookie(w http.ResponseWriter, r *http.Request, name string) {
  http.SetCookie(w, &http.Cookie{
    Name:     name,
    Value:    "",
    Expires:  time.Now().Add(-24 * time.Hour),
    HttpOnly: true,
  })
}

func respondData(w http.ResponseWriter, status int, data interface{}) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(status)
  _ = json.NewEncoder(w).Encode(map[string]interface{}{
    "data": data,
  })
}

type JsonUserParser struct{}

func (parser *JsonUserParser) parseUserInfoPassword(r *http.Request) (*core.User, string, error) {
  temp := struct {
    Id       *string `json:"id"`
    Password *string `json:"password"`
    MaxTodo  *int    `json:"max_todo"`
  }{}
  err := json.NewDecoder(r.Body).Decode(&temp)
  if err != nil || temp.Id == nil || temp.Password == nil || temp.MaxTodo == nil {
    return nil, "", ErrParseUserInfoPassword
  }
  return &core.User{ID: *temp.Id, MaxTodo: *temp.MaxTodo}, *temp.Password, nil
}

func (parser *JsonUserParser) parseIdPassword(r *http.Request) (string, string, error) {
  temp := struct {
    Id       *string `json:"id"`
    Password *string `json:"password"`
  }{}
  err := json.NewDecoder(r.Body).Decode(&temp)
  if err != nil || temp.Id == nil || temp.Password == nil {
    return "", "", ErrParseIdPassword
  }
  return *temp.Id, *temp.Password, nil
}

type JsonUserResponder struct{}

func (responder *JsonUserResponder) respondLoginSuccess(w http.ResponseWriter, r *http.Request, token string) {
  setCookie(w, "Authorization", token)
  respondData(w, http.StatusOK, token)
}

func (responder *JsonUserResponder) respondLoginError(w http.ResponseWriter, r *http.Request, err error) {
  switch err {
  case core.ErrWrongIdPassword, core.ErrEmptyId, core.ErrEmptyPassword:
    respondError(w, http.StatusUnauthorized, "User ID or password incorrect")
  default:
    respondUnknownError(w)
  }
}

func (responder *JsonUserResponder) respondLogout(w http.ResponseWriter, r *http.Request) {
  removeCookie(w, r, "Authorization")
  respondData(w, http.StatusOK, "Logged out")
}

func (responder *JsonUserResponder) respondSignupSuccess(w http.ResponseWriter, r *http.Request, token string) {
  setCookie(w, "Authorization", token)
  respondData(w, http.StatusOK, token)
}

func (responder *JsonUserResponder) respondSignupError(w http.ResponseWriter, r *http.Request, err error) {
  switch err {
  case core.ErrEmptyId, core.ErrEmptyPassword, core.ErrInvalidMaxTodo, ErrParseUserInfoPassword:
    respondError(w, http.StatusBadRequest, "User info invalid")
  case core.ErrUserAlreadyExists:
    respondError(w, http.StatusConflict, "User already exists")
  default:
    respondUnknownError(w)
  }
}

type jsonTask struct {
  ID          string `json:"id"`
  Content     string `json:"content"`
  CreatedDate string `json:"created_date"`
  Done        bool   `json:"done"`
}

func (jsTask *jsonTask) read(task *core.Task) {
  jsTask.ID = task.ID
  jsTask.Content = task.Content
  jsTask.CreatedDate = task.CreatedDate.Format(timeLayout)
  jsTask.Done = task.Done
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

type JsonTaskParser struct{}

func (parser *JsonTaskParser) parseDate(r *http.Request) (time.Time, error) {
  date, err := time.Parse(timeLayout, r.FormValue("created_date"))
  return date, err
}

func (parser *JsonTaskParser) parseNewTask(r *http.Request) (*core.Task, error) {
  var req struct {
    Content string `json:"content"`
  }
  err := json.NewDecoder(r.Body).Decode(&req)
  if err != nil {
    return nil, ErrCannotParseTaskInfo
  }
  return &core.Task{Content: req.Content}, nil
}

func (parser *JsonTaskParser) parseTask(r *http.Request) (*core.Task, error) {
  var req struct {
    Id          string `json:"id"`
    Content     string `json:"content"`
    CreatedDate string `json:"created_date"`
    Done        bool   `json:"done"`
  }
  err := json.NewDecoder(r.Body).Decode(&req)
  if err != nil {
    return nil, ErrCannotParseTaskInfo
  }
  createdDate, err := time.Parse(timeLayout, req.CreatedDate)
  fmt.Println(err)
  if err != nil {
    return nil, ErrCannotParseTaskInfo
  }
  return &core.Task{ID: req.Id, Done: req.Done, Content: req.Content, CreatedDate: createdDate}, nil
}

func (parser *JsonTaskParser) parseTaskId(r *http.Request) (string, error) {
  var req struct {
    Id string `json:"id"`
  }
  err := json.NewDecoder(r.Body).Decode(&req)
  if err != nil {
    return "", ErrCannotParseTaskId
  }
  return req.Id, nil
}

type JsonTaskResponder struct{}

func (responder *JsonTaskResponder) respondCreateSuccess(w http.ResponseWriter, r *http.Request, task *core.Task) {
  var jsTask jsonTask
  jsTask.read(task)
  respondData(w, http.StatusOK, jsTask)
}

func (responder *JsonTaskResponder) respondCreateError(w http.ResponseWriter, r *http.Request, err error) {
  switch err {
  case ErrCannotParseTaskInfo:
    respondError(w, http.StatusBadRequest, "cannot parse task")
  case core.ErrMaxTodoReached:
    respondError(w, http.StatusBadRequest, "maximum number of todo reached")
  default:
    respondUnknownError(w)
  }
}

func (responder *JsonTaskResponder) respondIndexSuccess(w http.ResponseWriter, r *http.Request, tasks []*core.Task) {
  respondData(w, http.StatusOK, taskSliceToJson(tasks))
}

func (responder *JsonTaskResponder) respondIndexError(w http.ResponseWriter, r *http.Request, err error) {
  respondUnknownError(w)
}

func (responder *JsonTaskResponder) respondUpdateSuccess(w http.ResponseWriter, r *http.Request, task *core.Task) {
  var jsTask jsonTask
  jsTask.read(task)
  respondData(w, http.StatusOK, jsTask)
}

func (responder *JsonTaskResponder) respondUpdateError(w http.ResponseWriter, r *http.Request, err error) {
  switch err {
  case ErrCannotParseTaskInfo:
    respondError(w, http.StatusBadRequest, "cannot parse task info")
  default:
    respondUnknownError(w)
  }
}

func (responder *JsonTaskResponder) respondDeleteSuccess(w http.ResponseWriter, r *http.Request, id string) {
  respondData(w, http.StatusOK, id)
}

func (responder *JsonTaskResponder) respondDeleteError(w http.ResponseWriter, r *http.Request, err error) {
  switch err {
  case ErrCannotParseTaskId:
    respondError(w, http.StatusBadRequest, "cannot parse task id")
  default:
    respondUnknownError(w)
  }
}

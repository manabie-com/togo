package http

import (
  "errors"
  "github.com/manabie-com/togo/internal/context"
  "github.com/manabie-com/togo/internal/core"
  "log"
  "net/http"
  "time"
)

var (
  ErrCannotParseTask = errors.New("cannot parse task")
  ErrMaxTodoReached = errors.New("user's maximum number of todo today reached")
)

type TaskHandler struct {
  taskRepo core.TaskRepo

  parseTask      func(r *http.Request) (*core.Task, error)
  generateTaskId func() string
  getToday       func() time.Time

  respondCreateSuccess func(w http.ResponseWriter, r *http.Request, task *core.Task)
  respondCreateError   func(w http.ResponseWriter, r *http.Request, err error)

  respondIndexSuccess func(w http.ResponseWriter, r *http.Request, tasks []*core.Task)
  respondIndexError   func(w http.ResponseWriter, r *http.Request, err error)
}

func (handler *TaskHandler) Index(w http.ResponseWriter, r *http.Request) {
  user := context.User(r.Context())
  if user == nil {
    w.WriteHeader(http.StatusInternalServerError)
    log.Fatalf("[http::TaskHandler::Index - user not set]\n")
    return
  }

  var tasks []*core.Task
  date, err := time.Parse(timeLayout, r.FormValue("created_date"))
  if err != nil {
    tasks, err = handler.taskRepo.ByUser(r.Context(), user.ID)
  } else {
    tasks, err = handler.taskRepo.ByUserDate(r.Context(), user.ID, date)
  }
  if err != nil {
    handler.respondIndexError(w, r, err)
    return
  }

  // success
  handler.respondIndexSuccess(w, r, tasks)
}

func (handler *TaskHandler) create(ctx context.Context, user *core.User, task *core.Task) error {
  // check number of tasks created today
  tasks, err := handler.taskRepo.ByUserDate(ctx, user.ID, handler.getToday())
  if err != nil {
    return err
  }

  // check if max todos reached
  if len(tasks) >= user.MaxTodo {
    return ErrMaxTodoReached
  }

  // add task
  task.UserID = user.ID
  task.CreatedDate = handler.getToday()
  return handler.taskRepo.Create(ctx, task)
}

func (handler *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
  user := context.User(r.Context())
  if user == nil {
    w.WriteHeader(http.StatusInternalServerError)
    log.Fatalf("http::TaskHandler::Create - user not set")
    return
  }

  // parse task
  task, err := handler.parseTask(r)
  if err != nil {
    log.Printf("[http::TaskHandler::Create - parse task error: %v]\n", err)
    handler.respondCreateError(w, r, err)
    return
  }
  task.ID = handler.generateTaskId()

  err = handler.create(r.Context(), user, task)
  if err != nil {
    handler.respondCreateError(w, r, err)
    return
  }

  // success
  handler.respondCreateSuccess(w, r, task)
}

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
  ErrCannotParseTaskInfo = errors.New("cannot parse task info")
  ErrCannotParseTaskId   = errors.New("cannot parse task id")
)

type taskParser interface {
  parseDate(r *http.Request) (time.Time, error)
  parseNewTask(r *http.Request) (*core.Task, error)
  parseTask(r *http.Request) (*core.Task, error)
  parseTaskId(r *http.Request) (string, error)
}

type taskResponder interface {
  respondCreateSuccess(w http.ResponseWriter, r *http.Request, task *core.Task)
  respondCreateError(w http.ResponseWriter, r *http.Request, err error)

  respondIndexSuccess(w http.ResponseWriter, r *http.Request, tasks []*core.Task)
  respondIndexError(w http.ResponseWriter, r *http.Request, err error)

  respondUpdateSuccess(w http.ResponseWriter, r *http.Request, tasks *core.Task)
  respondUpdateError(w http.ResponseWriter, r *http.Request, err error)

  respondDeleteSuccess(w http.ResponseWriter, r *http.Request, id string)
  respondDeleteError(w http.ResponseWriter, r *http.Request, err error)
}

type TaskHandler struct {
  parser    taskParser
  service   core.TaskService
  responder taskResponder
}

func (handler *TaskHandler) Index(w http.ResponseWriter, r *http.Request) {
  user := context.User(r.Context())
  if user == nil {
    w.WriteHeader(http.StatusInternalServerError)
    log.Fatalf("[http::TaskHandler::Index - user not set]\n")
    return
  }

  date, err := handler.parser.parseDate(r)
  var tasks []*core.Task
  if err != nil {
    tasks, err = handler.service.IndexAll(r.Context(), user)
  } else {
    tasks, err = handler.service.IndexDate(r.Context(), user, date)
  }

  if err != nil {
    handler.responder.respondIndexError(w, r, err)
    return
  }
  handler.responder.respondIndexSuccess(w, r, tasks)
}

func (handler *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
  user := context.User(r.Context())
  if user == nil {
    w.WriteHeader(http.StatusInternalServerError)
    log.Fatalf("http::TaskHandler::Create - user not set")
    return
  }

  // parse task
  task, err := handler.parser.parseNewTask(r)
  if err != nil {
    log.Printf("[http::TaskHandler::Create - parse task error: %v]\n", err)
    handler.responder.respondCreateError(w, r, err)
    return
  }

  err = handler.service.Create(r.Context(), user, task)

  if err != nil {
    handler.responder.respondCreateError(w, r, err)
    return
  }
  handler.responder.respondCreateSuccess(w, r, task)
}

func (handler *TaskHandler) Update(w http.ResponseWriter, r *http.Request) {
  user := context.User(r.Context())
  if user == nil {
    w.WriteHeader(http.StatusInternalServerError)
    log.Fatalf("http::TaskHandler::Create - user not set")
    return
  }

  // parse task
  task, err := handler.parser.parseTask(r)
  if err != nil {
    log.Printf("[http::TaskHandler::Update - parse task error: %v]\n", err)
    handler.responder.respondUpdateError(w, r, err)
    return
  }

  err = handler.service.Update(r.Context(), user, task)

  if err != nil {
    handler.responder.respondUpdateError(w, r, err)
    return
  }
  handler.responder.respondUpdateSuccess(w, r, task)
}

func (handler *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
  user := context.User(r.Context())
  if user == nil {
    w.WriteHeader(http.StatusInternalServerError)
    log.Fatalf("http::TaskHandler::Create - user not set")
    return
  }

  // parse task
  id, err := handler.parser.parseTaskId(r)
  if err != nil {
    log.Printf("[http::TaskHandler::Delete - parse task id error: %v]\n", err)
    handler.responder.respondDeleteError(w, r, err)
    return
  }

  err = handler.service.Delete(r.Context(), user, id)

  if err != nil {
    handler.responder.respondDeleteError(w, r, err)
    return
  }
  handler.responder.respondDeleteSuccess(w, r, id)
}

type TaskHandlerOption func(handler *TaskHandler)

func WithTaskParser(parser taskParser) TaskHandlerOption {
  return func(handler *TaskHandler) {
    handler.parser = parser
  }
}

func WithTaskService(service core.TaskService) TaskHandlerOption {
  return func(handler *TaskHandler) {
    handler.service = service
  }
}

func WithTaskResponder(responder taskResponder) TaskHandlerOption {
  return func(handler *TaskHandler) {
    handler.responder = responder
  }
}

func NewTaskHandler(opts ...TaskHandlerOption) *TaskHandler {
  handler := TaskHandler{}
  for _, opt := range opts {
    opt(&handler)
  }
  return &handler
}

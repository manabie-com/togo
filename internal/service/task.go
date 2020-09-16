package service

import (
  "context"
  "github.com/manabie-com/togo/internal/core"
  "sync"
  "time"
)


type BasicTaskService struct {
  TaskRepo core.TaskRepo
  Today    func() time.Time
  NewID    func() string
  mtx      sync.Mutex
}

func (service *BasicTaskService) IndexAll(ctx context.Context, user *core.User) ([]*core.Task, error) {
  tasks, err := service.TaskRepo.ByUser(ctx, user.ID)
  return tasks, err
}

func (service *BasicTaskService) IndexDate(ctx context.Context, user *core.User, date time.Time) ([]*core.Task, error) {
  tasks, err := service.TaskRepo.ByUserDate(ctx, user.ID, date)
  return tasks, err
}

func (service *BasicTaskService) Create(ctx context.Context, user *core.User, task *core.Task) error {
  // lock mutex for concurrent access
  service.mtx.Lock()
  defer service.mtx.Unlock()

  //check number of tasks created today
  tasks, err := service.TaskRepo.ByUserDate(ctx, user.ID, service.Today())
  if err != nil {
    return err
  }

  // check if max todos reached
  if len(tasks) >= user.MaxTodo {
    return core.ErrMaxTodoReached
  }

  // add task
  task.ID = service.NewID()
  task.UserID = user.ID
  task.CreatedDate = service.Today()
  task.Done = false
  task.Deleted = false
  return service.TaskRepo.Create(ctx, task)
}

func (service *BasicTaskService) Update(ctx context.Context, user *core.User, task *core.Task) error {
  return service.TaskRepo.Update(ctx, user, task)
}

func (service *BasicTaskService) Delete(ctx context.Context, user *core.User, id string) error {
  return service.TaskRepo.Delete(ctx, user, id)
}

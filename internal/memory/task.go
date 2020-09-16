package memory

import (
  "context"
  "github.com/manabie-com/togo/internal/core"
  "time"
)

type TaskRepo struct {
  tasks map[string]map[time.Time][]*core.Task
}

func NewTaskRepo() *TaskRepo {
  return &TaskRepo{tasks: make(map[string]map[time.Time][]*core.Task)}
}

func (repo *TaskRepo) Create(ctx context.Context, task *core.Task) error {
  if _, ok := repo.tasks[task.UserID]; !ok {
    repo.tasks[task.UserID] = make(map[time.Time][]*core.Task)
  }
  repo.tasks[task.UserID][task.CreatedDate] = append(repo.tasks[task.UserID][task.CreatedDate], task)
  return nil
}

func (repo *TaskRepo) ByUser(ctx context.Context, userId string) ([]*core.Task, error) {
  if _, ok := repo.tasks[userId]; !ok {
    return []*core.Task{}, nil
  }
  out := []*core.Task{}
  for _, dateList := range repo.tasks[userId] {
    out = append(out, dateList...)
  }
  return out, nil
}

func (repo *TaskRepo) ByUserDate(ctx context.Context, userId string, date time.Time) ([]*core.Task, error) {
  if _, ok := repo.tasks[userId]; !ok {
    return []*core.Task{}, nil
  }
  return repo.tasks[userId][date], nil
}

func (repo *TaskRepo) Update(ctx context.Context, user *core.User, task *core.Task) error {
  if _, ok := repo.tasks[user.ID]; !ok {
    return nil
  }
  for date, dateList := range repo.tasks[user.ID] {
    for i, t := range dateList {
      if t.ID == task.ID {
        repo.tasks[user.ID][date][i] = t
      }
    }
  }
  return nil
}

func (repo *TaskRepo) Delete(ctx context.Context, user *core.User, id string) error {
  if _, ok := repo.tasks[user.ID]; !ok {
    return nil
  }
  for date, dateList := range repo.tasks[user.ID] {
    for i, t := range dateList {
      if t.ID == id {
        repo.tasks[user.ID][date] = append(repo.tasks[user.ID][date][:i], repo.tasks[user.ID][date][(i+1):]...)
      }
    }
  }
  return nil
}

package http

import (
  "context"
  "fmt"
  "github.com/manabie-com/togo/internal/core"
  "testing"
  "time"
)

type dummyTaskRepo struct {
  tasks map[string]map[time.Time][]*core.Task
}

var today, _ = time.Parse(timeLayout, "2020-09-13")

func (repo *dummyTaskRepo) Create(ctx context.Context, task *core.Task) error {
  if _, ok := repo.tasks[task.UserID]; !ok {
    repo.tasks[task.UserID] = make(map[time.Time][]*core.Task)
  }
  repo.tasks[task.UserID][task.CreatedDate] = append(repo.tasks[task.UserID][task.CreatedDate], task)
  return nil
}

func (repo *dummyTaskRepo) ByUser(ctx context.Context, userId string) ([]*core.Task, error) {
  if _, ok := repo.tasks[userId]; !ok {
    return []*core.Task{}, nil
  }
  out := []*core.Task{}
  for _, dateList := range repo.tasks[userId] {
    out = append(out, dateList...)
  }
  return out, nil
}

func (repo *dummyTaskRepo) ByUserDate(ctx context.Context, userId string, date time.Time) ([]*core.Task, error) {
  if _, ok := repo.tasks[userId]; !ok {
    return []*core.Task{}, nil
  }
  return repo.tasks[userId][date], nil
}

var taskHandler TaskHandler

func reset() {
  taskHandler = TaskHandler{
    taskRepo: &dummyTaskRepo{make(map[string]map[time.Time][]*core.Task)},
    getToday: func() time.Time {
      return today
    },
  }
}

func TestTaskHandler_create(t *testing.T) {
  t.Run("Add task", func(t *testing.T) {
    reset()
    user := core.User{ID: "dummy", MaxTodo: 1}
    err := taskHandler.create(context.Background(), &user, &core.Task{Content: ""})
    if err != nil {
      t.Errorf("there should be no error, receive %v", err)
    }
    tasks, err := taskHandler.taskRepo.ByUser(context.Background(), user.ID)
    if err != nil {
      t.Errorf("there should be no error, receive %v", err)
    }
    if len(tasks) != 1 {
      t.Errorf("number of tasks must be 1, receive %v", len(tasks))
    }
  })

  for n := 2; n < 10; n++ {
    t.Run(fmt.Sprintf("Add less than max todo %d", n), func(t *testing.T) {
      reset()
      addCount := n-1
      user := core.User{ID: "dummy", MaxTodo: n}
      for i := 0; i < addCount; i++ {
        err := taskHandler.create(context.Background(), &user, &core.Task{Content: "", CreatedDate: today})
        if err != nil {
          t.Errorf("there should be no error, receive %v at iter %v", err, i)
        }
      }
      tasks, err := taskHandler.taskRepo.ByUserDate(context.Background(), user.ID, today)
      if err != nil {
        t.Errorf("there should be no error, receive %v", err)
      }
      if len(tasks) != addCount {
        t.Errorf("number of tasks must %v, receive %v", addCount, len(tasks))
      }
    })

    t.Run(fmt.Sprintf("Add equal to max todo %d", n), func(t *testing.T) {
      reset()
      addCount := n
      user := core.User{ID: "dummy", MaxTodo: n}
      for i := 0; i < addCount; i++ {
        err := taskHandler.create(context.Background(), &user, &core.Task{Content: "", CreatedDate: today})
        if err != nil {
          t.Errorf("there should be no error, receive %v at iter %v", err, i)
        }
      }
      tasks, err := taskHandler.taskRepo.ByUserDate(context.Background(), user.ID, today)
      if err != nil {
        t.Errorf("there should be no error, receive %v", err)
      }
      if len(tasks) != addCount {
        t.Errorf("number of tasks must %v, receive %v", addCount, len(tasks))
      }
    })

    t.Run(fmt.Sprintf("Add more than max todo %d", n), func(t *testing.T) {
      reset()
      addCount := n+1
      user := core.User{ID: "dummy", MaxTodo: n}
      for i := 0; i < n; i++ {
        err := taskHandler.create(context.Background(), &user, &core.Task{Content: "", CreatedDate: today})
        if err != nil {
          t.Errorf("there should be no error, receive %v at iter %v", err, i)
        }
      }
      for i := 0; i < addCount-n; i++ {
        err := taskHandler.create(context.Background(), &user, &core.Task{Content: "", CreatedDate: today})
        if err != ErrMaxTodoReached {
          t.Errorf("there should be max todo error, receive %v at iter %v", err, n+1+i)
        }
      }
      tasks, err := taskHandler.taskRepo.ByUserDate(context.Background(), user.ID, today)
      if err != nil {
        t.Errorf("there should be no error, receive %v", err)
      }
      if len(tasks) != n {
        t.Errorf("number of tasks must %v, receive %v", n, len(tasks))
      }
    })
  }
}

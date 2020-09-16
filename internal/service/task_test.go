package service

import (
  "context"
  "fmt"
  "github.com/manabie-com/togo/internal/core"
  "github.com/manabie-com/togo/internal/memory"
  "strconv"
  "testing"
  "time"
)

const timeLayout = "2006-01-02"
var counter = 0
var fixedDay, _ = time.Parse(timeLayout, "2020-09-15")
var basicTaskService BasicTaskService

func resetBasicTaskService() {
  counter = 0
  taskRepo := memory.NewTaskRepo()
  basicTaskService = BasicTaskService{
    TaskRepo: taskRepo,
    Today: func() time.Time {
      return fixedDay
    },
    NewID: func() string {
      counter++
      return strconv.Itoa(counter)
    },
  }
}

func TestTaskHandler_Create(t *testing.T) {
  t.Run("Add task", func(t *testing.T) {
    resetBasicTaskService()
    user := core.User{ID: "dummy", MaxTodo: 1}
    err := basicTaskService.Create(context.Background(), &user, &core.Task{Content: ""})
    if err != nil {
      t.Errorf("there should be no error, receive %v", err)
    }
    tasks, err := basicTaskService.IndexAll(context.Background(), &user)
    if err != nil {
      t.Errorf("there should be no error, receive %v", err)
    }
    if len(tasks) != 1 {
      t.Errorf("number of tasks must be 1, receive %v", len(tasks))
    }
  })

  for n := 2; n < 10; n++ {
    t.Run(fmt.Sprintf("Add less than max todo %d", n), func(t *testing.T) {
      resetBasicTaskService()
      addCount := n-1
      user := core.User{ID: "dummy", MaxTodo: n}
      for i := 0; i < addCount; i++ {
        err := basicTaskService.Create(context.Background(), &user, &core.Task{Content: "", CreatedDate: fixedDay})
        if err != nil {
          t.Errorf("there should be no error, receive %v at iter %v", err, i)
        }
      }
      tasks, err := basicTaskService.IndexDate(context.Background(), &user, fixedDay)
      if err != nil {
        t.Errorf("there should be no error, receive %v", err)
      }
      if len(tasks) != addCount {
        t.Errorf("number of tasks must %v, receive %v", addCount, len(tasks))
      }
    })

    t.Run(fmt.Sprintf("Add equal to max todo %d", n), func(t *testing.T) {
      resetBasicTaskService()
      addCount := n
      user := core.User{ID: "dummy", MaxTodo: n}
      for i := 0; i < addCount; i++ {
        err := basicTaskService.Create(context.Background(), &user, &core.Task{Content: "", CreatedDate: fixedDay})
        if err != nil {
          t.Errorf("there should be no error, receive %v at iter %v", err, i)
        }
      }
      tasks, err := basicTaskService.IndexDate(context.Background(), &user, fixedDay)
      if err != nil {
        t.Errorf("there should be no error, receive %v", err)
      }
      if len(tasks) != addCount {
        t.Errorf("number of tasks must %v, receive %v", addCount, len(tasks))
      }
    })

    t.Run(fmt.Sprintf("Add more than max todo %d", n), func(t *testing.T) {
      resetBasicTaskService()
      addCount := n+1
      user := core.User{ID: "dummy", MaxTodo: n}
      for i := 0; i < n; i++ {
        err := basicTaskService.Create(context.Background(), &user, &core.Task{Content: "", CreatedDate: fixedDay})
        if err != nil {
          t.Errorf("there should be no error, receive %v at iter %v", err, i)
        }
      }
      for i := 0; i < addCount-n; i++ {
        err := basicTaskService.Create(context.Background(), &user, &core.Task{Content: "", CreatedDate: fixedDay})
        if err != core.ErrMaxTodoReached {
          t.Errorf("there should be max todo error, receive %v at iter %v", err, n+1+i)
        }
      }
      tasks, err := basicTaskService.IndexDate(context.Background(), &user, fixedDay)
      if err != nil {
        t.Errorf("there should be no error, receive %v", err)
      }
      if len(tasks) != n {
        t.Errorf("number of tasks must %v, receive %v", n, len(tasks))
      }
    })
  }
}

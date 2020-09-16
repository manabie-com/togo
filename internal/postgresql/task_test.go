package postgresql

import (
  "context"
  "github.com/google/uuid"
  "github.com/manabie-com/togo/internal/core"
  "testing"
  "time"
)

func equalDate(a, b time.Time) bool {
  return a.Year() == b.Year() && a.YearDay() == b.YearDay()
}

func testTaskFields(t *testing.T, task *core.Task, id string, userId string, date time.Time) {
  t.Helper()
  if task.UserID != userId {
    t.Errorf("Expect task of %v, got %v", userId, task.UserID)
  }
  if task.ID != id {
    t.Errorf("Expect id to be %v, got %v", id, task.ID)
  }
  if !equalDate(task.CreatedDate, date) {
    t.Errorf("Expect date to be %v, got %v", date, task.CreatedDate)
  }
}

func TestTaskRepo_Create(t *testing.T) {
  t.Run("Create new task", func(t *testing.T) {
    reset()
    err := taskRepo.Create(context.Background(), &task)
    if err != nil {
      t.Errorf("Expect no error, got (%v)", err)
    }
  })
  t.Run("Create then select", func(t *testing.T) {
    reset()
    task.CreatedDate = time.Now()
    err := taskRepo.Create(context.Background(), &task)
    if err != nil {
      t.Errorf("Expect no error, got (%v)", err)
    }
    tasks, err := taskRepo.ByUser(context.Background(), firstUser.ID)
    if err != nil {
      t.Errorf("Expect no error, got (%v)", err)
    }
    if len(tasks) != 1 {
      t.Errorf("Expect %v tasks, got %v tasks", 1, len(tasks))
    }
    testTaskFields(t, tasks[0], task.ID, firstUser.ID, task.CreatedDate)
  })
  t.Run("Create 2 days then select", func(t *testing.T) {
    reset()
    // first task
    firstTask := core.Task{
      ID:          uuid.New().String(),
      Content:     "content",
      UserID:      firstUser.ID,
      CreatedDate: time.Now(),
    }
    err := taskRepo.Create(context.Background(), &firstTask)
    if err != nil {
      t.Errorf("Expect no error, got (%v)", err)
    }

    // second task
    secondTask := core.Task{
      ID:          uuid.New().String(),
      Content:     "content",
      UserID:      firstUser.ID,
      CreatedDate: time.Now().Add(time.Hour * -24),
    }
    err = taskRepo.Create(context.Background(), &secondTask)
    if err != nil {
      t.Errorf("Expect no error, got (%v)", err)
    }

    // get all tasks
    tasks, err := taskRepo.ByUser(context.Background(), firstUser.ID)
    if err != nil {
      t.Errorf("Expect no error, got (%v)", err)
    }
    if len(tasks) != 2 {
      t.Errorf("Expect %v tasks, got %v tasks", 2, len(tasks))
    }

    // get task today
    tasks, err = taskRepo.ByUserDate(context.Background(), firstUser.ID, firstTask.CreatedDate)
    if err != nil {
      t.Errorf("Expect no error, got (%v)", err)
    }
    if len(tasks) != 1 {
      t.Errorf("Expect %v tasks, got %v tasks", 1, len(tasks))
    }
    testTaskFields(t, tasks[0], firstTask.ID, firstUser.ID, firstTask.CreatedDate)

    // get tasks yesterday
    tasks, err = taskRepo.ByUserDate(context.Background(), firstUser.ID, secondTask.CreatedDate)
    if err != nil {
      t.Errorf("Expect no error, got (%v)", err)
    }
    if len(tasks) != 1 {
      t.Errorf("Expect %v tasks, got %v tasks", 1, len(tasks))
    }
    testTaskFields(t, tasks[0], secondTask.ID, firstUser.ID, secondTask.CreatedDate)
  })
}

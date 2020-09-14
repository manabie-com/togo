package postgresql

import (
  "database/sql"
  "github.com/manabie-com/togo/internal/context"
  "github.com/manabie-com/togo/internal/core"
  "log"
  "time"
)

const timeLayout = "2006-01-02"

type TaskRepo struct {
  DB *sql.DB
}

func rowsToTasks(rows *sql.Rows) (tasks []*core.Task, err error) {
  for rows.Next() {
    var task core.Task

    err = rows.Scan(&task.ID, &task.Content, &task.UserID, &task.CreatedDate)
    if err != nil {
      log.Printf("[postgresql::TaskRepo::rowsToTasks - failed to scan task %v]\n", err)
      continue
    }

    tasks = append(tasks, &task)
  }

  if err = rows.Err(); err != nil {
    log.Printf("[postgresql::TaskRepo::rowsToTasks - rows error : %v]\n", err)
    return nil, err
  }
  return tasks, nil
}

func (repo *TaskRepo) Create(ctx context.Context, task *core.Task) error {
  _, err := repo.DB.ExecContext(ctx, "insert into tasks(id,content,user_id,created_date) values ($1,$2,$3,$4)", task.ID,
    task.Content, task.UserID, task.CreatedDate.Format(timeLayout))
  if err != nil {
    log.Printf("[postgresql::TaskRepo::Create - insert error : %v]\n", err)
  }
  return err
}

func (repo *TaskRepo) ByUser(ctx context.Context, userId string) ([]*core.Task, error) {
  rows, err := repo.DB.QueryContext(ctx, "select id,content,user_id,created_date from tasks where user_id=$1", userId)
  if err != nil {
    log.Printf("[postgresql::TaskRepo::ByUser - select error : %v]\n", err)
    return nil, err
  }
  defer rows.Close()

  return rowsToTasks(rows)
}

func (repo *TaskRepo) ByUserDate(ctx context.Context, userId string, date time.Time) ([]*core.Task, error) {
  rows, err := repo.DB.QueryContext(ctx, "select id,content,user_id,"+
    "created_date from tasks where user_id=$1 and created_date=$2",
    userId, date.Format(timeLayout))
  if err != nil {
    log.Printf("[postgresql::TaskRepo::ByUserDate - select error : %v]\n", err)
    return nil, err
  }
  defer rows.Close()

  return rowsToTasks(rows)
}

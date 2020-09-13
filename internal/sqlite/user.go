package sqlite

import (
  "database/sql"
  "github.com/manabie-com/togo/internal/context"
  "github.com/manabie-com/togo/internal/core"
  "log"
)


type UserRepo struct {
  DB             *sql.DB
}

func (repo *UserRepo) ById(ctx context.Context, id string) (*core.User, error) {
  var user core.User
  user.ID = id
  row := repo.DB.QueryRowContext(ctx, "select password, max_todo from users where id=?", id)
  err := row.Scan(&user.Hash, &user.MaxTodo)
  if err != nil {
    if err != core.ErrUserNotFound {
      log.Printf("[sqlite::UserRepo::ByUser - row scan error: %v]\n", err)
    }
    return nil, err
  }
  return &user, nil
}

func (repo *UserRepo) Validate(ctx context.Context, userId string, password string) (*core.User, bool) {
  var user core.User
  user.ID = userId
  row := repo.DB.QueryRowContext(ctx, "select password, max_todo from users where id=?", userId)
  err := row.Scan(&user.Hash, &user.MaxTodo)
  if err != nil {
    log.Printf("[sqlite::UserRepo::ValidateUser - row scan error: %v]\n", err)
    return nil, false
  }

  if password != user.Hash {
    return nil, false
  }
  return &user, true
}

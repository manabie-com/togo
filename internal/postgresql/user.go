package postgresql

import (
  "database/sql"
  "github.com/lib/pq"
  "github.com/manabie-com/togo/internal/context"
  "github.com/manabie-com/togo/internal/core"
  "golang.org/x/crypto/bcrypt"
  "log"
)

var pqErrUniqueConstraint = pq.ErrorCode("23505")

type UserRepo struct {
  DB *sql.DB
  Cost int
}

func (repo *UserRepo) Hash(password string) (string, error) {
  hash, err := bcrypt.GenerateFromPassword([]byte(password), repo.Cost)
  return string(hash), err
}

func (repo *UserRepo) Compare(password, hash string) bool {
  return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func (repo *UserRepo) Create(ctx context.Context, user *core.User, password string) (err error) {
  user.Hash, err = repo.Hash(password)
  if err != nil {
    return
  }
  rs, err := repo.DB.ExecContext(ctx, "insert into users(id, hash, max_todo) values ($1,$2,$3);", user.ID, user.Hash,
    user.MaxTodo)
  if err != nil {
    if pqErr, ok := err.(*pq.Error); ok && pqErr.Code != pqErrUniqueConstraint {
      log.Printf("[postgresql::UserRepo::Create - exec error: %v (%v)]\n", err, pqErr.Code)
      return err
    } else if !ok {
      log.Printf("[postgresql::UserRepo::Create - exec error: %v]\n", err)
      return err
    }
    return core.ErrUserAlreadyExists
  }
  if affected, _ := rs.RowsAffected(); affected == 0 {
    return core.ErrUserAlreadyExists
  }
  return err
}

func (repo *UserRepo) ById(ctx context.Context, id string) (*core.User, error) {
  var user core.User
  user.ID = id
  row := repo.DB.QueryRowContext(ctx, "select hash, max_todo from users where id=$1;", id)
  err := row.Scan(&user.Hash, &user.MaxTodo)
  if err != nil {
    if err != sql.ErrNoRows {
      log.Printf("[postgresql::UserRepo::ByUser - row scan error: %v]\n", err)
      return nil, err
    }
    return nil, core.ErrUserNotFound
  }
  return &user, nil
}

func (repo *UserRepo) Validate(ctx context.Context, userId string, password string) (*core.User, error) {
  var user core.User
  user.ID = userId
  row := repo.DB.QueryRowContext(ctx, `select hash, max_todo from users where id = $1;`, userId)
  err := row.Scan(&user.Hash, &user.MaxTodo)
  if err != nil {
    if err != sql.ErrNoRows {
      log.Printf("[postgresql::UserRepo::ValidateUser - row scan error: %v]\n", err)
      return nil, err
    }
    return nil, core.ErrWrongIdPassword
  }

  if !repo.Compare(password, user.Hash) {
    return nil, core.ErrWrongIdPassword
  }
  return &user, nil
}

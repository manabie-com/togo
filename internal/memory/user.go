package memory

import (
  "context"
  "github.com/manabie-com/togo/internal/core"
)

type UserRepo struct{
  hashFunc    func(password string) (string, error)
  compareFunc func(password, hash string) bool
  users       map[string]*core.User
}

func NewUserRepo(hash func(string) (string, error), compare func(string, string) bool) *UserRepo {
  return &UserRepo{
    hashFunc:    hash,
    compareFunc: compare,
    users:       make(map[string]*core.User),
  }
}

func (repo *UserRepo) Hash(password string) (string, error) {
  return repo.hashFunc(password)
}

func (repo *UserRepo) Compare(password, hash string) bool {
  return repo.compareFunc(password, hash)
}

func (repo *UserRepo) Create(ctx context.Context, user *core.User, password string) (err error) {
  if _, ok := repo.users[user.ID]; ok {
    return core.ErrUserAlreadyExists
  }
  user.Hash, err = repo.hashFunc(password)
  repo.users[user.ID] = user
  return nil
}

func (repo *UserRepo) ById(ctx context.Context, id string) (*core.User, error) {
  if _, ok := repo.users[id]; ok {
    return repo.users[id], nil
  }
  return nil, core.ErrUserNotFound
}

func (repo *UserRepo) Validate(ctx context.Context, id string, password string) (*core.User, error) {
  if u, ok := repo.users[id]; ok && repo.compareFunc(password, u.Hash) {
    return u, nil
  }
  return nil, core.ErrWrongIdPassword
}

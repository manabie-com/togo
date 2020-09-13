package context

import (
  "context"
  "github.com/manabie-com/togo/internal/core"
  "log"
)

type contextKey string

const userKey contextKey = "user"

func WithUser(ctx context.Context, user *core.User) context.Context {
  return context.WithValue(ctx, userKey, user)
}

func User(ctx context.Context) *core.User {
  tmp := ctx.Value(userKey)
  if tmp == nil {
    return nil
  }
  user, ok := tmp.(*core.User)
  if !ok {
    log.Fatalf("context - user value set incorrectly. type=%T, value=%#v", tmp, tmp)
    return nil
  }
  return user
}

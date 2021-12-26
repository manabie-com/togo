package combined

import (
	"context"
	"togo-user-session-service/internal/model"
	"togo-user-session-service/internal/storage/redis"
	"togo-user-session-service/internal/storage/sqlite"
)

type CombinedDB struct {
	SessionDB *redis.RedisStorage
	UserDB    *sqlite.SQLiteDB
}

func (c *CombinedDB) RegisterOrLogin(ctx context.Context, username, password string) (string, error) {
	user, err := c.UserDB.RegisterOrLogin(ctx, username, password)
	if err != nil {
		return "", err
	}
	token, err := c.SessionDB.CreateToken(ctx, user)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (c *CombinedDB) VerifyToken(ctx context.Context, token string) (*model.User, error) {
	return c.SessionDB.VerifyToken(ctx, token)
}

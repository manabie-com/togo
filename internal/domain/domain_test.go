package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/shanenoi/togo/config"
	"github.com/shanenoi/togo/internal/storages/models"
	"github.com/shanenoi/togo/internal/storages/postgresql"
	"github.com/shanenoi/togo/internal/storages/redis"
	"math/rand"
	"net/http/httptest"
	"testing"
	"time"
)

func AddConfigToContext(ctx *gin.Context) {
	databaseConfigs := config.DatabaseConfigs()
	dbEngine := postgresql.Connect(databaseConfigs.DB_URI)

	redisConfigs := config.RedisConfigs()
	redisEngine := redis.Connect(
		redisConfigs.REDIS_URI,
		redisConfigs.REDIS_PASS,
		redisConfigs.REDIS_DEFAULT_DB,
	)
	postgresql.MakeMigrations(dbEngine.Database)

	ctx.Set(config.POSTGRESQL_DB, dbEngine)
	ctx.Set(config.REDIS_DB, redisEngine)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestAll(t *testing.T) {
	background, _ := gin.CreateTestContext(httptest.NewRecorder())
	AddConfigToContext(background)

	user := &models.User{Username: RandStringRunes(50), Password: RandStringRunes(50)}
	CaseUser(t, background, user)
	CaseCreateOneTask(t, background, user)
}

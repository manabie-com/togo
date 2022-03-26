package storage

import (
	"context"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"

	"github.com/luongdn/togo/database"
	"github.com/luongdn/togo/models"
)

func TestGetCache(t *testing.T) {
	SQLdb, _ := database.OpenMockDBConn()
	db, _ := SQLdb.DB()
	defer db.Close()

	redisServer := miniredis.RunT(t)
	redisClient := database.NewMockRedisClient(redisServer)
	ruleStore := NewRuleStore(SQLdb, redisClient)

	user_id := "user_id"
	action := models.TaskCreate
	rule, err := ruleStore.GetCache(context.TODO(), user_id, action)

	assert.Equal(t, redis.Nil, err)
	assert.Equal(t, models.Rule{}, rule)
}

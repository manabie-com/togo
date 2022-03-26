package routes

import (
	"context"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/luongdn/togo/database"
	"github.com/luongdn/togo/models"
	"github.com/stretchr/testify/assert"
)

func TestThrottled(t *testing.T) {
	redisServer := miniredis.RunT(t)
	redisClient := database.NewMockRedisClient(redisServer)

	rule := models.Rule{
		ID:              "id",
		UserID:          "user_id",
		Action:          models.TaskCreate,
		Unit:            models.Second,
		RequestsPerUnit: 10,
	}
	counterKey := buildCounterKey(rule)

	assert.Equal(t, false, throttled(context.TODO(), redisClient, rule))
	redisServer.CheckGet(t, counterKey, "1")

	redisServer.FastForward(time.Second)
	if redisServer.Exists(counterKey) {
		t.Fatalf("%s should not have existed anymore", counterKey)
	}

	redisServer.IncrByFloat(counterKey, 10)
	assert.Equal(t, true, throttled(context.TODO(), redisClient, rule))
}

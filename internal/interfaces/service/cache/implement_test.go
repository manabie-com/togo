package cache

import (
	"context"
	"testing"

	"github.com/datshiro/togo-manabie/internal/interfaces/models"
	"github.com/stretchr/testify/assert"
)

const DefaultRedisLocation = "redis://localhost:6379"

func TestIncreaseQuotaSuccess(t *testing.T) {
	cache, err := CreateRedisClient(DefaultRedisLocation)
	assert.NoError(t, err)
	assert.NotEqual(t, nil, cache)

	ctx := context.Background()
	user := &models.User{ID: 1, Name: "Dat", Quota: 5}
	err = cache.IncreaseQuota(ctx, user)
	assert.NoError(t, err)

	valid, err := cache.ValidateQuota(ctx, user)
	assert.NoError(t, err)
	assert.Equal(t, true, valid)
}

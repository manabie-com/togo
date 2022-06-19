package domain

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"togo/pkg/random"
)

func TestNewUser(t *testing.T) {
	id, _ := uuid.NewUUID()
	dailyTaskLimit := random.RandomInt(1, 10)
	user := NewUser(id, dailyTaskLimit)
	assert.Equal(t, id, user.ID)
	assert.Equal(t, dailyTaskLimit, user.DailyTaskLimit)
}

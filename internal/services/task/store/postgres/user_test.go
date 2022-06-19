package postgres

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"togo/internal/services/task/domain"
	"togo/pkg/random"
)

func createUser(t *testing.T) *domain.User {
	id, _ := uuid.NewUUID()
	dailyTaskLimit := random.RandomInt(1, 10)
	user := domain.NewUser(id, dailyTaskLimit)
	err := userRepo.Save(user)
	assert.NoError(t, err)
	return user
}

func TestUserRepository_Save(t *testing.T) {
	t.Run("should save user", func(t *testing.T) {
		createUser(t)
	})
	t.Run("should error if daily task limit is empty", func(t *testing.T) {
		id, _ := uuid.NewUUID()
		user := domain.NewUser(id, 0)
		err := userRepo.Save(user)
		assert.Error(t, err)
	})
}

func TestUserRepository_FindByID(t *testing.T) {
	t.Run("should find user", func(t *testing.T) {
		user := createUser(t)
		foundUser, err := userRepo.FindByID(user.ID)
		assert.NoError(t, err)
		assert.Equal(t, user.ID, foundUser.ID)
		assert.Equal(t, user.DailyTaskLimit, foundUser.DailyTaskLimit)
	})
	t.Run("should return nil if user is not exists", func(t *testing.T) {
		id := uuid.New()
		foundUser, err := userRepo.FindByID(id)
		assert.NoError(t, err)
		assert.Nil(t, foundUser)
	})
}

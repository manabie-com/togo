package services

import (
	"github.com/manabie-com/togo/mocks"
	"github.com/manabie-com/togo/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestGetUserByUserName(t *testing.T) {
	userRepo := new(mocks.IUserRepository)

	user := models.User{
		Username:  "huyha",
		Password:  "123456",
		MaxTodo:   2,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	c := userRepo.On("GetUserByUserName", user.Username).Return(&user, nil)

	require.Equal(t, []*mock.Call{c}, userRepo.ExpectedCalls)

	userService := UserService{userRepo}

	actualUser, err := userService.GetUserByUserName(user.Username)

	if err != nil {
		t.Errorf("expected %q but got %q", &user, actualUser)
	}

	assert.Equal(t, &user, actualUser)
}

package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/shanenoi/togo/internal/storages/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func CaseUser(t *testing.T, background *gin.Context, user *models.User) {
	password := user.Password
	userDomain := NewUserDomain(background)

	token := ""
	user.Password = password
	err := userDomain.SignupUser(user)
	assert.Equal(t, err, nil)

	user.Password = password
	err = userDomain.SignupUser(user)
	assert.NotEqual(t, err, nil)

	user.Password = password
	token, err = userDomain.LoginUser(user)
	assert.Equal(t, err, nil)
	assert.NotEqual(t, token, "")
}

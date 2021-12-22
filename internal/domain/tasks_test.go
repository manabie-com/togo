package domain

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/shanenoi/togo/config"
	"github.com/shanenoi/togo/internal/storages/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func CaseCreateOneTask(t *testing.T, background *gin.Context, user *models.User) {
	userDomain := NewUserDomain(background)
	taskDomain := NewTaskDomain(background)

	_, err := userDomain.LoginUser(user)
	assert.Equal(t, err, nil)

	data, _ := json.Marshal(map[string]string{"a": "a", "b": "b"})

	var mess string
	for index := 0; index < config.MAX_TASK_PER_DAY; index++ {
		mess, _ = taskDomain.CreateOneTask(float64(user.ID), data)
		assert.NotEqual(t, mess, config.RESP_OUT_OF_SLOT)
	}
	mess, _ = taskDomain.CreateOneTask(float64(user.ID), data)
	assert.Equal(t, mess, config.RESP_OUT_OF_SLOT)
}

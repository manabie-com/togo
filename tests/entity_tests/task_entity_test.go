package entity_tests

import (
	"github.com/go-playground/assert/v2"
	"github.com/google/uuid"
	"log"
	"manabie-com/togo/entity"
	"manabie-com/togo/util"
	"testing"
	"time"
)

func TestSaveTask(t *testing.T) {

	err := refreshUserAndTaskTable()
	if err != nil {
		log.Fatalf("Error user and task refreshing table %v\n", err)
	}

	user, err := seedOneUser()
	if err != nil {
		log.Fatalf("Cannot seed user %v\n", err)
	}

	newTask := entity.Task{
		ID:          uuid.New().String(),
		Content:     "content",
		UserID:      user.ID,
		CreatedDate: time.Now().Format(util.DefaultTimeFormat),
	}
	err = newTask.Create()
	if err != nil {
		t.Errorf("this is the error getting the post: %v\n", err)
		return
	}

	assert.Equal(t, newTask.Content, "content")
	assert.Equal(t, newTask.UserID, user.ID)
}
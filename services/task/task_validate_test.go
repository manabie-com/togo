package task

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/manabie-com/backend/entity"
	"github.com/manabie-com/backend/utils"

	"github.com/stretchr/testify/assert"

	mockTakValicateService "github.com/manabie-com/backend/mocks/taskservicevalidate"
)

func callMockValidate(serv I_TaskServiceValidate, task entity.Task) *utils.ErrorRest {
	err := serv.Validate(&task)
	return err
}
func TestTaskServiceValidate(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	task := CreateTask()

	taskServiceValidateMock := mockTakValicateService.NewMockI_TaskServiceValidate(ctl)
	taskServiceValidateMock.EXPECT().Validate(&task).Return(nil)

	validate := callMockValidate(taskServiceValidateMock, task)

	assert.Nil(t, validate)
}

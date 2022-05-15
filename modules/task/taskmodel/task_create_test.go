package taskmodel_test

import (
	"github.com/japananh/togo/modules/task/taskmodel"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTaskCreate_Validate(t *testing.T) {
	t.Run("validate success", func(t *testing.T) {
		task := taskmodel.TaskCreate{}
		task.Title = "task name"
		task.Description = "description"
		task.FakeCreatedBy = "3mHP8w3u35tRZa"
		task.FakeAssigneeId = "3mHP8w3u35tRZa"
		task.FakeParentId = "3zRzKxgk8Vayfn"

		err := task.Validate()
		require.Nil(t, err, err)
	})

	t.Run("validate err", func(t *testing.T) {
		task := taskmodel.TaskCreate{}
		task.Title = "task name"
		task.Description = "description"
		task.FakeCreatedBy = "3mHP8w3"
		task.FakeAssigneeId = "3mHP8w3u35tRZa"
		task.FakeParentId = "3zRzKxgk8Vayfn"

		err := task.Validate()
		require.NotNil(t, err)
	})
}

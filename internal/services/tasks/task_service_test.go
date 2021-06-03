package tasks

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/manabie-com/togo/internal/models"
	"github.com/manabie-com/togo/internal/utils/random"
	timeUtils "github.com/manabie-com/togo/internal/utils/time"
	"github.com/stretchr/testify/assert"
)

func TestTaskServiceTaskCount(t *testing.T) {
	//TODO: implement unit test for task count

}

func TestTaskServiceGetTasksCreatedOn(t *testing.T) {
	//TODO: implement unit test for getting task list

}

var TaskServiceCreateNewTestCase = map[string]struct {
	inputModel  *models.Task
	expectedErr error
}{
	"Should create task successfully": {
		inputModel: &models.Task{
			Content:    random.RandString(10),
			UserID:     random.RandString(10),
			CreateDate: timeUtils.CurrentDate(),
		},
		expectedErr: nil,
	},
}

func TestTaskServiceCreateNew(t *testing.T) {
	t.Parallel()

	for caseName, tCase := range TaskServiceCreateNewTestCase {
		t.Run(caseName, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			model := tCase.inputModel
			mock.ExpectQuery(`INSERT INTO "manabie"."tasks" (.+)`).
				WithArgs(model.Content, model.UserID, model.CreateDate).
				WillReturnRows(sqlmock.NewRows([]string{
					models.TaskColumns.ID,
				}).AddRow(1))
			defer db.Close()
			s := NewService(db)
			got := s.CreateNew(context.Background(), model)
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
			assert.Equal(t, tCase.expectedErr, got)
		})
	}
}

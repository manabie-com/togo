package postgres

import (
	"context"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bxcodec/faker/v3"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/stretchr/testify/assert"
)

func Test_Task__RetrieveTasksFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	req := &storages.Task{
		UserID:      faker.Username(),
		CreatedDate: "2021-07-31",
	}
	rErr := fmt.Errorf("sql error")

	mock.ExpectQuery(regexp.QuoteMeta(RetrieveTasksStmt)).WillReturnError(rErr)

	taskStore := NewTaskStore(db)
	result, err := taskStore.RetrieveTasks(context.Background(), req)

	assert.Equal(t, rErr, err)
	assert.Nil(t, result)
}

func Test_Task_RetrieveTasksSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	userID := faker.Username()
	req := &storages.Task{
		UserID:      userID,
		CreatedDate: "2021-07-31",
	}

	ans := []*storages.Task{
		{
			UserID:       userID,
			ID:           "id1",
			Content:      "something content",
			CreatedDate:  "2021-07-31",
			NumberInDate: 1,
		},
		{
			UserID:       userID,
			ID:           "id2",
			Content:      "something content",
			CreatedDate:  "2021-07-31",
			NumberInDate: 2,
		},
		{
			UserID:       userID,
			ID:           "id3",
			Content:      "something content",
			CreatedDate:  "2021-07-31",
			NumberInDate: 3,
		},
	}

	rows := sqlmock.NewRows([]string{"id", "content", "user_id", "created_date", "number_in_date"})

	for _, ta := range ans {
		rows = rows.AddRow(ta.ID, ta.Content, ta.UserID, ta.CreatedDate, ta.NumberInDate)
	}
	mock.ExpectQuery(regexp.QuoteMeta(RetrieveTasksStmt)).WillReturnRows(rows)

	taskStore := NewTaskStore(db)
	result, err := taskStore.RetrieveTasks(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, ans, result)
}

func Test_Task_AddTaskFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	req := &storages.Task{
		UserID:       faker.Username(),
		Content:      "something content",
		CreatedDate:  "2021-07-31",
		NumberInDate: 3,
	}
	rErr := fmt.Errorf("sql error")

	mock.ExpectExec(regexp.QuoteMeta(AddTaskStmt)).WillReturnError(rErr)

	taskStore := NewTaskStore(db)
	err = taskStore.AddTask(context.Background(), req)

	assert.Equal(t, rErr, err)
}

func Test_Task_AddTaskSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	req := &storages.Task{
		UserID:       faker.Username(),
		Content:      "something content",
		CreatedDate:  "2021-07-31",
		NumberInDate: 3,
	}

	mock.ExpectExec(regexp.QuoteMeta(AddTaskStmt)).WillReturnResult(sqlmock.NewResult(1, 1))

	taskStore := NewTaskStore(db)
	err = taskStore.AddTask(context.Background(), req)

	assert.NoError(t, err)
}

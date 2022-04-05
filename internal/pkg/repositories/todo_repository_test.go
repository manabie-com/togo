package repositories

import (
	"context"
	"regexp"
	"testing"
	"time"
	"togo/internal/pkg/mocks"
	"togo/pkg/utils"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var ctx = context.Background()

func TestUserRepoCountTodosByDay(t *testing.T) {
	userID := 1
	now := time.Now()
	startOfDay := utils.StartOfDay(now)
	endOfDay := utils.EndOfDay(now)
	dbMock, mock := mocks.NewDatabaseMock()
	row := sqlmock.NewRows([]string{"count"}).AddRow(1)
	repo := NewToDoRepository(mock)
	query := regexp.QuoteMeta(`SELECT * FROM "todos" WHERE user_id = $1 and created_at > $2 and created_at < $3`)
	dbMock.ExpectQuery(query).WithArgs(userID, startOfDay, endOfDay).WillReturnRows(row)
	count, err := repo.CountTodosByDay(ctx, userID)

	assert.NoError(t, err)
	assert.Equal(t, count, userID)
}

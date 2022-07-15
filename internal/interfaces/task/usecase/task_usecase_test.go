package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/datshiro/togo-manabie/internal/interfaces/domain"
	"github.com/datshiro/togo-manabie/internal/interfaces/models"
	"github.com/datshiro/togo-manabie/internal/interfaces/service/cache"
	mock_service "github.com/datshiro/togo-manabie/internal/interfaces/service/cache/mocks"
	"github.com/stretchr/testify/suite"
	"github.com/volatiletech/null/v8"
)

var (
	mockUser = &models.User{
		ID:        1,
		Email:     null.StringFrom("datshiro@gmail.com"),
		Password:  "",
		Name:      "Dat",
		Quota:     5,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: null.TimeFrom(time.Time{}),
	}

	mockTask = &models.Task{
		ID:          1,
		Title:       "Task title",
		Description: null.StringFrom("Task Description"),
		Priority:    1,
		IsDone:      false,
		UserID:      mockUser.ID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   null.TimeFrom(time.Time{}),
	}
)

type TaskUseCaseTestSuite struct {
	suite.Suite
	taskUC       domain.TaskUseCase
	CacheService cache.CacheService
	mock         sqlmock.Sqlmock
	db           *sql.DB
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestEventService(t *testing.T) {
	suite.Run(t, new(TaskUseCaseTestSuite))
}

func (s *TaskUseCaseTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	// Inject mock instance into boil.
	// boil.SetDB(db)

	s.NoError(err)
	s.mock = mock
	s.db = db
}

func (s *TaskUseCaseTestSuite) TestCreateSuccess() {
	cacheService := mock_service.NewCacheService(s.T())
	ctx := context.Background()

	cacheService.On("IncreaseQuota", ctx, mockUser).Return(nil)
	cacheService.On("ValidateQuota", ctx, mockUser).Return(true, nil)
	s.mock.ExpectBegin()

	rows := sqlmock.NewRows([]string{"id", "is_done"}).
		AddRow(mockTask.ID, false)

	s.mock.ExpectQuery("INSERT INTO \"togo\".\"tasks\" (.+)").
		WithArgs(
			mockTask.Title,
			mockTask.Description,
			mockTask.Priority,
			mockTask.UserID,
			mockTask.CreatedAt,
			mockTask.UpdatedAt,
			mockTask.DeletedAt,
		).WillReturnRows(rows)
	// . WillReturnResult(sqlmock.NewResult(int64(mockTask.ID), 1))
	s.mock.ExpectCommit()

	s.taskUC = NewTaskUseCase(s.db, cacheService)
	createdTask := &models.Task{
		Title:       mockTask.Title,
		Description: mockTask.Description,
		Priority:    mockTask.Priority,
		IsDone:      mockTask.IsDone,
		UserID:      mockTask.UserID,
		UpdatedAt:   mockTask.UpdatedAt,
		CreatedAt:   mockTask.CreatedAt,
		DeletedAt:   mockTask.DeletedAt,
	}
	err := s.taskUC.CreateTask(ctx, createdTask, mockUser)
	if err != nil {
		panic(err)
	}
	s.NoError(err)
	s.Equal(createdTask, mockTask)

	// ensure all expectations have been met
	if err := s.mock.ExpectationsWereMet(); err != nil {
		fmt.Printf("unmet expectation error: %s", err)
	}

}

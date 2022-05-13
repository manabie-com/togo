package task

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	r "github.com/nvhai245/togo/internal/domain"
	"github.com/stretchr/testify/assert"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestCountTaskByUserID(t *testing.T) {
	type testCountTasks struct {
		testCase      string
		userID        int64
		limit         int
		tasks         []r.Task
		expectedError error
	}
	var testData = []testCountTasks{
		{
			testCase: "success",
			userID:   3,
			limit:    3,
			tasks: []r.Task{
				{
					ID:      1,
					UserID:  1,
					Content: "Do homework",
					Timestamp: r.Timestamp{
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
					},
				},
				{
					ID:      2,
					UserID:  1,
					Content: "Cleaning",
					Timestamp: r.Timestamp{
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
					},
				},
			},
			expectedError: nil,
		},
		{
			userID:        4,
			testCase:      "user does not exist",
			tasks:         []r.Task{},
			expectedError: errors.New("user does not exist"),
		},
		{
			userID:   1,
			limit:    1,
			testCase: "task limit exceeded",
			tasks: []r.Task{
				{
					ID:      1,
					UserID:  3,
					Content: "Do homework",
					Timestamp: r.Timestamp{
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
					},
				},
			},
			expectedError: errors.New(fmt.Sprintf("exceeded daily %v tasks limit", 1)),
		},
	}

	for _, tc := range testData {
		t.Run(tc.testCase, func(t *testing.T) {
			db, mock := NewMock()
			repo := &repository{db}
			defer func() {
				repo.Close()
			}()

			query := "SELECT daily_limit FROM users WHERE id = ?"
			rows := sqlmock.NewRows([]string{"daily_limit"}).AddRow(tc.limit)
			if tc.limit == 0 {
				mock.ExpectQuery(query).WillReturnError(errors.New("user does not exist"))
			} else {
				mock.ExpectQuery(query).WillReturnRows(rows)
			}

			query = "SELECT COUNT(*) FROM tasks WHERE user_id = ? AND (created_at BETWEEN Date('now') AND Date('now','+1 day','-0.001 second'))"
			rows = sqlmock.NewRows([]string{"count"}).AddRow(len(tc.tasks))
			mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

			err := repo.CheckTaskByUserID(context.Background(), tc.userID)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestCreateTask(t *testing.T) {
	type testCreateTask struct {
		testCase string
		userID   int64
		task     string
		expected *r.Task
	}

	testData := []testCreateTask{
		{
			testCase: "create success",
			userID:   1,
			task:     "Do homework",
			expected: &r.Task{
				ID:      1,
				UserID:  1,
				Content: "Do homework",
			},
		},
		{
			testCase: "create error",
			userID:   1,
			expected: nil,
		},
	}

	for _, tc := range testData {
		t.Run(tc.testCase, func(t *testing.T) {
			db, mock := NewMock()
			repo := &repository{db}
			defer func() {
				repo.Close()
			}()

			query := "INSERT INTO tasks (user_id, content) VALUES (?, ?) RETURNING *"
			rows := sqlmock.NewRows([]string{"id", "user_id", "content", "created_at", "updated_at"})
			if tc.expected != nil {
				rows.AddRow(tc.expected.ID, tc.expected.UserID, tc.expected.Content, time.Now(), time.Now())
			}
			mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

			task, err := repo.CreateTask(context.Background(), tc.userID, tc.task)
			if tc.expected != nil {
				assert.Equal(t, tc.expected.ID, task.ID)
				assert.NoError(t, err)
			} else {
				assert.Equal(t, tc.expected, task)
				assert.Error(t, err)
			}
		})
	}

}

func TestGetAllTaskByUserID(t *testing.T) {
	type testCreateTask struct {
		testCase      string
		userID        int64
		expected      []*r.Task
		expectedError error
	}

	testData := []testCreateTask{
		{
			testCase: "success",
			userID:   1,
			expected: []*r.Task{
				{
					ID:      1,
					UserID:  1,
					Content: "Do homework",
					Timestamp: r.Timestamp{
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
					},
				},
				{
					ID:      2,
					UserID:  1,
					Content: "Cleaning",
					Timestamp: r.Timestamp{
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
					},
				},
			},
		},
		{
			testCase:      "user does not exist",
			userID:        1,
			expected:      nil,
			expectedError: errors.New("user does not exist"),
		},
	}

	for _, tc := range testData {
		t.Run(tc.testCase, func(t *testing.T) {
			db, mock := NewMock()
			repo := &repository{db}
			defer func() {
				repo.Close()
			}()

			query := "SELECT * FROM tasks WHERE user_id = ?"
			rows := sqlmock.NewRows([]string{"id", "user_id", "content", "created_at", "updated_at"})
			if tc.expectedError != nil {
				mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(tc.expectedError)
			} else {
				for _, task := range tc.expected {
					rows.AddRow(task.ID, task.UserID, task.Content, task.CreatedAt, task.UpdatedAt)
				}
				mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)
			}

			tasks, err := repo.GetAllTaskByUserID(context.Background(), tc.userID)
			assert.Equal(t, len(tc.expected), len(tasks))
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

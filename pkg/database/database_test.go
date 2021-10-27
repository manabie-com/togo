package database

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatabase(t *testing.T) {
	t.Run("error_cases", testDatabaseErrorCases)
	t.Run("success_cases", testDatabaseSuccessCases)
}

func testDatabaseErrorCases(t *testing.T) {
	ctx := context.Background()
	connector, mock, err := NewSqlMockConnector()
	assert.NoError(t, err)
	db := NewDatabase(connector)

	// Test without connection
	err = db.WithoutTransaction(ctx, func(ctx context.Context, conn Connection) error {
		return nil
	})
	assert.Error(t, err, "should be return error: forgot connect to database")

	err = db.Transaction(ctx, func(ctx context.Context, conn Connection) error {
		return nil
	})
	assert.Error(t, err, "should be return error: forgot connect to database")

	// Test with connection
	err = db.Connect(nil)
	assert.NoError(t, err)

	expectedErr := errors.New("error occurred")
	err = db.WithoutTransaction(ctx, func(ctx context.Context, conn Connection) error {
		return expectedErr
	})
	assert.ErrorIs(t, err, expectedErr, "should be return error: "+expectedErr.Error())

	mock.ExpectBegin()
	mock.ExpectRollback()
	err = db.Transaction(ctx, func(ctx context.Context, conn Connection) error {
		return expectedErr
	})
	assert.ErrorIs(t, err, expectedErr, "should be return error: "+expectedErr.Error())
	if err := mock.ExpectationsWereMet(); err != nil {
		assert.NoError(t, err)
	}
}

func testDatabaseSuccessCases(t *testing.T) {
	ctx := context.Background()
	connector, mock, err := NewSqlMockConnector()
	assert.NoError(t, err)
	db := NewDatabase(connector)

	err = db.Connect(nil)
	assert.NoError(t, err)

	err = db.WithoutTransaction(ctx, func(ctx context.Context, conn Connection) error {
		return nil
	})
	assert.NoError(t, err)

	mock.ExpectBegin()
	mock.ExpectCommit()
	err = db.Transaction(ctx, func(ctx context.Context, conn Connection) error {
		return nil
	})
	assert.NoError(t, err)
	if err := mock.ExpectationsWereMet(); err != nil {
		assert.NoError(t, err)
	}
}

package txmanager

import (
	"context"

	"gorm.io/gorm"
)

const TX string = "tx"

//go:generate mockgen -destination=./mock_$GOFILE -source=$GOFILE -package=txmanager
type TransactionManager interface {
	Begin(ctx context.Context) TransactionManager
	End(ctx context.Context, err error) error
	Rollback(ctx context.Context) error
	Commit(ctx context.Context) error
	Recover(ctx context.Context)
	InjectTransaction(ctx context.Context) context.Context
}

func NewTransactionManager(db *gorm.DB) TransactionManager {
	return &txManager{
		DB: db,
	}
}

type txManager struct {
	DB     *gorm.DB
	isDone bool
}

func (tx *txManager) Begin(ctx context.Context) TransactionManager {
	return &txManager{
		DB: tx.DB.Begin().WithContext(ctx),
	}
}

func (tx *txManager) End(ctx context.Context, err error) error {
	if err != nil {
		return tx.Rollback(ctx)
	}

	return tx.Commit(ctx)
}

func (tx *txManager) Rollback(ctx context.Context) error {
	if tx.isDone {
		return nil
	}

	tx.isDone = true
	return tx.DB.Rollback().WithContext(ctx).Error
}

func (tx *txManager) Commit(ctx context.Context) error {
	if tx.isDone {
		return nil
	}

	tx.isDone = true
	return tx.DB.Commit().WithContext(ctx).Error
}

func (tx *txManager) Recover(ctx context.Context) {
	if r := recover(); r != nil {
		tx.Rollback(ctx)
		panic(r)
	}
}

func (tx *txManager) InjectTransaction(ctx context.Context) context.Context {
	return context.WithValue(ctx, TX, tx.DB)
}

func GetTxFromContext(ctx context.Context, db *gorm.DB) *gorm.DB {
	tx, ok := ctx.Value(TX).(*gorm.DB)
	if !ok || tx == nil {
		return db
	}
	return tx
}

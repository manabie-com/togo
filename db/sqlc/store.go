package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	common "togo/common/model"
)

type Store interface {
	Querier
	CreateTaskTx(ctx context.Context, arg CreateTaskTxParams) (CreateTaskTxResult, error)
}

type DBStore struct {
	db *sql.DB
	*Queries
}

func NewStore(db *sql.DB) Store {
	return &DBStore{
		db:      db,
		Queries: New(db),
	}
}

func (store *DBStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

// Debugging concurrent transactions
// var txKey = struct{}{}

type CreateTaskTxParams struct {
	User    User   `json:"user"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

type CreateTaskTxResult struct {
	User common.UserResponse `json:"user"`
	Task Task                `json:"task"`
}

func (store *DBStore) CreateTaskTx(ctx context.Context, arg CreateTaskTxParams) (CreateTaskTxResult, error) {
	var result CreateTaskTxResult
	err := store.execTx(ctx, func(q *Queries) error {
		// txName := ctx.Value(txKey)
		user := common.UserResponse{
			Username:         arg.User.Username,
			FullName:         arg.User.FullName,
			Email:            arg.User.Email,
			DailyCap:         arg.User.DailyCap,
			DailyQuantity:    arg.User.DailyQuantity,
			PasswordChangeAt: arg.User.PasswordChangeAt,
			CreatedAt:        arg.User.CreatedAt,
		}
		// check if dailyQuantity+1 > dailyCap?
		if user.DailyQuantity >= user.DailyCap {
			result = CreateTaskTxResult{
				User: user,
				Task: Task{},
			}
			return errors.New("daily limit exceed")
		} else {
			count, err := q.CountTasksCreatedToday(ctx, user.Username)
			if err != nil {
				return err
			}
			if user.DailyQuantity != count {
				user.DailyQuantity = 0
			}
			user.DailyQuantity++
			_, err = q.UpdateUserDailyQuantity(ctx, UpdateUserDailyQuantityParams{
				Username:      user.Username,
				DailyQuantity: user.DailyQuantity,
			})
			if err != nil {
				return err
			}
			task, err := q.CreateTask(ctx, CreateTaskParams{
				Name:    arg.Name,
				Owner:   arg.User.Username,
				Content: arg.Content,
			})
			result = CreateTaskTxResult{
				User: user,
				Task: task,
			}
			return err
		}
	})
	return result, err
}

package repositories

import (
	"database/sql"
	"sync"
	"context"
	"manabie.com/internal/common"
)

type RepositoryFactory struct {
	db *sql.DB
	transactions map[TransactionId]*sql.Tx
	mtx sync.Mutex
	count int
}

func MakeRepositoryFactory(
	iDb *sql.DB,
) RepositoryFactory {
	return RepositoryFactory {
		db: iDb,
		count: 0,
		transactions: map[TransactionId]*sql.Tx{},
	}
}

func isolationLevelToSqlIsolationLevel(iLevel TransactionLevel) sql.IsolationLevel {
	switch iLevel {
	case ReadUncommitted:
		return sql.LevelReadUncommitted
	case ReadCommitted:
		return sql.LevelReadCommitted
	case RepeatableRead:
		return sql.LevelRepeatableRead
	case Serializable:
		return sql.LevelSerializable
	}
	return sql.LevelReadCommitted
}

func (f *RepositoryFactory) StartTransactionAuto(
	iContext context.Context, 
	iIsolationLevel TransactionLevel,
	iHandler TransactionHandler,
) error {
	tx, err := f.db.BeginTx(iContext, &sql.TxOptions{Isolation: isolationLevelToSqlIsolationLevel(iIsolationLevel)})
	if err != nil {
		return err
	}
	f.mtx.Lock()
	count := TransactionId(f.count)
	f.transactions[count] = tx
	f.count += 1
	f.mtx.Unlock()
	err = iHandler(count)

	f.mtx.Lock()
	delete(f.transactions, count)
	f.mtx.Unlock()

	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}

	return err
}

func (f *RepositoryFactory) GetTaskRepository(iId TransactionId) (TaskRepositoryI, error) {
	f.mtx.Lock()
	defer f.mtx.Unlock()
	if tx, ok := f.transactions[iId]; !ok {
		return TaskRepositorySql{}, common.NotFound
	} else {
		repo := MakeTaskRepositorySql(tx)
		return repo, nil
	}
}

func (f *RepositoryFactory) GetUserRepository(iId TransactionId) (UserRepositoryI, error) {
	f.mtx.Lock()
	defer f.mtx.Unlock()
	if tx, ok := f.transactions[iId]; !ok {
		return UserRepositorySql{}, common.NotFound
	} else {
		repo := MakeUserRepositorySql(tx)
		return repo, nil
	}
}
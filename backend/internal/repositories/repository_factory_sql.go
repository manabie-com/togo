package repositories

import (
	"database/sql"
	"sync"
)

type RepositoryFactoriesFactory struct {
	db *sql.DB
	transactions map[TransactionId]*sql.Tx
	mu sync.Mutex
	count int
}
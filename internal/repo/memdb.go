package repo

import (
	"github.com/hashicorp/go-memdb"
)

type InMemConn struct {
	*memdb.MemDB
}

func (conn *InMemConn) GetTxn() (interface{}, error) {
	return conn.Txn(true), nil
}

type InMemStorage struct {
	Schema *memdb.DBSchema
}

func (store *InMemStorage) Connect() (Conn, error) {
	// Create a new data base
	db, err := memdb.NewMemDB(store.Schema)
	if err != nil {
		panic(err)
	}
	return &InMemConn{db}, nil
}

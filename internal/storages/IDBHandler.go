package storages

type IDBHandler interface {
	Execute(statement string) error
	Query(statement string) (IRow, error)
}

type IRow interface {
	Scan(dest ...interface{}) error
	Next() bool
}
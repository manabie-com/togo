package repo

type Conn interface {
	GetTxn() (interface{}, error)
}

type Storage interface {
	Connect() (Conn, error)
}

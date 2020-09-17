package constants

type DBType string

const (
	POSTGRES DBType = "POSTGRES"
	SQLITE DBType = "SQLITE"
)

const (
	HOST     = "localhost"
	PORT     = 5432
	USER     = "postgres"
	PASSWORD = "mysecretpassword"
	DB_NAME   = "togo-db"
	DB_TYPE = POSTGRES
)

const (
	JWT_KEY = "wqGyEBBfPK9w3Lxw"
)
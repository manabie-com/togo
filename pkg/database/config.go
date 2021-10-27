package database

type Config struct {
	UserName        string
	Password        string
	Address         string
	Database        string
	NumberMaxConns  int
	NumberIdleConns int
}

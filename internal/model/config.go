package model

type AppSettings struct {
	DatabaseSettings     DatabaseSettings     `json:"databaseSettings"`
}

type DatabaseSettings struct {
	ConnectionString   string `json:"connectionString"`
	MaxIdleConnections int    `json:"maxIdleConnections"`
	MaxOpenConnections int    `json:"maxOpenConnections"`
}

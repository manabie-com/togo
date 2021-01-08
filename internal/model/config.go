package model

const (
	DatabaseProviderPostgresql = "postgresql"
	DatabaseProviderSQLite     = "sqlite"
)

type AppSettings struct {
	JWTSecretKey     string           `json:"jwtSecretKey"`
	DatabaseSettings DatabaseSettings `json:"databaseSettings"`
}

type DatabaseSettings struct {
	Provider   string     `json:"provider"`
	Postgresql Postgresql `json:"postgresql"`
	SQLite     SQLite     `json:"sqlite"`
}

type Postgresql struct {
	ConnectionString   string `json:"connectionString"`
	MaxIdleConnections int    `json:"maxIdleConnections"`
	MaxOpenConnections int    `json:"maxOpenConnections"`
}

type SQLite struct {
	DriverName     string `json:"driverName"`
	DataSourceName string `json:"dataSourceName"`
}

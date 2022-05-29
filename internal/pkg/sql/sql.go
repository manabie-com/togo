package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/dinhquockhanh/togo/internal/pkg/config"
)

func NewSqlConnection(config *config.DB) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.Name)
	dbCon, err := sql.Open(config.Driver, connStr)
	if err != nil {
		return nil, err
	}
	return dbCon, nil
}

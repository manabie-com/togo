package repositories

import (
	"database/sql"
	"os"
	"fmt"
	_ "github.com/lib/pq"
	"go.uber.org/dig"
)

func ConnectPostgres() *sql.DB {
	connStr := ""
	connStr += " user=" + os.Getenv("PGUSER")
	connStr += " dbname=" + os.Getenv("PGDATABASE")
	connStr += " password=" + os.Getenv("PGPASSWORD")
	connStr += " host=" + os.Getenv("PGHOST")
	connStr += " sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(fmt.Sprintf("could not connect to database %s", err.Error()))
	}

	return db
}

func ProvideSqlConnection(iContainer *dig.Container) error {
	return iContainer.Provide(ConnectPostgres)
}
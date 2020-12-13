package sqllite

import "fmt"

var (
	pgPassword = "postgres"
	pgDbName   = "togo"
	pgHost     = "localhost"
	pgUser     = "postgres"
	pgPort     = "5432"
)

// DBConnectionURL returns URL to connect to database
func DBConnectionURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		pgUser, pgPassword, pgHost, pgPort, pgDbName)
}

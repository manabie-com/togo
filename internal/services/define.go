package services

import "fmt"

const (
	host     = "postgres"
	port     = 5432
	user     = "dev"
	password = "passwd"
	dbname   = "togo"
)

var (
	psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
)

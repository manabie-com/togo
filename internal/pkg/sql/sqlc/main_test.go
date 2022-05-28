package db

import (
	"log"
	"os"
	"testing"

	"github.com/dinhquockhanh/togo/internal/pkg/config"
	db "github.com/dinhquockhanh/togo/internal/pkg/sql"
	_ "github.com/lib/pq"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	var err error
	cf := config.DB{
		Host:     "localhost",
		Port:     "5432",
		User:     "username",
		Password: "password",
		Name:     "togo",
		Driver:   "postgres",
	} // TODO: load config from file.
	cnn, err := db.NewSqlConnection(&cf)
	if err != nil {
		log.Fatalln("cannot connect to DB: ", err.Error())
	}
	testQueries = New(cnn)
	os.Exit(m.Run())
}

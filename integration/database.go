package integration

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"strings"
)

var (
	dbConn *sql.DB
	qSeed  []string
)

func prepareSeedData() error {
	seed, err := ioutil.ReadFile("./seed.sql")
	if err != nil {
		return err
	}
	qSeed = strings.Split(string(seed), "\n\n")
	return nil
}

func seedDB() error {
	for i, query := range qSeed {
		_, err := dbConn.Exec(query)
		if err != nil {
			return fmt.Errorf("failed to exececute query % of seed.sql, error: %s", i, err.Error())
		}
	}
	return nil
}

func refreshDB() error {
	_, err := dbConn.Exec(`SELECT truncate_tables();`)
	if err != nil {
		return err
	}
	return seedDB()
}

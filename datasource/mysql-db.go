package datasource

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/manabie-com/backend/utils"
)

func Conn() (*sql.DB, error) {
	config, errConfig := utils.LoadConfig(".")
	if errConfig != nil {
		log.Fatal("cannot load config datasource")
	}
	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		config.MYSQL_USERNAME,
		config.MYSQL_PASSWORD,
		config.MYSQL_HOST,
		config.MYSQL_DB,
	)

	var err error
	Client, err := sql.Open("mysql", dataSource)
	if err != nil {
		return nil, err
	}

	if err = Client.Ping(); err != nil {
		return nil, err
	}

	log.Println("Data sucessfully configured.")

	return Client, nil
}

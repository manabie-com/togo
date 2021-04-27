package models

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"

	_ "github.com/lib/pq"
	"github.com/manabie-com/togo/internal/app/config"
	"github.com/manabie-com/togo/internal/utils"
)

var configs = config.LoadConfigs()

type DBInterface interface {
	Connect() *sql.DB
}

func Connect() *sql.DB {
	URL := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		configs.Database.User,
		configs.Database.Pass,
		configs.Database.Host,
		configs.Database.Port,
		configs.Database.Name)
	var db *sql.DB
	db, err := sql.Open("postgres", URL)

	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}

func MigrationDB() error {
	con := Connect()
	defer con.Close()
	var err error

	sqlStr := `
			DROP TABLE IF EXISTS users CASCADE;
			DROP TABLE IF EXISTS tasks CASCADE;

			CREATE TABLE IF NOT EXISTS users
			(
				id       BIGINT             NOT NULL,
				username VARCHAR(50)        NOT NULL,
				password TEXT               NOT NULL,
				max_todo BIGINT DEFAULT 5   NOT NULL,
				CONSTRAINT users_PK PRIMARY KEY (id)
			);

			INSERT INTO users (id, username, password, max_todo) VALUES ('100', 'firstUser', 'example', 5);
			INSERT INTO users (id, username, password, max_todo) VALUES ('200', 'secondUser', 'example', 5);

			CREATE TABLE IF NOT EXISTS tasks
			(
				id           BIGINT     NOT NULL,
				content      TEXT       NOT NULL,
				user_id      BIGINT     NOT NULL,
				created_date timestamp  NOT NULL,
				CONSTRAINT tasks_PK PRIMARY KEY (id),
				CONSTRAINT tasks_FK FOREIGN KEY (user_id) REFERENCES users (id)
			);
	`
	sqlRaw := strings.Split(sqlStr, ";")
	for _, stmt := range sqlRaw {
		_, err = con.Exec(stmt)
		if err != nil {
			return err
		}
	}

	return nil
}

func TestConnection() {
	con := Connect()
	defer con.Close()
	err := con.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database connected!")
}

func RefreshTable(resp http.ResponseWriter, req *http.Request) {
	con := Connect()
	defer con.Close()

	sqlStr := "TRUNCATE tasks CASCADE;"
	_, err := con.Exec(sqlStr)
	if err != nil {
		utils.JSON(resp, http.StatusInternalServerError, nil)
	}

	utils.JSON(resp, http.StatusOK, nil)
}

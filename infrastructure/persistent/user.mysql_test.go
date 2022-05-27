package persistent_test

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"
	"togo/infrastructure/persistent"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var db *sql.DB

func TestMain(m *testing.M) {
	var err error
	db, err = sql.Open("mysql", fmt.Sprintf("root:123456aA@(localhost:3309)/mysql?multiStatements=true"))
	if err != nil {
		log.Fatalf("Cant connect to db: %s", err)
	}
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	migrateIns, err := migrate.NewWithDatabaseInstance(
		"file://migrates",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatalf("Cant create migrate Instance: %s", err)
	}
	migrateIns.Steps(2)
	err = db.Ping()
	if err != nil {
		log.Fatalf("Cant connect to db: %s", err)
	}
	log.Println("Connect success")
	code := m.Run()
	os.Exit(code)
}

func TestNewUserMysqlRepository(t *testing.T) {
	repo := persistent.NewUserMysqlRepository(db)
	t.Logf("%#v", repo)
}

func TestUserMysqlRepo_Create(t *testing.T) {

}

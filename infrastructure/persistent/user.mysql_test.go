package persistent_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"
	"togo/domain/model"
	"togo/infrastructure/persistent"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var db *sql.DB
var user = model.User{
	Username: "admin",
	Password: "admin",
	Limit:    10,
}

func TestMain(m *testing.M) {
	var err error
	db, err = sql.Open("mysql", fmt.Sprintf("root:123456aA@(localhost:3309)/todo?multiStatements=true"))
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
	repo := persistent.NewUserMysqlRepository(db)
	repo.Delete(context.Background(), user.Username)
	err := repo.Create(context.Background(), user)
	if err != nil {
		t.Errorf("%s", err.Error())
		return
	}
	t.Log("Create user success")
}

func TestUserMysqlRepo_Get(t *testing.T) {
	repo := persistent.NewUserMysqlRepository(db)
	u2, err := repo.Get(context.Background(), user.Username)
	if err != nil {
		t.Errorf("%s", err.Error())
		return
	}
	t.Logf("User: %#v", u2)
}

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

func TestMain(m *testing.M) {
	var err error
	db, err = sql.Open("mysql", fmt.Sprintf("root:123456aA@(localhost:3309)/todo?multiStatements=true"))
	if err != nil {
		log.Fatalf("Cant connect to db: %s", err)
	}
	//driver, _ := mysql.WithInstance(db, &mysql.Config{})
	//migrateIns, err := migrate.NewWithDatabaseInstance(
	//	"file://migrations",
	//	"todo",
	//	driver,
	//)
	//if err != nil {
	//	log.Fatalf("Cant create migrate Instance: %s", err)
	//}
	////migrateIns.Drop()
	//err = migrateIns.Up()
	//if err != nil {
	//	log.Fatalf("migrateIns fail: %s", err)
	//}
	//err = db.Ping()
	//if err != nil {
	//	log.Fatalf("Cant connect to db: %s", err)
	//}
	log.Println("Connect success")
	code := m.Run()
	//
	//err = migrateIns.Down()
	//migrateIns.Close()
	os.Exit(code)
}

func TestNewUserMysqlRepository(t *testing.T) {
	repo := persistent.NewUserMysqlRepository(db)
	t.Logf("%#v", repo)
}

func TestUserMysqlRepo_Create(t *testing.T) {
	repo := persistent.NewUserMysqlRepository(db)
	user := model.User{
		Username: "admin",
		Password: "admin",
		Limit:    10,
	}
	err := repo.Create(context.Background(), user)
	if err != nil {
		t.Errorf("%s", err.Error())
		return
	}
	t.Log("Create user success")
}

func TestUserMysqlRepo_Get(t *testing.T) {
	repo := persistent.NewUserMysqlRepository(db)
	user := model.User{
		Username: "admin",
		Password: "admin",
		Limit:    10,
	}
	u2, err := repo.Get(context.Background(), user.Username)
	if err != nil {
		t.Errorf("%s", err.Error())
		return
	}
	t.Logf("User: %#v", u2)
}

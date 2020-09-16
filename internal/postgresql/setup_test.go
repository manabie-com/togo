package postgresql

import (
  "context"
  "database/sql"
  "fmt"
  "github.com/google/uuid"
  "github.com/manabie-com/togo/internal/config"
  "github.com/manabie-com/togo/internal/core"
  "golang.org/x/crypto/bcrypt"
  "io/ioutil"
  "log"
  "os"
  "strings"
  "testing"
  "time"
)

var db *sql.DB
var resetStatements []string
var userRepo UserRepo
var taskRepo TaskRepo

const bcryptCost = 10

var hashedPassword, _ = bcrypt.GenerateFromPassword([]byte("password"), bcryptCost)
var firstUser = core.User{
  ID:      "firstUser",
  Hash:    string(hashedPassword),
  MaxTodo: 5,
}
var task = core.Task{
  ID:          uuid.New().String(),
  Content:     "content",
  UserID:      firstUser.ID,
  CreatedDate: time.Now(),
}

func reset() {
  if len(resetStatements) == 0 {
    file, err := ioutil.ReadFile("./reset.sql")
    if err != nil {
      log.Fatalf("postgresql::test - Error loading file: %v", err)
    }
    resetStatements = strings.Split(string(file), "\n\n")
  }
  for _, stmt := range resetStatements {
    _, err := db.ExecContext(context.Background(), stmt)
    if err != nil {
      fmt.Printf("%v\n", err)
    }
  }
  userRepo = UserRepo{
    DB:   db,
    Cost: 10,
  }
  taskRepo = TaskRepo{DB: db}
}

func TestMain(t *testing.M) {
  var err error
  config.Load("../../.env")
  db, err = sql.Open("postgres", config.PostgresDataSourceTest)
  if err != nil {
    log.Fatalf("postgresql::test - Error connecting database: %v", err)
  }
  os.Exit(t.Run())
}

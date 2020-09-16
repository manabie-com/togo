package integration

import (
  "context"
  "database/sql"
  "fmt"
  "github.com/dgrijalva/jwt-go"
  "github.com/manabie-com/togo/internal/config"
  "github.com/manabie-com/togo/internal/core"
  "github.com/manabie-com/togo/internal/http"
  "github.com/manabie-com/togo/internal/postgresql"
  "io/ioutil"
  "log"
  "os"
  "strings"
  "testing"
  "time"
)

type jwtManager struct {
  Key        string
  Method     jwt.SigningMethod
  ExpireTime time.Duration
}

func (manager *jwtManager) Token(user *core.User) (string, error) {
  claims := jwt.MapClaims{}
  claims["user_id"] = user.ID
  claims["exp"] = time.Now().Add(manager.ExpireTime).Unix()
  token := jwt.NewWithClaims(manager.Method, claims)
  signed, err := token.SignedString([]byte(manager.Key))
  if err != nil {
    return "", err
  }
  return signed, nil
}

var db *sql.DB
var resetStatements []string

var server *http.Server
//
//const bcryptCost = 10
//
//var hashedPassword, _ = bcrypt.GenerateFromPassword([]byte("password"), bcryptCost)
//var firstUser = core.User{
//  ID:      "firstUser",
//  Hash:    string(hashedPassword),
//  MaxTodo: 5,
//}
//var task = core.Task{
//  ID:          uuid.New().String(),
//  Content:     "content",
//  UserID:      firstUser.ID,
//  CreatedDate: time.Now(),
//}

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

  server = http.DefaultServer(&postgresql.UserRepo{
    DB:   db,
    Cost: 10,
  }, &postgresql.TaskRepo{DB: db})
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

package main

import (
  "database/sql"
  "fmt"
  "github.com/joho/godotenv"
  _ "github.com/lib/pq"
  "github.com/manabie-com/togo/internal/http"
  "github.com/manabie-com/togo/internal/postgresql"
  "os"
  "strconv"

  //"github.com/manabie-com/togo/internal/sqlite"
  //_ "github.com/mattn/go-sqlite3"
  "log"
)

var port int
var dataSource string
var jwtKey string

const MODE_LOCAL = "local"

func init() {
  var err error
  runMode := os.Getenv("RUN_MODE")
  if runMode == MODE_LOCAL {
    err := godotenv.Load()
    if err != nil {
      log.Fatalf("togo::init - Error loading .env file: %v", err)
    }
  }

  dataSource = os.Getenv("DATASOURCE")
  if dataSource == "" {
    log.Fatal("togo::init - Data source not set")
  }
  jwtKey = os.Getenv("JWT_KEY")
  if jwtKey == "" {
    log.Fatal("togo::init - JWT key not set")
  }

  portStr := os.Getenv("TOGO_PORT")
  if portStr == "" {
    log.Fatal("togo::init - application port not set")
  }
  port, err = strconv.Atoi(portStr)
  if err != nil {
    log.Fatal("togo::init - invalid port")
  }
}

func main() {
  fmt.Println("togo::main - Connecting database...")
  db, err := sql.Open("postgres", dataSource)
  if err != nil {
    log.Fatalf("togo::main - Error connecting database: %v", err)
  }
  fmt.Println("togo::main - Database connected")

  userRepo := &postgresql.UserRepo{DB: db, Cost: 10}
  taskRepo := &postgresql.TaskRepo{DB: db}
  fmt.Println(fmt.Sprintf("togo::main - Listening on port %v...", port))
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), http.JSONServer(jwtKey, userRepo, taskRepo)))
}

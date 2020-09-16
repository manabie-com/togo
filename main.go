package main

import (
  "database/sql"
  "fmt"
  _ "github.com/lib/pq"
  "github.com/manabie-com/togo/internal/config"
  "github.com/manabie-com/togo/internal/http"
  "github.com/manabie-com/togo/internal/postgresql"
  //"github.com/manabie-com/togo/internal/sqlite"
  _ "github.com/mattn/go-sqlite3"
  "log"
)

func main() {
  config.Load()
  fmt.Println("togo::main - Connecting database...")
  db, err := sql.Open(string(config.Dialect), config.DataSource)
  if err != nil {
    log.Fatalf("togo::main - Error connecting database: %v", err)
  }
  fmt.Println("togo::main - Database connected")

  userRepo := &postgresql.UserRepo{DB: db, Cost: 10}
  taskRepo := &postgresql.TaskRepo{DB: db}
  server := http.DefaultServerWithSPA(userRepo, taskRepo)
  fmt.Println(fmt.Sprintf("togo::main - Listening on port %v...", config.Port))
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.Port), server))
}

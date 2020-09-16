package config

import (
  "github.com/joho/godotenv"
  "log"
  "os"
  "strconv"
)

type DBDialect string
const (
  POSTGRESQL DBDialect = "postgres"
  SQLITE DBDialect = "sqlite"
)

type RunMode string
const (
  MODE_DEVELOPMENT RunMode = "development"
  MODE_PRODUCTION RunMode = "production"
)

var Port int
var PostgresDataSource string
var SqliteDataSource string
var PostgresDataSourceTest string
var SqliteDataSourceTest string
var DataSource string
var Dialect DBDialect
var JwtKey string

func Load(envPath ...string) {
  var err error
  runMode := os.Getenv("RUN_MODE")
  if runMode == string(MODE_DEVELOPMENT) {
    err := godotenv.Load(envPath...)
    if err != nil {
      log.Fatalf("config::init - Error loading .env file: %v", err)
    }
  }

  PostgresDataSourceTest = os.Getenv("POSTGRES_DATASOURCE_TEST")
  SqliteDataSourceTest = os.Getenv("SQLITE_DATASOURCE_TEST")

  PostgresDataSource = os.Getenv("POSTGRES_DATASOURCE")
  SqliteDataSource = os.Getenv("SQLITE_DATASOURCE")
  dialect := os.Getenv("DIALECT")
  switch dialect {
  case string(SQLITE):
    Dialect = SQLITE
    DataSource = SqliteDataSource
  case string(POSTGRESQL):
    Dialect = POSTGRESQL
    DataSource = PostgresDataSource
  default:
    log.Fatal("config::init - Unknown database dialect")
  }
  if DataSource == "" {
    log.Fatal("config::init - Data source not set")
  }
  JwtKey = os.Getenv("JWT_KEY")
  if JwtKey == "" {
    log.Fatal("config::init - JWT key not set")
  }

  portStr := os.Getenv("TOGO_PORT")
  if portStr == "" {
    log.Fatal("config::init - Application port not set")
  }
  Port, err = strconv.Atoi(portStr)
  if err != nil {
    log.Fatal("config::init - Invalid port")
  }
}

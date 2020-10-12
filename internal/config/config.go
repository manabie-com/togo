package config

import (
	"fmt"
	"os"
)

// Config : address service configurations
type Config struct {
	ServerPort string `json: "server_port"`
	PGHost     string `json: "pg_host"`
	PGPort     string `json: "pg_port"`
	PGUser     string `json: "pg_user"`
	PGPassword string `json: "pg_password"`
	PGDBName   string `json: "pg_db_name"`
}

var config *Config

func init() {
	serverPort := os.Getenv("server_port")
	pgHost := os.Getenv("pg_host")
	pgPort := os.Getenv("pg_port")
	pgUser := os.Getenv("pg_user")
	pgPassword := os.Getenv("pg_password")
	pgDBName := os.Getenv("pg_db_name")

	config = &Config{
		ServerPort: serverPort,
		PGHost:     pgHost,
		PGPort:     pgPort,
		PGUser:     pgUser,
		PGPassword: pgPassword,
		PGDBName:   pgDBName,
	}
}

// GetConfig :
func GetConfig() *Config {
	return config
}

// Print configurations
func (conf *Config) Print() {
	fmt.Println("------------ Configurations --------------")
	fmt.Printf("Postgres host:\t\t%s\n", conf.PGHost)
	fmt.Printf("Postgres port:\t\t%s\n", conf.PGPort)
	fmt.Printf("Postgres user:\t\t%s\n", conf.PGUser)
	fmt.Printf("Postgres password:\t%s\n", conf.PGPassword)
	fmt.Printf("Postgres DB name:\t%s\n", conf.PGDBName)
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/manabie-com/togo/internal/api"
	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	SVPort     string `mapstructure:"SV_PORT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}

func (cf Config) PsqlInfo() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cf.DBHost, cf.DBPort, cf.DBUser, cf.DBPassword, cf.DBName)
}

func (cf Config) ServerPort() string {
	return fmt.Sprintf(":%s", cf.SVPort)
}

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("You should run application with config file")
	}
	config, err := LoadConfig(args[1])
	if err != nil {
		log.Fatal("You should run application with valid file path", err)
	}
	db, err := sqlx.Open("postgres", config.PsqlInfo())
	if err != nil {
		log.Fatal("error opening db", err)
	}
	todoApi := api.NewToDoApi("wqGyEBBfPK9w3Lxw", db)
	err = http.ListenAndServe(config.ServerPort(), &todoApi)
	if err != nil {
		log.Fatal("error listen and serve api", err)
	}
}

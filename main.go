package main

import (
	"github.com/manabie-com/togo/internal/model"
	"github.com/manabie-com/togo/internal/transport"
	"github.com/manabie-com/togo/internal/usecase"
	"github.com/spf13/viper"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
		return
	}

	conf := model.AppSettings{}
	err = viper.Unmarshal(&conf)
	if err != nil {
		log.Fatal(err)
		return
	}

	a, err := usecase.New(conf)
	if err != nil {
		log.Fatal(err)
		return
	}

	http.ListenAndServe(":5050", &transport.ToDoService{
		JWTKey:  conf.JWTSecretKey,
		Usecase: a,
	})
}

package main

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/manabie-com/togo/internal/storages/psql"
	"github.com/manabie-com/togo/internal/transport/http"
	"github.com/manabie-com/togo/internal/usecase"
)

var (
	conf struct {
		Transport  http.APIConf
		AuthUC     usecase.AuthUCConf `envconfig:"AUTH"`
		Storage    psql.Config
		Port       int
		MetricPort int
	}
)

func init() {
	err := envconfig.Process("app", &conf)
	if err != nil {
		panic(err)
	}
}

func main() {
	storage, err := psql.NewStorage(conf.Storage)
	if err != nil {
		panic(err)
	}
	authUc, err := usecase.NewAuthUseCase(conf.AuthUC, storage)
	if err != nil {
		panic(err)
	}
	if len(os.Args) < 2 {
		fmt.Println("Please insert more command")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "add-admin":
		if len(os.Args) != 4 {
			fmt.Println("add-admin needs 2 arguments")
			os.Exit(1)
		}
		err := authUc.CreateUser(os.Args[2], os.Args[3])
		if err != nil {
			fmt.Printf("Error creating user: %s\n", err)
			os.Exit(1)
		}
		fmt.Println("Success!!!")
	}

}

package main

import (
	server "github.com/HoangMV/todo/src"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func main() {
	server := server.New()
	server.SetupConfig()
	server.Run()
}

package main

import (
	"github.com/huuthuan-nguyen/manabie/app"
	"github.com/huuthuan-nguyen/manabie/app/utils"
)

func main() {
	config := utils.ReadConfig() // read config from env
	application := app.NewApp(config)
	application.Run()
}

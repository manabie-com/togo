package main

import (
	"ManabieProject/env"
	"ManabieProject/src/dbcontrol"
	routes "ManabieProject/src/router"
)

func main() {
	env.Init()
	dbcontext.Init()
	router := routes.InitRouter()
	go routes.ResetCounterTask()
	router.Run("0.0.0.0:33333")

}

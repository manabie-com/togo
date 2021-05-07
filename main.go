package main

import (
	"manabie-com/togo/controller"
	"manabie-com/togo/entity"
	"manabie-com/togo/global"
	pkg_logrus "manabie-com/togo/pkg/logger"
	pkg_rd "manabie-com/togo/pkg/rd"
)

func main() {
	global.FetchProductionEnvironmentVariables()

	entity.InitializeDb()

	pkg_logrus.InitLogrus()

	pkg_rd.InitializeRdV8()

	controller.Initialize()
}

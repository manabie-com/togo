package main

import (
	"github.com/manabie-com/togo/internal/app"
)

func main() {
	app := new(app.APP)
	app.Initialize()
	app.Serve()
}

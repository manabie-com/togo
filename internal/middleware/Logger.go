package middleware

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	fiberLog "github.com/gofiber/fiber/v2/middleware/logger"
)

func Logger(app *fiber.App) {

	currentTime := time.Now()
	file, err := os.OpenFile("../logs/"+currentTime.Format("01022006")+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	app.Use(fiberLog.New(fiberLog.Config{
		Format: "-----------------------------------------------------------\n" +
			"PID:${pid} STATUS:${status} - Method:${method} Path:${path}\n" +
			"Response Body: ${resBody}\n",
		TimeFormat: "02-Jan-2006",
		Output:     file,
	}))

}

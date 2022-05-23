package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
	http "net/http"
	"os"
	"togo/database"
	taskHandler "togo/internal/task/controller/http"
	taskRepo "togo/internal/task/repository"
	taskService "togo/internal/task/service"
	userRepo "togo/internal/user/repository"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	//init database
	db := database.InitDB()

	//binding Repository
	taskRepo := taskRepo.NewTaskRepository(db)
	userRepo := userRepo.NewUserRepository(db)

	//binding Service
	taskService := taskService.NewTaskService(taskRepo, userRepo)
	//binding handler
	taskHandler.NewTaskHandler(e, taskService)

	httpPort := os.Getenv("HTTP_PORT")
	e.Logger.Fatal(e.Start(":" + httpPort))
}

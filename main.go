package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	appctx "github.com/manabie-com/togo/app_ctx"
	"github.com/manabie-com/togo/token_provider/jwt"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
)

func main() {

	DBUser := os.Getenv("DB_USER")
	DBPasswd := os.Getenv("DB_PASSWORD")
	DBHost := os.Getenv("DB_HOST")
	DBPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local&charset=%s",
		DBUser,
		DBPasswd,
		DBHost,
		DBPort,
		dbName,
		"utf8mb4",
	)
	tokenProvider := jwt.NewTokenJWTProvider(os.Getenv("JWT_SECRET"), 60*60*24*30)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	appCtx := appctx.NewAppContext(db.Debug(), tokenProvider)

	engine := gin.Default()
	setupHandlers(engine, appCtx)

	engine.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}

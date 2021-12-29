package main

import (
	"log"
	"os"
	"time"
	api2 "togo/handler"
	"togo/db"
	"togo/middleware"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	_ "github.com/golang/mock/gomock"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	r := gin.Default()

	db.SetupDatabaseConnection(false)

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "togo-api",
		Key:             []byte("secret key"),
		Timeout:         time.Hour,
		MaxRefresh:      time.Hour,
		IdentityKey:     middleware.IdentityKey,
		PayloadFunc:     middleware.PayloadFunc,
		IdentityHandler: middleware.IdentityHandler,
		Authenticator:   middleware.Authenticator,
		Authorizator:    middleware.Authorizator,
		Unauthorized:    middleware.Unauthorized,
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	errInit := authMiddleware.MiddlewareInit()
	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	api := &api2.APIEnv{
		DB: db.DB,
	}

	r.POST("/register", api.Register)
	r.POST("/login", authMiddleware.LoginHandler)

	auth := r.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.POST("/task", api.CreateTask)
		auth.GET("/task", api.GetTask)
	}

	log.Fatal(r.Run(":" + port))
}

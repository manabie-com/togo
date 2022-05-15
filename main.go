package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/japananh/togo/common"
	"github.com/japananh/togo/component"
	"github.com/japananh/togo/component/tokenprovider"
	"github.com/japananh/togo/middleware"
	"github.com/japananh/togo/modules/task/tasktransport/gintask"
	"github.com/japananh/togo/modules/user/usertransport/ginuser"
	"github.com/joho/godotenv"
	goose "github.com/pressly/goose/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	// load env from `.env` file
	if err := loadEnv(); err != nil {
		log.Fatalln("Missing env file")
	}

	dsn := os.Getenv("DB_CONNECTION_STR")
	secretKey := os.Getenv("SYSTEM_KEY")
	atExpiryStr := os.Getenv("ACCESS_TOKEN_EXPIRY")
	rtExpiryStr := os.Getenv("REFRESH_TOKEN_EXPIRY")

	if dsn == "" || secretKey == "" || atExpiryStr == "" || rtExpiryStr == "" {
		log.Fatalln("Missing some env")
	}

	// connect to database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	// run database migrations
	if err := runDBMigrations(db); err != nil {
		log.Fatalln(err)
	}

	// create token configs
	tokenConfig, err := tokenprovider.NewTokenConfig(atExpiryStr, rtExpiryStr)
	if err != nil {
		log.Fatalln(err)
	}

	// run api service
	if err := runService(db, secretKey, tokenConfig); err != nil {
		log.Fatalln(err)
	}
}

func loadEnv() error {
	cwd, _ := os.Getwd()
	return godotenv.Load(cwd + `/.env`)
}

func runDBMigrations(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("mysql"); err != nil {
		return err
	}

	if err := goose.Up(sqlDB, "migrations"); err != nil {
		return err
	}

	return nil
}

func runService(db *gorm.DB,
	secretKey string,
	tokenConfig *tokenprovider.TokenConfig,
) error {
	r := gin.Default()

	appCtx := component.NewAppContext(db, secretKey, tokenConfig)

	r.Use(middleware.Recover(appCtx))

	v1 := r.Group("/api/v1")

	v1.POST("/register", ginuser.Register(appCtx))
	v1.POST("/login", ginuser.Login(appCtx))

	tasks := v1.Group("/tasks", middleware.RequiredAuth(appCtx))
	{
		tasks.POST("/", gintask.CreateTask(appCtx))
	}

	// TODO: How to only show these APIs in development?
	// api for encode uid receives real id and database type, then return fake uid
	// e.g: id: 16, db_type: 2 -> fakeId: 3w5rMJ8raFkfXS
	v1.GET("/encode-uid", func(c *gin.Context) {
		type reqData struct {
			DBType int `form:"db_type" binding:"required"`
			RealId int `form:"id" binding:"required"`
		}

		var d reqData
		if err := c.ShouldBind(&d); err != nil {
			c.JSON(http.StatusBadRequest, "invalid request")
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id": common.NewUID(uint32(d.RealId), d.DBType, 1),
		})
	})

	// api for decode uid receives fake uid then return real id and database type
	// e.g: fakeId: 3w5rMJ8raFkfXS -> id: 16, db_type: 2
	v1.GET("/decode-uid", func(c *gin.Context) {
		type reqData struct {
			FakeId string `form:"id" binding:"required"`
		}

		var d reqData
		if err := c.ShouldBind(&d); err != nil {
			c.JSON(http.StatusBadRequest, "invalid request")
			return
		}

		realId, err := common.FromBase58(d.FakeId)
		if err != nil {
			c.JSON(http.StatusBadRequest, "invalid request")
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id":      realId.GetLocalID(),
			"db_type": realId.GetObjectType(),
		})
	})

	return r.Run()
}

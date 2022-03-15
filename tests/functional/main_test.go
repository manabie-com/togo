package functional_test

import (
	"os"
	"testing"
	"togo-service/app/models"
	"togo-service/pkg/database"
	"togo-service/pkg/middleware"
	"togo-service/pkg/routes"
	"togo-service/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var db *gorm.DB
var app *fiber.App
var user1 models.User
var user2 models.User

func TestMain(m *testing.M) {
	println("Start Testing")
	godotenv.Load("../../.env.testing")
	db = database.SetupDB()
	resetDB(db)

	app = fiber.New()
	middleware.FiberMiddleware(app)
	routes.Setup(app, db)

	os.Exit(m.Run())
}

func resetDB(db *gorm.DB) {
	// drops table
	db.Migrator().DropTable(&models.User{})
	db.Migrator().DropTable(&models.Setting{})
	db.Migrator().DropTable(&models.Task{})
	// migrate
	database.DoMigrate(db)

	pass, _ := utils.HashPassword("secret")

	db.Create(&models.User{
		Username: "admin",
		Password: pass,
		Role:     "admin",
	})

	// user1
	user1.Username = "user"
	user1.Password = pass
	db.Create(&user1)

	var setting1 models.Setting
	setting1.UserID = uint64(user1.ID)
	setting1.QuotaPerDay = 2
	db.Create(&setting1)

	// user2
	user2.Username = "user2"
	user2.Password = pass
	db.Create(&user2)

	var setting2 models.Setting
	setting2.UserID = uint64(user2.ID)
	setting2.QuotaPerDay = 2
	db.Create(&setting2)
}

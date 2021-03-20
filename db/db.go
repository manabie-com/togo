package db

import (
	"fmt"
	"github.com/manabie-com/togo/config"
	"github.com/manabie-com/togo/models"
	"github.com/manabie-com/togo/repositories"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var DB *gorm.DB //database

func ConnectDB() {
	DSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		"localhost", config.NewEnv.DbUser, config.NewEnv.DbPass, config.NewEnv.DbName, config.NewEnv.DbPort, "Asia/Ho_Chi_Minh")

	conn, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  DSN,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	log.Println(fmt.Sprintf("Succeed to connect to database: %s", config.NewEnv.DbName))

	DB = conn

}

func Migrate() {
	err := DB.AutoMigrate(&models.User{}, &models.Task{})

	if err != nil {
		log.Println(err)
		log.Fatal("Failed to migrate")
	}
}

func Seed() {
	userRepo := repositories.NewUserRepository(DB)
	//UserRepo.AddUser(&models.User{Username: "huyha", Password: "123456", MaxTodo: 3})
	aa, err := userRepo.GetUserByUserName("huyha")
	if err != nil {
		print("err", err)
	}
	print(aa.Password)
	//err = utils.VerifyPassword(aa.Password, "123456")
	//if err != nil {
	//	print("no no no")
	//}
	////log.Printf(err)
	////log.Printf(aa)
}

package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/manabie-com/togo/internal/config"
	"github.com/manabie-com/togo/internal/models"
)

//ConnectToDB will connect the BE to the database
func ConnectToDB() (*gorm.DB, error) {
	config := config.GetConfig()
	// dsn := ""host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai""
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.DatabaseHost,
		config.DatabaseUsername,
		config.DatabasePassword,
		config.DatabaseName,
		config.DatabasePort,
		config.DatabaseSslMode,
		config.DatabaseTimezone,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err == nil {
		fmt.Println("DB connected.")
		// fmt.Printf("DBURI: %s", dsn)
		return db, nil
	} else {
		fmt.Println("ERROR FOUND!")
		panic(err)
	}
}

//SyncDB will migrate & seed DB
func SyncDB(isForced bool) {
	//DB will be synced forcefully.
	if isForced {
		modelList := models.GetModelList()
		db, err := ConnectToDB()
		if err != nil {
			panic(err)
		}
		for i := 0; i < len(modelList); i++ {
			if db.Migrator().HasTable(modelList[i]) {
				db.Migrator().DropTable(modelList[i])
			}
			db.Migrator().AutoMigrate(modelList[i])
		}
		SeedDB(db)
	} else {
		fmt.Println("No seeding needed.")
	}

}

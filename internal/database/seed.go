package database

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/models"

	"gorm.io/gorm"
)

//SeedDB function will trigger all seed functions below
func SeedDB(db *gorm.DB) {
	accounts := []models.Account{}
	uuid1, _ := uuid.Parse("1f78cabc-b268-43cb-9935-c3a0a53f4f82")
	uuid2, _ := uuid.Parse("0edb6398-fa61-43c9-9ffd-e83127fc6060")
	uuid3, _ := uuid.Parse("cd9d9123-a7cc-48ed-87e1-045b21eaf466")
	uuid4, _ := uuid.Parse("4b481c87-f208-4dfe-bc44-18c631a95a34")
	ids := []uuid.UUID{uuid1, uuid2, uuid3, uuid4}
	usernames := []string{"TuanAnh", "SpacePotato", "TestUsername", "Duke40199"}
	for i := 0; i < len(usernames); i++ {
		accounts = append(accounts,
			models.Account{
				ID:                 ids[i],
				Password:           "password",
				Username:           usernames[i],
				MaxDailyTasksCount: 5,
			})
	}
	SeedAccounts(db, &accounts)
}

//SeedAccounts will seed users to the DB
func SeedAccounts(db *gorm.DB, accounts *[]models.Account) {
	db.Create(&accounts)
	fmt.Println("======= User seeded.")
}

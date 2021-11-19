package conn

import (
	sqlite "manabie/manabie/databases/drivers/sqlite"

	"gorm.io/gorm"
)

// Connect func
func Connect() *gorm.DB {
	return sqlite.Connect()
}
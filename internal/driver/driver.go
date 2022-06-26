package driver

import (
	"fmt"

	"example.com/m/v2/utils"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pkg/errors"
)

var DB *gorm.DB

// Connect to Database postgres
func ConnectDatabase() (*gorm.DB, error) {
	dsn := ""
	switch utils.Env.DBDriver {
	case "postgres":
		dsn = utils.Env.DSNPostgres
	//To Do for another database
	default:
		return nil, errors.New("cannot find db driver")
	}

	database, err := gorm.Open(utils.Env.DBDriver, dsn)
	if err != nil {
		return nil, errors.Wrap(err, "Fail to Connect Database!")
	}

	if err := database.DB().Ping(); err != nil {
		return nil, err
	}

	DB = database
	fmt.Println("Connection DB Complete")

	return database, nil

}

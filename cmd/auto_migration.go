package cmd

import (
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"log"
	"togo/common"
	"togo/handler"
	"togo/pkg/mysql"
)

var autoMigrationCmd = &cobra.Command{
	Use:   "auto-migration",
	Short: "Run auto migration to setup database",
	Run: func(cmd *cobra.Command, args []string) {
		mySql := mysql.NewMySQL(common.MySQL_Key)
		if err := mySql.Run(); err != nil {
			log.Fatal(err.Error())
		}

		db := mySql.Get().(*gorm.DB)
		autoMigration := handler.NewAutoMigration("./migration", db)
		if err := autoMigration.Run(2); err != nil {
			log.Fatal(err.Error())
		}
	},
}

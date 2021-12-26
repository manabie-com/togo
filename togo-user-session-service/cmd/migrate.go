package cmd

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate user db",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		dbType := "sqlite"
		name := "./togo.db"
		switch dbType {
		case "sqlite":
			if err := migrateSQLite(name); err != nil {
				panic(err)
			}
		}

		fmt.Println("migrate success")
	},
}

func migrateSQLite(dbName string) error {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return err
	}

	stm := `
  CREATE TABLE users(
    id TEXT NOT NULL PRIMARY KEY,
    username TEXT UNIQUE,
    password TEXT
  );
  `
	_, err = db.Exec(stm)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

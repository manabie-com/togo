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
	Short: "Migrate DB",
	Long:  `Migrate DB`,
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
  CREATE TABLE tasks(
    id TEXT NOT NULL PRIMARY KEY,
    user_id TEXT,
    title TEXT,
    content TEXT,
    created_time NUMERIC
  );
  CREATE INDEX idx_user_id_created_time ON tasks(user_id ASC, created_time DESC);
  `
	_, err = db.Exec(stm)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.AddCommand(migrateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

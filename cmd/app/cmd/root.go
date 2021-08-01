package cmd

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
	"os"
	"togo/config"
	"togo/router"
)

var rootCmd = &cobra.Command{
	Use:   "Togo",
	Short: "Togo api",
	Long:  "Togo Web api",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := config.New()
		if err != nil {
			return err
		}

		dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", c.DbUsername, c.DbPassword, c.DbHost, c.DbPort, c.DbName, c.SslMode)

		db, err := sql.Open("postgres", dataSourceName)
		if err != nil {
			return err
		}

		if err = db.Ping(); err != nil {
			return err
		}

		r := router.NewRouter(c, db)

		if err := r.Start(); err != nil {
			return err
		}

		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

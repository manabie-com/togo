package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"togo/cmd/internal"
	"togo/config"
	"togo/router"
)

var rootCmd = &cobra.Command{
	Use:   "Togo",
	Short: "Togo api",
	Long:  "Togo Web api",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := config.NewConfig()
		if err != nil {
			return err
		}

		db, err := internal.NewPostgresql(c)
		if err != nil {
			return err
		}

		r := router.NewRouter(c, db)

		err = r.Start()
		if err != nil {
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

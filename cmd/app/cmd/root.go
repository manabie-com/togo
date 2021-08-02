package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"togo/cmd/internal"
	"togo/config"
)

var rootCmd = &cobra.Command{
	Use:   "Togo",
	Short: "Togo api",
	Long:  "Togo Web api",
	RunE: func(cmd *cobra.Command, args []string) error {
		conf, err := config.NewConfig()
		if err != nil {
			return err
		}

		db, err := internal.NewPostgresql(conf)
		if err != nil {
			return err
		}

		rdb, err := internal.NewRedis(conf)
		if err != nil {
			return err
		}

		err = NewServer(&config.ServerConfig{Config: conf, DB: db, Redis: rdb}).Start()
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

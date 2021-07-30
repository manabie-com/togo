package cmd

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
	"os"
	"togo/cmd/router"
	"togo/config"
	"togo/db/postgres"
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
		store, err := postgres.NewStore(dataSourceName)
		if err != nil {
			return err
		}

		serviceContext := config.NewServiceContext(store, c)

		err = router.NewRouter(serviceContext)
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

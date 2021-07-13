package main

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	_ "github.com/lib/pq"

	"github.com/manabie-com/togo/config"
	"github.com/manabie-com/togo/server"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use: "application",
	}
	cmd.AddCommand(&cobra.Command{
		Use: "server",
		Run: func(cmd *cobra.Command, args []string) {
			server.Serve()
		},
	})
	cmd.AddCommand(&cobra.Command{
		Use: "migrate up",
		Run: func(cmd *cobra.Command, args []string) {
			m, err := migrate.New(
				config.MigrationFolder,
				config.PostgreSQL.String())
			if err != nil {
				log.Fatal(err)
			}
			err = m.Up()
			if err != nil && err != migrate.ErrNoChange {
				log.Fatal(err)
			}
		},
	})

	cmd.AddCommand(&cobra.Command{
		Use: "migrate down",
		Run: func(cmd *cobra.Command, args []string) {
			m, err := migrate.New(
				config.MigrationFolder,
				config.PostgreSQL.String())
			if err != nil {
				log.Fatal(err)
			}
			err = m.Down()
			if err != nil && err != migrate.ErrNoChange {
				log.Fatal(err)
			}
		},
	})
	if err := cmd.Execute(); err != nil {
		log.Panic(err)
	}
}

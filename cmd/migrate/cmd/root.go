package cmd

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	"github.com/spf13/cobra"
	"log"
	"os"
	"togo/config"
)

const dialect = "postgres"

var (
	flags = flag.NewFlagSet("migrate", flag.ExitOnError)
	dir   = flags.String("dir", "./migrations", "directory with migration files")
)

var rootCmd = &cobra.Command{
	Use:   "Togo",
	Short: "Togo Migrate",
	Long:  "Togo Migrate",
	Run: func(cmd *cobra.Command, args []string) {
		flags.Usage = usage
		flags.Parse(os.Args[1:])
		if len(args) == 0 || args[0] == "-h" || args[0] == "--help" {
			flags.Usage()
			return
		}

		command := args[0]
		switch command {
		case "create":
			if err := goose.Run("create", nil, *dir, args[1:]...); err != nil {
				log.Fatalf("migrate run: %v", err)
			}
			return
		case "fix":
			if err := goose.Run("fix", nil, *dir); err != nil {
				log.Fatalf("migrate run: %v", err)
			}
			return
		}

		c, err := config.New()
		if err != nil {
			return
		}

		dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", c.DbUsername, c.DbPassword, c.DbHost, c.DbPort, c.DbName, c.SslMode)

		db, err := sql.Open(dialect, dataSourceName)

		if err := goose.SetDialect(dialect); err != nil {
			log.Fatal(err)
		}

		if err := goose.Run(command, db, *dir, args[1:]...); err != nil {
			log.Fatalf("migrate run: %v", err)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func usage() {
	fmt.Println(usagePrefix)
	flags.PrintDefaults()
	fmt.Println(usageCommands)
}

var (
	usagePrefix = `Usage: migrate [OPTIONS] COMMAND
Examples:
    migrate status
Options:
`
	usageCommands = `
Commands:
    up                   Migrate the DB to the most recent version available
    up-by-one            Migrate the DB up by 1
    up-to VERSION        Migrate the DB to a specific VERSION
    down                 Roll back the version by 1
    down-to VERSION      Roll back to a specific VERSION
    redo                 Re-run the latest migration
    reset                Roll back all migrations
    status               Dump the migration status for the current DB
    version              Print the current version of the database
    create NAME [sql|go] Creates new migration file with the current timestamp
    fix                  Apply sequential ordering to migrations
`
)

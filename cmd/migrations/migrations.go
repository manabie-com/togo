package migrations

import (
	"fmt"
	migrateV4 "github.com/golang-migrate/migrate/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

const versionTimeFormat = "20060102150405"

func MigrateCommand(sourceUrl string, databaseUrl string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "database migration command",
	}

	cmd.AddCommand(&cobra.Command{
		Use:   "up",
		Short: "migration up",
		Run: func(cmd *cobra.Command, args []string) {
			m, err := migrateV4.New(sourceUrl, databaseUrl)
			fmt.Println(sourceUrl)
			fmt.Println(databaseUrl)
			if err != nil {
				logrus.Fatal(err)
			}
			logrus.Info("migration up")
			if err := m.Up(); err != nil && err != migrateV4.ErrNoChange {
				logrus.Fatal(err)
			}
		},
	}, &cobra.Command{
		Use:   "down",
		Short: "step down migration by N(int)",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			m, err := migrateV4.New(sourceUrl, databaseUrl)
			if err != nil {
				logrus.Fatal(err)
			}
			down, err := strconv.Atoi(args[0])
			if err != nil {
				logrus.Fatal("rev should be a number", err)
			}
			logrus.Infof("migration down %d", -down)
			if err := m.Steps(-down); err != nil {
				logrus.Fatal(err)
			}
		},
	}, &cobra.Command{
		Use:  "create",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			folder := strings.ReplaceAll(sourceUrl, "file://", "")
			now := time.Now()
			ver := now.Format(versionTimeFormat)
			name := strings.Join(args, "-")

			up := fmt.Sprintf("%s/%s_%s.up.sql", folder, ver, name)
			down := fmt.Sprintf("%s/%s_%s.down.sql", folder, ver, name)

			logrus.Infof("create migration %s", name)
			logrus.Infof("up script %s", up)
			logrus.Infof("down script %s", down)

			if err := ioutil.WriteFile(up, []byte{}, 0644); err != nil {
				logrus.Fatalf("Create migration up error %v", err)
			}
			if err := ioutil.WriteFile(down, []byte{}, 0644); err != nil {
				logrus.Fatalf("Create migration down error %v", err)
			}
		},
	})
	return cmd
}

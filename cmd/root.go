package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use: "main",
	Short: "Start a server",
}

func Execute()  {
	rootCmd.AddCommand(autoMigrationCmd)
	rootCmd.AddCommand(serverCmd)
	rootCmd.Execute()
}

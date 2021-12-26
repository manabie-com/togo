package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "togo",
	Short: "Manabie togo app",
	Long:  "Manabie togo app implement",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	viper.AutomaticEnv()

	cfgFile := os.Getenv("CONFIG_FILE")
	viper.SetConfigFile(cfgFile)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading config file: %s\n", cfgFile)
		fmt.Fprintf(os.Stderr, "%s\n", err)
	} else {
		fmt.Println("Using config file:", cfgFile)
	}
}

package main

import (
	"fmt"
	"github.com/banhquocdanh/togo/internal/commands"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	rootCmd, err := commands.ToDoCommand()
	if err != nil {
		panic(err)
	}

	rootCmd.HelpFunc()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("rootCmd:", err)
	}

}

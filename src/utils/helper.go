package utils

import (
	"os"
)

/*
Write all data into a file as recording
 */

func WriteFile(filePath string, data string) {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err = f.Write([]byte(data)); err != nil {
		panic(err)
	}
	if _, err = f.WriteString("\n"); err != nil {
		panic(err)
	}
}

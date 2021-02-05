package utils

import (
	"log"
	"os"
)

var (
	_info  = log.New(os.Stderr, "INFO: ", log.LstdFlags)
	_error = log.New(os.Stderr, "ERROR: ", log.LstdFlags)
)

func Println(message string) {
	_info.Println(message)
}

func Printf(format string, v ...interface{}) {
	_info.Printf(format, v...)
}

func Error(message string) {
	_error.Println(message)
}

func Errorf(format string, v ...interface{}) {
	_error.Printf(format, v...)
}

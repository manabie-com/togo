package utils

//import (
//	"log"
//	"os"
//)
//
//var DebugEnabled bool
//
//const (
//	ERROR   = "ERROR"
//	WARNING = "WARNING"
//	INFO    = "INFO"
//	DEBUG   = "DEBUG"
//)
//
//type Log struct {
//	LogMessage string
//	Message    string
//	ExitCode   int
//}
//
//func Error(l Log) {
//	log.Printf("%s: %v", ERROR, l.LogMessage)
//	if DebugEnabled {
//		if l.Message != "" {
//			log.Printf("%s: %v", DEBUG, l.Message)
//		}
//	}
//	os.Exit(1)
//}
//
//func Warning(l Log) {
//	log.Printf("%s: %v", WARNING, l.LogMessage)
//	if DebugEnabled {
//		if l.Message != "" {
//			log.Printf("%s: %v", DEBUG, l.Message)
//		}
//	}
//}
//
//func Info(l Log) {
//	log.Printf("%s: %v", INFO, l.LogMessage)
//	if DebugEnabled {
//		if l.Message != "" {
//			log.Printf("%s: %v", DEBUG, l.Message)
//		}
//	}
//}
//
//func Debug(l Log) {
//	log.Printf("%s: %v", DEBUG, l.LogMessage)
//	log.Printf("%s: %v", DEBUG, l.Message)
//}

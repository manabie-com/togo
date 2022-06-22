package controllers

import (
	"log"
	"net/http"
)

func ErrorHandle(w http.ResponseWriter,err error, message string, status int) {
	if err != nil {
		http.Error(w, message, status)
		log.Fatal("asdfsafsfsaf")
	}
}
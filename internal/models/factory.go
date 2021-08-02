package models

import (
	"log"
)

var modelList = []interface{}{
	//Account-related
	Account{},
	Task{},
}

func init() {
	log.Println("Initializing Models Factory")
}

//GetModelList will get all models
func GetModelList() []interface{} {
	return modelList
}

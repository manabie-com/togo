package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ReturnJSON(w http.ResponseWriter, data interface{}) {
	response, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(500)
		fmt.Errorf("Cannot encode to JSON : %s ", err.Error())
	}
	w.Write(response)
}

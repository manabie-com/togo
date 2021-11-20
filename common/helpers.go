package common

import (
	"encoding/json"
	"fmt"
)

func Interface2String(i interface{}) map[string]string {
	var p map[string]interface{}
	var result = map[string]string{}
	data, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &p)
	if err != nil {
		panic(err)
	}
	for k, v := range p {
		if k == "password" {
			result[k] = "secret token"
		} else {
			result[k] = fmt.Sprintf("%v", v)
		}
	}
	return result
}

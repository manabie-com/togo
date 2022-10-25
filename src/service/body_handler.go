package service

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

//JsonBody handle data of body in request

func JsonBody(g *gin.Context) (error, map[string]interface{}) {
	value, err := ioutil.ReadAll(g.Request.Body)
	g.Request.Body = ioutil.NopCloser(bytes.NewBuffer(value))
	if err != nil {
		return err, nil
	}
	var data map[string]interface{}
	json.Unmarshal(value, &data)
	return nil, data
}

package utils

import (
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"net/http"
	"reflect"
	"strings"
)

func InitHTTPRequest(method, url string, headers map[string]string, params map[string]string, content interface{}) *http.Request {
	// init request URL
	if params != nil && len(params) > 0 {
		url += "?"
		for k := range params {
			url += fmt.Sprintf("%s=%s&", k, params[k])
		}

		// Remove last '&' in URL
		url = url[:len(url)]
	}

	// init body
	var body io.Reader
	if content != nil {
		b, err := json.Marshal(content)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		body = strings.NewReader(string(b))
	}

	// setup new request
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	// set headers
	if headers != nil {
		for k := range headers {
			req.Header.Set(k, headers[k])
		}
	}

	return req
}

func ToMap(obj interface{}) map[string]string {
	var output map[string]string

	if obj == nil {
		return nil
	}

	b, err := json.Marshal(obj)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	err = json.Unmarshal(b, &output)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return output
}

func IsEqual(obj1, obj2 interface{}) bool {
	return reflect.DeepEqual(obj1, obj2)
}
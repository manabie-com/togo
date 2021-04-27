package test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func auth(user string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodPost, "http://localhost:5050/login", strings.NewReader(user))
	if err != nil {
		return nil, err
	}
	return req, nil
}
func CreateToken(user string) (string, int, error) {
	req, err := auth(user)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}
	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}
	var data string
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}
	return data, http.StatusOK, nil
}

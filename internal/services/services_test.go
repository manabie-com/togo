package services

import (
	"net/http"
	"testing"
)

func TestLogin(t *testing.T) {
	const correctCredUri = "http://localhost:5050/login?user_id=firstUser&password=example"
	const wrongCredUri = "http://localhost:5050/login?user_id=shadow&password=null"
	resp, err := http.Get(correctCredUri)
	if err != nil {
		t.Fatal("Case 1: Unable to make request")
		t.Fatal(err.Error())
		t.FailNow()
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatal("Wrong authentication")
		t.FailNow()
	}

	resp, err = http.Get(wrongCredUri)
	if err != nil {
		t.Fatal("Case 2: Unable to make request")
		t.Fatal(err.Error())
		t.FailNow()
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusUnauthorized {
		t.Fatal("Wrong status code/ authentication")
		t.FailNow()
	}

}

package integration_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"github.com/manabie-com/togo/internal/services"
	sqllite "github.com/manabie-com/togo/internal/storages/sqlite"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
)

func init_db_test() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./data_test.db")
	if err != nil {
		log.Fatal("error opening db", err)
		return nil, err
	}
	return db, nil
}

func init_test_server(db *sql.DB) *httptest.Server {
	sv := &services.ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store: &sqllite.LiteDB{
			DB: db,
		}}
	ts := httptest.NewServer(sv)
	return ts
}

func getData(ts *httptest.Server, url string, header map[string]string) (int, map[string]interface{}, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", ts.URL+url, nil)
	if err != nil {
		return 0, nil, err
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}
	res, err := client.Do(req)
	if err != nil {
		return 0, nil, err
	}
	resBody, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return 0, nil, err
	}

	var rs map[string]interface{}
	json.Unmarshal([]byte(resBody), &rs)
	//fmt.Printf("%s\n-h %s\nRes>> code: %d > body: %s\n",req.URL, req.Header, res.StatusCode, rs)

	return res.StatusCode, rs, err
}

func postData(ts *httptest.Server, url string, body []byte, header map[string]string) (int, map[string]interface{}, error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", ts.URL+url, bytes.NewBuffer(body))
	if err != nil {
		return 0, nil, err
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}
	res, err := client.Do(req)
	if err != nil {
		return 0, nil, err
	}
	resBody, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return 0, nil, err
	}

	var rs map[string]interface{}
	json.Unmarshal([]byte(resBody), &rs)
	//fmt.Printf("%s\n-body: %s\n-h: %s\nRes>> code: %d > resBody: %s\n",
	//	req.URL, string(body), req.Header, res.StatusCode, rs)

	return res.StatusCode, rs, err
}

func clearData(db *sql.DB) error {
	stmt := `DELETE FROM tasks`
	_, err := db.Exec(stmt)
	if err != nil {
		return err
	}

	stmt = `DELETE FROM users`
	_, err = db.Exec(stmt)
	if err != nil {
		return err
	}

	return nil
}

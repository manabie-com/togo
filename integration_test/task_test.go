package integration_test

import (
	"database/sql"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"net/http/httptest"
	"testing"
)

type DataTask struct {
	id      string
	content string
	user_id string
	date    string
}

func TestGetTask(t *testing.T) {
	//prepare
	db, err := init_db_test()
	if err != nil {
		t.Skip("error init db")
		return
	}
	clearData(db)
	ts := init_test_server(db)
	defer ts.Close()
	_, tk_1, tk_2, errInit := initDataGetTask(db, ts)
	if errInit != nil {
		t.Skip("error init data")
		return
	}

	cases := []struct {
		caseId string
		tk     string
		url    string
		code   int
		user   string
		rs     int
	}{
		{
			caseId: "get_task_1",
			tk:     "abcefd1kad",
			url:    "/tasks?created_date=2020-06-29",
			code:   http.StatusUnauthorized,
			user:   "",
			rs:     0,
		},
		{
			caseId: "get_task_2",
			tk:     tk_1,
			url:    "/tasks?created_date=2020-06-29",
			code:   http.StatusOK,
			user:   "user_1",
			rs:     2,
		},
		{
			caseId: "get_task_3",
			tk:     tk_1,
			url:    "/tasks?created_date=2020-06-25",
			code:   http.StatusOK,
			user:   "user_1",
			rs:     1,
		},
		{
			caseId: "get_task_4",
			tk:     tk_2,
			url:    "/tasks?created_date=2020-06-29",
			code:   http.StatusOK,
			user:   "user_2",
			rs:     0,
		},
	}
	for _, c := range cases {
		header := map[string]string{
			"Authorization": c.tk,
		}
		code, _, err := getData(ts, c.url, header)
		if err != nil {
			t.Errorf("case fail - %s\nRequest fail: %s", c.caseId, err)
		}
		if code != c.code {
			//t.Logf("wrong status code %d - true result %d", code, c.code)
			t.Errorf("case fail - %s\nwrong status code %d >< true result %d", c.caseId, code, c.code)
		} else {
			if code != 200 {
				t.Logf("case pass - %s", c.caseId)
			} else {

			}
		}
	}
	//
	//header = map[string]string{
	//	"Authorization": tk_2,
	//}
	//_, _, err = getData(ts, "/tasks?created_date="+data[0].date, header)
	//if err != nil {
	//	t.Error(err)
	//}
}

func TestaddTask(t *testing.T) {
	//prepare
	db, err := init_db_test()
	if err != nil {
		t.Skip("error init db")
		return
	}
	clearData(db)
	ts := init_test_server(db)
	defer ts.Close()
	//cases, tk_1, tk_2, errInit := initDataGetTask(db, ts)
	//if errInit != nil {
	//	t.Skip("error init data")
	//	return
	//}
	//
	//header := map[string]string{
	//	"Authorization": token,
	//}
	//js := []byte(`{"content":"abcdfkjwe cj fkajr gk drg "}`)
	//_, _, _ = postData(ts, "/tasks", js, header)
	//if "a" != "want" {
	//	t.Errorf("Hello() = %q, want %q", "a", "want")
	//}
}

func initDataGetTask(db *sql.DB, ts *httptest.Server) ([]DataTask, string, string, error) {
	stmt := `INSERT INTO tasks (id, content, user_id, created_date) VALUES (?, ?, ?, ?)`

	dataSuite := []DataTask{
		{
			content: "content a from user_1",
			user_id: "user_1",
			date:    "2020-06-29",
		},
		{
			content: "content b from user_1",
			user_id: "user_1",
			date:    "2020-06-29",
		},
		{
			content: "content c from user_1",
			user_id: "user_1",
			date:    "2020-06-25",
		},
		{
			content: "content a from user_2",
			user_id: "user_2",
			date:    "2020-06-29",
		},
		{
			content: "content d from user_2",
			user_id: "user_2",
			date:    "2020-06-29",
		},
		{
			content: "content f from user_2",
			user_id: "user_2",
			date:    "2020-06-29",
		},
	}
	for _, c := range dataSuite {
		_, err := db.Exec(stmt, uuid.New().String(), c.content, c.user_id, c.date)
		if err != nil {
			return nil, "", "", err
		}
	}

	initDataUser(db, "user_1", "123", 5)
	initDataUser(db, "user_2", "123", 5)
	tk_1, tk_2 := "", ""
	code, rsl, er := getData(ts, "/login?user_id=user_1&password=123", nil)
	if code != http.StatusOK || er != nil {
		return nil, "", "", er
	} else {
		tk_1 = rsl["data"].(string)
	}

	code, rsl, er = getData(ts, "/login?user_id=user_2&password=123", nil)
	if code != http.StatusOK || er != nil {
		return nil, "", "", er
	} else {
		tk_2 = rsl["data"].(string)
	}

	return dataSuite, tk_1, tk_2, nil
}

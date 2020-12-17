package integration_test

import (
	"database/sql"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/storages"
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
	_, tk_1, tk_2, errInit := initDataGetTask(db, "postgres", ts)
	if errInit != nil {
		t.Skipf("error init data: %s", errInit)
		return
	}

	cases := []struct {
		caseId string
		tk     string
		url    string
		code   int
		user   string
		date   string
		rs     int
	}{
		{
			caseId: "get_task_1",
			tk:     "abcefd1kad",
			url:    "/tasks?created_date=2020-06-29",
			code:   http.StatusUnauthorized,
			user:   "",
			date:   "2020-06-29",
			rs:     0,
		},
		{
			caseId: "get_task_2",
			tk:     tk_1,
			url:    "/tasks?created_date=2020-06-29",
			code:   http.StatusOK,
			user:   "user_1",
			date:   "2020-06-29",
			rs:     2,
		},
		{
			caseId: "get_task_3",
			tk:     tk_1,
			url:    "/tasks?created_date=2020-06-25",
			code:   http.StatusOK,
			user:   "user_1",
			date:   "2020-06-25",
			rs:     1,
		},
		{
			caseId: "get_task_4",
			tk:     tk_2,
			url:    "/tasks?created_date=2020-06-29",
			code:   http.StatusOK,
			user:   "user_2",
			date:   "2020-06-29",
			rs:     3,
		},
		{
			caseId: "get_task_5",
			tk:     tk_2,
			url:    "/tasks?created_date=2020-06-01",
			code:   http.StatusOK,
			user:   "user_2",
			date:   "2020-06-29",
			rs:     0,
		},
	}
	for _, c := range cases {
		header := map[string]string{
			"Authorization": c.tk,
		}
		code, rs, err := getData(ts, c.url, header)
		if err != nil {
			t.Errorf("case fail - %s\nRequest fail: %s", c.caseId, err)
			continue
		}
		if code != c.code {
			t.Errorf("case fail - %s\nwrong status code %d >< true result %d", c.caseId, code, c.code)
			continue
		} else {
			if code != 200 {
				t.Logf("case pass - %s", c.caseId)
				continue
			} else {
				type ArrayTask []storages.Task
				_, exist := rs["data"]
				if !exist {
					t.Errorf("case fail - %s\nwrong schema result", c.caseId)
					continue
				}
				data, _ := json.Marshal(rs["data"])
				ars := ArrayTask{}
				json.Unmarshal(data, &ars)

				if len(ars) != c.rs {
					t.Errorf("case fail - %s\ngot %d item(s) >< true result %d item(s)", c.caseId, len(rs), c.rs)
					continue
				} else {
					for _, d := range ars {
						if d.UserID != c.user || d.CreatedDate != c.date {
							t.Errorf("case fail - wrong item")
							continue
						}
					}
					t.Logf("case pass - %s", c.caseId)
				}
			}
		}
	}
}

func initDataGetTask(db *sql.DB, driveName string, ts *httptest.Server) ([]DataTask, string, string, error) {
	tk_1, tk_2, err := initUserTest(db, ts)
	if err != nil {
		return nil, "", "", err
	}

	stmt := `INSERT INTO tasks (id, content, user_id, created_date) VALUES (?, ?, ?, ?)`
	if driveName == "postgres" {
		stmt = `INSERT INTO tasks (id, content, user_id, created_date) VALUES ($1, $2, $3, $4)`
	}

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

	return dataSuite, tk_1, tk_2, nil
}

func initUserTest(db *sql.DB, ts *httptest.Server) (string, string, error) {
	initDataUser(db, "postgres", "user_1", "123", 5)
	initDataUser(db, "postgres", "user_2", "123", 5)
	tk_1, tk_2 := "", ""
	code, rsl, er := getData(ts, "/login?user_id=user_1&password=123", nil)
	if code != http.StatusOK || er != nil {
		return "", "", er
	} else {
		tk_1 = rsl["data"].(string)
	}

	code, rsl, er = getData(ts, "/login?user_id=user_2&password=123", nil)
	if code != http.StatusOK || er != nil {
		return "", "", er
	} else {
		tk_2 = rsl["data"].(string)
	}
	return tk_1, tk_2, nil
}

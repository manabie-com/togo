package integration_test

import (
	"database/sql"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/storages"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestAddFirstTask(t *testing.T) {
	//prepare
	db, err := init_db_test()
	if err != nil {
		t.Skip("error init db")
		return
	}
	clearData(db)
	ts := init_test_server(db)
	defer ts.Close()

	tk_1, tk_2, errInit := initUserTest(db, ts)
	if errInit != nil {
		t.Skip("error init data")
		return
	}

	cases := []struct {
		caseId  string
		tk      string
		url     string
		body    string
		content string
		code    int
		user    string
		rs      int
	}{
		{
			caseId:  "add_task_first_1",
			tk:      "abcefd1kad",
			url:     "/tasks",
			body:    `{"content":"con tent task test unauthor"}`,
			content: "con tent task test unauthor",
			code:    http.StatusUnauthorized,
			user:    "",
			rs:      0,
		},
		{
			caseId:  "add_task_first_2",
			tk:      tk_1,
			url:     "/tasks",
			body:    `{"content":"content test us1 success"}`,
			content: "content test us1 success",
			code:    http.StatusOK,
			user:    "user_1",
			rs:      0,
		},
		{
			caseId:  "add_task_first_3",
			tk:      tk_2,
			url:     "/tasks",
			body:    `{"content":"content test us2 success"}`,
			content: "content test us2 success",
			code:    http.StatusOK,
			user:    "user_2",
			rs:      0,
		},
	}
	for _, c := range cases {
		header := map[string]string{
			"Authorization": c.tk,
		}
		code, rs, err := postData(ts, c.url, []byte(c.body), header)
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
				_, exist := rs["data"]
				if !exist {
					t.Errorf("case fail - %s\nwrong schema result", c.caseId)
					continue
				}
				data, _ := json.Marshal(rs["data"])
				d := storages.Task{}
				json.Unmarshal(data, &d)
				if d.UserID != c.user || d.Content != c.content {
					t.Errorf("case fail - wrong item return")
					continue
				} else {
					if checkTaskExist(db, d) {
						t.Logf("case pass - %s", c.caseId)
						continue
					} else {
						t.Errorf("case fail - %s\nnot found in db task_id: %s", c.caseId, d.ID)
						continue
					}
				}
			}
		}
	}
}

func TestAddTask(t *testing.T) {
	//prepare
	db, err := init_db_test()
	if err != nil {
		t.Skip("error init db")
		return
	}
	clearData(db)
	ts := init_test_server(db)
	defer ts.Close()
	_, tk_1, tk_2, errInit := initDataAddTask(db, ts)
	if errInit != nil {
		t.Skip("error init data")
		return
	}

	cases := []struct {
		caseId  string
		tk      string
		url     string
		body    string
		content string
		code    int
		user    string
		rs      int
	}{
		{
			caseId:  "add_task_1",
			tk:      "abcefd1kad",
			url:     "/tasks",
			body:    `{"content":"con tent task test unauthor"}`,
			content: "con tent task test unauthor",
			code:    http.StatusUnauthorized,
			user:    "",
			rs:      0,
		},
		{
			caseId:  "add_task_2",
			tk:      tk_1,
			url:     "/tasks",
			body:    `{"content":"content test us1 success"}`,
			content: "content test us1 success",
			code:    http.StatusOK,
			user:    "user_1",
			rs:      0,
		},
		{
			caseId:  "add_task_3",
			tk:      tk_2,
			url:     "/tasks",
			body:    `{"content":"content test us2 success"}`,
			content: "content test us2 success",
			code:    http.StatusOK,
			user:    "user_2",
			rs:      0,
		},
	}
	for _, c := range cases {
		header := map[string]string{
			"Authorization": c.tk,
		}
		code, rs, err := postData(ts, c.url, []byte(c.body), header)
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
				_, exist := rs["data"]
				if !exist {
					t.Errorf("case fail - %s\nwrong schema result", c.caseId)
					continue
				}
				data, _ := json.Marshal(rs["data"])
				d := storages.Task{}
				json.Unmarshal(data, &d)
				if d.UserID != c.user || d.Content != c.content {
					t.Errorf("case fail - wrong item return")
					continue
				} else {
					if checkTaskExist(db, d) {
						t.Logf("case pass - %s", c.caseId)
						continue
					} else {
						t.Errorf("case fail - %s\nnot found in db task_id: %s", c.caseId, d.ID)
						continue
					}
				}
			}
		}
	}
}

func TestAddTaskReachLimit(t *testing.T) {
	//prepare
	db, err := init_db_test()
	if err != nil {
		t.Skip("error init db")
		return
	}
	clearData(db)
	ts := init_test_server(db)
	defer ts.Close()
	_, tk_1, tk_2, errInit := initDataAddTaskLimit(db, ts)
	if errInit != nil {
		t.Skip("error init data")
		return
	}

	cases := []struct {
		caseId  string
		tk      string
		url     string
		body    string
		content string
		code    int
		user    string
		rs      int
	}{
		{
			caseId:  "add_task_limit_1",
			tk:      tk_1,
			url:     "/tasks",
			body:    `{"content":"content task test limit reach"}`,
			content: "content task test limit reach",
			code:    http.StatusBadRequest,
			user:    "",
			rs:      0,
		},
		{
			caseId:  "add_task_limit_2",
			tk:      tk_2,
			url:     "/tasks",
			body:    `{"content":"content test us2 success"}`,
			content: "content test us2 success",
			code:    http.StatusOK,
			user:    "user_2",
			rs:      0,
		},
	}
	for _, c := range cases {
		header := map[string]string{
			"Authorization": c.tk,
		}
		code, rs, err := postData(ts, c.url, []byte(c.body), header)
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
				_, exist := rs["data"]
				if !exist {
					t.Errorf("case fail - %s\nwrong schema result", c.caseId)
					continue
				}
				data, _ := json.Marshal(rs["data"])
				d := storages.Task{}
				json.Unmarshal(data, &d)
				if d.UserID != c.user || d.Content != c.content {
					t.Errorf("case fail - wrong item return")
					continue
				} else {
					if checkTaskExist(db, d) {
						t.Logf("case pass - %s", c.caseId)
						continue
					} else {
						t.Errorf("case fail - %s\nnot found in db task_id: %s", c.caseId, d.ID)
						continue
					}
				}
			}
		}
	}
}

func initDataAddTask(db *sql.DB, ts *httptest.Server) ([]DataTask, string, string, error) {
	stmt := `INSERT INTO tasks (id, content, user_id, created_date) VALUES (?, ?, ?, ?)`
	today := time.Now().Format("2006-01-02")
	dataSuite := []DataTask{
		{
			content: "content a from user_1",
			user_id: "user_1",
			date:    today,
		},
		{
			content: "content b from user_1",
			user_id: "user_1",
			date:    today,
		},
		{
			content: "content c from user_1",
			user_id: "user_1",
			date:    today,
		},
		{
			content: "content a from user_2",
			user_id: "user_2",
			date:    today,
		},
		{
			content: "content d from user_2",
			user_id: "user_2",
			date:    today,
		},
	}
	for _, c := range dataSuite {
		_, err := db.Exec(stmt, uuid.New().String(), c.content, c.user_id, c.date)
		if err != nil {
			return nil, "", "", err
		}
	}

	tk_1, tk_2, err := initUserTest(db, ts)
	if err != nil {
		return nil, "", "", err
	}

	return dataSuite, tk_1, tk_2, nil
}

func initDataAddTaskLimit(db *sql.DB, ts *httptest.Server) ([]DataTask, string, string, error) {
	stmt := `INSERT INTO tasks (id, content, user_id, created_date) VALUES (?, ?, ?, ?)`
	today := time.Now().Format("2006-01-02")
	dataSuite := []DataTask{
		{
			content: "content a from user_1",
			user_id: "user_1",
			date:    today,
		},
		{
			content: "content b from user_1",
			user_id: "user_1",
			date:    today,
		},
		{
			content: "content c from user_1",
			user_id: "user_1",
			date:    today,
		},
		{
			content: "content x from user_1",
			user_id: "user_1",
			date:    today,
		},
		{
			content: "content y from user_1",
			user_id: "user_1",
			date:    today,
		},
		{
			content: "content a from user_2",
			user_id: "user_2",
			date:    today,
		},
		{
			content: "content d from user_2",
			user_id: "user_2",
			date:    today,
		},
		{
			content: "content e from user_2",
			user_id: "user_2",
			date:    today,
		},
		{
			content: "content f from user_2",
			user_id: "user_2",
			date:    today,
		},
		{
			content: "content g from user_2",
			user_id: "user_2",
			date:    "2020-12-10",
		},
	}
	for _, c := range dataSuite {
		_, err := db.Exec(stmt, uuid.New().String(), c.content, c.user_id, c.date)
		if err != nil {
			return nil, "", "", err
		}
	}

	tk_1, tk_2, err := initUserTest(db, ts)
	if err != nil {
		return nil, "", "", err
	}

	return dataSuite, tk_1, tk_2, nil
}

func checkTaskExist(db *sql.DB, task storages.Task) bool {
	stmt := `SELECT count(*) FROM tasks WHERE user_id = ? AND created_date = ? AND content = ? AND id= ?`
	row := db.QueryRow(stmt, task.UserID, task.CreatedDate, task.Content, task.ID)

	var countTask int
	err := row.Scan(&countTask)
	if err == nil {
		if countTask == 1 {
			return true
		}
	}

	return false
}

package integration_test

import (
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"testing"
)

func TestLogin(t *testing.T) {
	//prepare condition
	id := "user_1"
	pass := "123"
	wrongId := "abc"
	wrongPass := "adsdf"
	db, err := init_db_test()
	if err != nil {
		t.Skip("error init db")
		return
	}
	clearData(db)
	errInit := initDataUser(db, id, pass, 5)
	if errInit != nil {
		t.Skip("error init data")
		return
	}
	ts := init_test_server(db)
	defer ts.Close()

	//execute test
	cases := []struct {
		caseId string
		url    string
		code   int
	}{
		{
			caseId: "login_1",
			url:    "/login?user_id=" + id + "&password=" + pass,
			code:   http.StatusOK,
		},
		{
			caseId: "login_2",
			url:    "/login?user_id=" + wrongId + "&password=" + pass,
			code:   http.StatusUnauthorized,
		},
		{
			caseId: "login_3",
			url:    "/login?user_id=" + id + "&password=" + wrongPass,
			code:   http.StatusUnauthorized,
		},
		{
			caseId: "login_4",
			url:    "/login?user_id=" + id,
			code:   http.StatusUnauthorized,
		},
		{
			caseId: "login_5",
			url:    "/login?user_id=" + wrongId,
			code:   http.StatusUnauthorized,
		},
		{
			caseId: "login_6",
			url:    "/login?password=" + pass,
			code:   http.StatusUnauthorized,
		}}
	for _, c := range cases {
		code, rs, err := getData(ts, c.url, nil)
		if err != nil {
			t.Errorf("case fail - %s\nRequest fail: %s", c.caseId, err)
			continue
		}
		if code != c.code {
			//t.Logf("wrong status code %d - true result %d", code, c.code)
			t.Errorf("case fail - %s\nwrong status code %d >< true result %d", c.caseId, code, c.code)
			continue
		} else {
			if code != 200 {
				t.Logf("case pass - %s", c.caseId)
				continue
			} else {
				_, exist := rs["data"]
				if !exist {
					t.Errorf("case fail - wrong schema result")
					continue
				}
				tokenString := rs["data"].(string)
				claims := jwt.MapClaims{}
				jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
					return []byte(""), nil
				})
				if claims["user_id"] != id {
					t.Errorf("case fail - %s\nwrong token info encode: user_id %s >< true result %s", c.caseId, claims["user_id"], id)
					continue
				} else {
					t.Logf("case pass - %s", c.caseId)
				}
			}
		}
	}
}

func initDataUser(db *sql.DB, id, pass string, maxTodo int) error {
	stmt := `INSERT INTO users (id, password, max_todo) VALUES (?, ?, ?)`
	_, err := db.Exec(stmt, id, pass, maxTodo)
	if err != nil {
		return err
	}
	return nil
}

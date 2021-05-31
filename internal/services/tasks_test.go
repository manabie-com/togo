package services

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	postgres "github.com/manabie-com/togo/internal/storages/postgres"
)

func mock(t *testing.T) *ToDoService {
	db := &postgres.DataBase{}
	err := db.Init("todotest", "postgres", "postgres")
	if err != nil {
		t.Fatal(err)
	}
	return &ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store: db,
	}
}

func resetDB(t *testing.T) {
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "postgres", "postgres", "todotest")
	db, err := sql.Open("postgres", psqlConn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}
	stmt := `DELETE FROM users`
	_, err = db.Exec(stmt)
	if err != nil {
		t.Fatal(err)
	}
	stmt = `DELETE FROM tasks`
	_, err = db.Exec(stmt)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRegister(t *testing.T) {
	form := url.Values{}
	form.Add("user_id", "testuser")
	form.Add("password", "123456")
	req, err := http.NewRequest(http.MethodPost, "/register", strings.NewReader(form.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	service := mock(t)
	defer service.Store.Finalize()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(service.register)
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("register returned wrong status code: got %v, want %v", rr.Code, http.StatusOK)
	}
	resetDB(t)
}

func TestRegisterEmptyUser(t *testing.T) {
	form := url.Values{}
	form.Add("user_id", "")
	form.Add("password", "123456")
	req, err := http.NewRequest(http.MethodPost, "/register", strings.NewReader(form.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	service := mock(t)
	defer service.Store.Finalize()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(service.register)
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusBadRequest {
		t.Fatalf("register returned wrong status code: got %v, want %v", rr.Code, http.StatusBadRequest)
	}
}

func TestRegisterUserExisted(t *testing.T) {
	form := url.Values{}
	form.Add("user_id", "testuser")
	form.Add("password", "123456")
	req, err := http.NewRequest(http.MethodPost, "/register", strings.NewReader(form.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	service := mock(t)
	defer service.Store.Finalize()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(service.register)
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("register returned wrong status code: got %v, want %v", rr.Code, http.StatusOK)
	}
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusBadRequest {
		t.Fatalf("register returned wrong status code: got %v, want %v", rr.Code, http.StatusBadRequest)
	}
	resetDB(t)
}

func TestGetAuthToken(t *testing.T) {
	//TODO: Same at above
}

func TestGetAuthTokenUserIncorrect(t *testing.T) {
	//TODO: Same at above
}

func TestAddTask(t *testing.T) {
	//TODO: Same at above
}

func TestAddTaskReachLimit(t *testing.T) {
	//TODO: Same at above
}

func TestUpdateTask(t *testing.T) {
	//TODO: Same at above
}

func TestUpdateTaskReachLimit(t *testing.T) {
	//TODO: Same at above
}

func TestUpdateTaskIDEmpty(t *testing.T) {
	//TODO: Same at above
}

func TestUpdateTaskWrongStatus(t *testing.T) {
	//TODO: Same at above
}

func TestListTask(t *testing.T) {
	//TODO: Same at above
}

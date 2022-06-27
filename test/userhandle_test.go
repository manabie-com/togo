package test

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/huynhhuuloc129/todo/controllers"
)

func TestResponse(t *testing.T) {
	db, _ := NewMock()
	h := controllers.NewBaseHandler(db)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "localhost:8000/users", nil)

	h.ResponseAllUser(w, req)
	resp := w.Result()
	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Errorf("Can't read body response")
	}
	fmt.Print(string(respBody))
	t.Errorf("asdfffff")
}

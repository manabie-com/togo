package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/manabie-com/togo/controllers"
)

func TestAddTodoTask(t *testing.T) {

	r := gin.Default()

	params := map[string]interface{}{
		"task":   "intergrate test",
		"userid": 1,
	}

	body, _ := json.Marshal(params)

	r.POST("/todo/add", controllers.AddTodoTask)

	req, _ := http.NewRequest("POST", "/todo/add", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/raw")

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {

		statusOK := w.Code == http.StatusOK
		return statusOK
	})
}

func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {

	// Create a response recorder
	w := httptest.NewRecorder()

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}

package test

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/huynhhuuloc129/todo/controllers"
// )

// func TestGetOneUser(t *testing.T) {
// 	req, err := http.NewRequest("GET", "/users/1", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	req.Header.Set("token","eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTYwODIwMjEsImlkIjoxMywidXNlcm5hbWUiOiJBRE1JTiJ9.yH-0Krzyu92cBIPy_TxMYqSIA3eso8bKRFW3yFhGKRI")
// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(controllers.ResponeAllUser)
// 	handler.ServeHTTP(rr, req)
// 	if status:= rr.Code; status != http.StatusOK {
// 		t.Errorf("Handler return wrong status code: got %v want %v",status, http.StatusOK);
// 	}
// }
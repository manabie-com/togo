package middlewares

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/huynhhuuloc129/todo/middlewares"
// )
// type loggingHandler struct{
// 	w http.ResponseWriter
// 	r *http.Request
// }
// // func TestLogging(t *testing.T) {
// // 	loggingHandler := loggingHandler{}

// // 	req := httptest.NewRequest(http.MethodGet, "localhost:8000/users/", nil)
// // 	req.Header.Add("token", Token)

// // 	w := httptest.NewRecorder()
// // 	logging := middlewares.Logging(loggingHandler)
// // 	logging.ServeHTTP(w, req)
// // }

// // func (l loggingHandler)ServeHTTP(w http.ResponseWriter, r *http.Request){

// // }
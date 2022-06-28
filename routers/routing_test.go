package routers

// import (
// 	"bytes"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"net/http/httptest"
// 	"strconv"
// 	"testing"

// 	"github.com/gorilla/mux"
// )

// func TestRouting(t *testing.T) {
// 	r := mux.NewRouter()
// 	srv := httptest.NewServer(r)
// 	defer srv.Close()

// 	res, err := http.Get(srv.URL+"/users")
// 	if err != nil {
// 		t.Fatal("Coult not send Get request, err: "+err.Error())
// 	}
// 	if res.StatusCode != http.StatusOK{
// 		t.Errorf("expected status OK; got %v", res.Status)
// 	}

// 	b, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		t.Fatalf("Could not read response: %v", err)
// 	}

// 	d, err := strconv.Atoi(string(bytes.TrimSpace(b)))
// 	if err != nil {
// 		t.Fatalf("expected an integer; got %s", b)
// 	}
// 	fmt.Println(d)
// }

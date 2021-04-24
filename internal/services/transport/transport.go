package transport

import "net/http"

type Transport interface {
	Login(resp http.ResponseWriter, req *http.Request)
	ListTasks(resp http.ResponseWriter, req *http.Request)
	AddTask(resp http.ResponseWriter, req *http.Request)
}

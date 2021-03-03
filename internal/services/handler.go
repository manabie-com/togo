package services

import "net/http"

func methodOptionsHandler(resp http.ResponseWriter, req *http.Request){
	resp.WriteHeader(http.StatusOK)
	return
}

func (s *ToDoService) loginHandler(resp http.ResponseWriter, req *http.Request) {
	s.getAuthToken(resp, req)
	return
}

func (s *ToDoService) getTasksHandler(resp http.ResponseWriter, req *http.Request) {
	s.listTasks(resp, req)
}

func (s *ToDoService) addTaskHandler(resp http.ResponseWriter, req *http.Request) {
	s.addTask(resp, req)
}


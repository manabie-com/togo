package api

import (
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	Initialize()
	json.NewDecoder(r.Body).Decode(&task)

	un, err := task.GetUser()
	if err != nil {
		http.Error(w, err.Error(), 404)
		log.Println("Error: ", err)
		return
	}
	tc := task.GetTasksCount(int64(un.ID))
	// check if user has already exceeded max daily tasks limit
	if tc < int64(un.DailyTaskLimit) {
		task.CreateOneTask(int64(un.ID))
		json.NewEncoder(w).Encode(&task)
	} else {
		http.Error(w, "You have reached the maximum limit of tasks today!", http.StatusUnprocessableEntity)
		return
	}
}

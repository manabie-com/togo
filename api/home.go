package api

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "task limiter service is running!")
}

func RespondError(w http.ResponseWriter, httpStatusCode int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	w.Write([]byte(err.Error()))
	return
}

func RespondWithJSON(w http.ResponseWriter, httpStatusCode int, data interface{}) {
	resp, err := json.Marshal(data)
	if err != nil {
		log.WithError(err).WithField("data", data).Error("failed to marshal data")
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	w.Write(resp)
	return
}

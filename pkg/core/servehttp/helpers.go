package servehttp

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func WaitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Shutting down")
	os.Exit(0)
}

func ResponseJSON(w http.ResponseWriter, httpCode int, data interface{}) {
	jData, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.WriteHeader(httpCode)
	w.Write(jData)
}

func ResponseSuccessJSON(w http.ResponseWriter, data interface{}) {
	ResponseJSON(w, http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   data,
	})
}

func ResponseErrorJSON(w http.ResponseWriter, httpCode int, message string) {
	ResponseJSON(w, httpCode, map[string]interface{}{
		"status":  "error",
		"message": message,
	})
}

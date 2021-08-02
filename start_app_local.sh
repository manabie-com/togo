#/bin/bash

export GO_ENV=local PORT=8080
LOG_FILE="log-"$GO_ENV"-"`date '+%Y%m%d'`".txt" 

go run main.go  #2>&1 | tee $LOG_FILE
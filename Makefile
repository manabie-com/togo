include internal/.env
run:
	go build -o internal/build/togo togo/internal/ 
	sh run.sh
	
testing.unit:
	@go test -v togo/internal/dao/

kill:
	@lsof -n -i4TCP:${PORT} | grep LISTEN | awk '{ print $2 }'| sh kill.sh | echo "SUCCESSFULLY KILLED PROCESS ON PORT ${PORT}"
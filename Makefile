run:
	go build -o internal/build/togo togo/internal/
	nohup ./internal/build/togo >/dev/null 2>&1 &
	

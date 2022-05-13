BINARY_NAME=main

build:

run:
	go test main_test.go -v
	nodemon --exec go run main.go

all: clean run

clean:
	go clean
	go mod tidy
all: run test

run:
	go run main.go

test:
	go test --cover ./...

test:
	go test ./...
test-coverage:
	go test -cover
run:
	go run main.go
clean:
	go mod tidy
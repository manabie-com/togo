test:
	go test ./...
test-coverage:
	go test -cover
test-integration:
	go test *.go
run:
	go run main.go
clean:
	go mod tidy
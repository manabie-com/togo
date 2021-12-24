.PHONY: all test clean

gen-proto:
	buf build
	buf generate

#go install github.com/matryer/moq@latest
gen-mock:
	moq -stub -pkg mock_storage -out ./mock/mock_storage/mock_storage.go ./internal/storage  Storage 

build:
	go build
	#CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-extldflags "-static"'

run:
	PORT=9090 METRIC_PORT=8080 CONFIG_FILE=./config/dev.toml ./togo-internal-service serve

migrate:
	go build
	#CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-extldflags "-static"'
	./togo-internal-service migrate

TEST_PACKAGE = $(shell go list ./... | grep -v /test)
#COVER_PACKAGE = $(shell go list ./... | grep -v '/test\|/cmd')
test:
	go test -v -race ./...

test-cover:
	go test -v -race -coverpkg=./... $(TEST_PACKAGE) -coverprofile=coverage.out -covermode=atomic
	go tool cover -html coverage.out

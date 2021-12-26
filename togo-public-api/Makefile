.PHONY: all test clean

#go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
#go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
gen-proto:
	buf build
	buf generate

#go install github.com/matryer/moq@latest
gen-mock:
	moq -stub -pkg mock_togo_internal_v1 -out ./mock/mock_togo_internal_v1/mock_togo_internal_v1.go ./internal/service/togo_internal_v1  TogoInternalServiceClient 
	moq -stub -pkg mock_togo_user_session_v1 -out ./mock/mock_togo_user_session_v1/mock_togo_user_session_v1.go ./internal/service/togo_user_session_v1  TogoUserSessionServiceClient 

build:
	go build
	#CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-extldflags "-static"'

run:
	PORT=9090 METRIC_PORT=8080 CONFIG_FILE=./config/dev.toml ./togo-public-api serve

TEST_PACKAGE = $(shell go list ./... | grep -v /test)
#COVER_PACKAGE = $(shell go list ./... | grep -v '/test\|/cmd')
test:
	go test -v -race ./...

test-cover:
	go test -v -race -coverpkg=./... $(TEST_PACKAGE) -coverprofile=coverage.out -covermode=atomic
	go tool cover -html coverage.out

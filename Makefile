.PHONY: local-db build proto gen-mock run gci lint lint-consistent sec

local-db:
	@docker-compose down
	@docker-compose up -d

build:
	go build -ldflags "-s -w" -o ./tmp/server ./cmd/main.go

proto:
	@protoc --go_out=plugins=grpc:./proto --proto_path=proto --go_opt=paths=source_relative proto/togo_service.proto

gen-mock:
	# repository
	@mockery --inpackage --name=Repository --dir=./repository/user

run:
	@GO111MODULE=off go get -u github.com/cosmtrek/air
	@air -c .air.conf

unit-test:
	@mkdir coverage || true
	@go test -race -v -coverprofile=coverage/coverage.txt.tmp -count=1  ./...
	@cat coverage/coverage.txt.tmp | grep -v "mock_" > coverage/coverage.txt
	@go tool cover -func=coverage/coverage.txt
	@go tool cover -html=coverage/coverage.txt -o coverage/index.html

gci:
	@GO111MODULE=off go get github.com/daixiang0/gci
	gci -w -local github.com/khangjig/togo .

lint:
	@hash golangci-lint 2>/dev/null || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.43.0
	@GO111MODULE=on CGO_ENABLED=0 golangci-lint run

lint-consistent:
	@hash go-consistent 2>/dev/null || GO111MODULE=off go get -v github.com/quasilyte/go-consistent
	@go-consistent ./...

sec:
	@curl -sfL https://raw.githubusercontent.com/securego/gosec/master/install.sh | sh -s -- -b $(GOPATH)/bin latest
	@gosec ./...
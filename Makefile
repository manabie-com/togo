TARGET = bin
TARGET_BIN = go-runtime
GO_CMD_MAIN = cmd/server/main.go

build:
	go build -o $(TARGET)/$(TARGET_BIN) $(GO_CMD_MAIN)

run:
	go run $(GO_CMD_MAIN) server

migrate-up:
	go run $(GO_CMD_MAIN) migrate up

test:
	go test ./...  -count=1 -v -cover -race

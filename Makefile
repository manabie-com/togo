
all: proto format analyzer build-server

minipro-proto:
	cd rpc_services && buf mod update && buf generate --path rpc_services.proto

auth-proto:
	cd auth_services && buf mod update && buf generate --path proto/auth.proto

proto: minipro-proto auth-proto


format:
	gofmt -l -s -w .
analyzer:
	go vet ./...
build-server:
	CGO_ENABLED=1 go build -ldflags=-w -o mini_project main.go  #for release	

clean:
	go clean
	rm -rf mini_project
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)
GOIMP = "goimports"
GOFMT = "gofumpt"
MODULE_PATH = "github.com/vchitai/togo"
GOTOOLS = golang.org/x/lint/golint \
	github.com/golangci/golangci-lint/cmd/golangci-lint \
	golang.org/x/tools/cmd/goimports \
	mvdan.cc/gofumpt \
	github.com/bufbuild/buf/cmd/buf\
	github.com/gogo/protobuf/protoc-gen-gogo \
	github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
	github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
	github.com/mwitkow/go-proto-validators/protoc-gen-govalidators \


init:
	go mod init $(MODULE_PATH)

tidy:
	GOSUMDB=off go mod tidy

vendor:
	GOSUMDB=off go mod vendor

install-go-tools:
	go get $(GOTOOLS)

fmt: ## Run gofmt for all .go files
	@$(GOIMP) -w $(GOFMT_FILES)
	@#$(GOFMT) -w $(GOFMT_FILES)

generate-v1: ## Generate proto
	protoc \
		-I proto/ \
		-I .third_party/googleapis \
		-I .third_party/envoyproxy \
		-I .third_party/gogoprotobuf \
		-I .third_party/protoc-gen-swagger \
		-I .third_party/protoc-gen-govalidators \
		--descriptor_set_out=./descriptors.protoset\
		--include_source_info --include_imports -I. \
		--gogo_out=plugins=grpc,paths=source_relative,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,\
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types:\
pb \
		--grpc-gateway_out=paths=source_relative,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,\
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types:\
pb \
		--swagger_out=docs \
		--validate_out=lang=go,gogoimport=true,paths=source_relative,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,\
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types:\
pb \
		--auth_out=paths=source_relative,.:pb\
		proto/*.proto

run:
	@go run cmd/server/*.go

generate:
	buf generate
	@(cd ./pb ; GOSUMDB=off go mod tidy)

test: ## Run go test for whole project
	@go test -v ./...

lint: ## Run linter
	@golangci-lint run ./...

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: init install-go-tools update fmt generate test lint help
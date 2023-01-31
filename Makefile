VERSION=0.1

BUILD := `git rev-parse HEAD`

# Operating System Default (LINUX)
TARGETOS=linux

GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)
GOFMT := "goimports"
LDFLAGS=-ldflags "-s -w -X=main.Version=$(VERSION) -X=main.Build=$(BUILD) -extldflags -static"

install:
	go install github.com/google/wire/cmd/wire
	go install entgo.io/ent/cmd/ent@latest
	go install github.com/golang/mock/mockgen@v1.6.0
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1

test:
	./scripts/test.sh

lint: ## Run linter
	@golangci-lint run ./...

fmt: ## Run gofmt for all .go files
	@$(GOFMT) -w $(GOFMT_FILES)

build_local:
	go build $(LDFLAGS) -o bin/togo ./cmd/togo

build:
	@GOOS=$(TARGETOS) CGO_ENABLED=0 go build $(LDFLAGS) -o bin/togo ./cmd/togo

gen: gen-schema
	## Go generate
	@go generate ./...

gen-constructors:
	cd internal && wire

gen-schema:
	@ent generate ./database/schema --target database/ent

dev-up:
	docker-compose up -d

dev-down:
	docker-compose down

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
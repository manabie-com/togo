VERSION=0.1

BUILD := `git rev-parse HEAD`

# Operating System Default (LINUX)
TARGETOS=linux

LDFLAGS=-ldflags "-s -w -X=main.Version=$(VERSION) -X=main.Build=$(BUILD) -extldflags -static"

check:
	go mod tidy
	test -z "$(git status --porcelain)"
	test -z $(shell gofmt -l main.go | tee /dev/stderr) || echo "[WARN] Fix formatting issues with 'make fmt'"
	golint ./...
	go vet ./...

build_local:
	go build $(LDFLAGS) -o bin/togo ./cmd/togo

build:
	@GOOS=$(TARGETOS) CGO_ENABLED=0 go build $(LDFLAGS) -o bin/togo ./cmd/togo

install:
	go get github.com/google/wire/cmd/wire && \
	go get entgo.io/ent/cmd/ent

gen: gen-schema
	## Go generate
	@go generate ./...

gen-constructors:
	cd internal && wire

gen-schema:
	go run entgo.io/ent/cmd/ent generate ./database/schema --target database/ent

dev-up:
	docker-compose up -d

dev-down:
	docker-compose down

test:
	./scripts/test.sh

coverage: test
	@go tool cover -html=test-results/.testcoverage.txt -o test-results/coverage.html && open test-results/coverage.html

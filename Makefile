SHELL=/bin/bash -o pipefail

export GO111MODULE        := on
export PATH               := .bin:${PATH}
export PWD                := $(shell pwd)

export BUILD_DATE         := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
export VCS_REF            := $(shell git rev-parse HEAD)
export POSTGRESQL_URL 	  :='postgres://togo:secret@localhost:5432/togo?sslmode=disable'

.PHONY: install-deps
install-deps:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest && \
	go install github.com/golang/mock/mockgen@latest

.PHONY: migrate-up
migrate-up:
	migrate -database ${POSTGRESQL_URL} -path internal/persistence/migrations/sql up

.PHONY: migrate-down
migrate-down:
	migrate -database ${POSTGRESQL_URL} -path internal/persistence/migrations/sql down

.PHONY: docker-compose
docker-compose:
	DOCKER_BUILDKIT=1 docker build -f .docker/togo/Dockerfile --build-arg=COMMIT=$(VCS_REF) --build-arg=BUILD_DATE=$(BUILD_DATE) -t trangmaiq/togo .
	docker-compose -f .docker/quickstart_postgres.yml up -d

.PHONY: generate
generate:
	go generate ./...

.PHONY: unit-test
unit-test:
	go test -timeout 5m -coverprofile cover.out ./...
	go tool cover -html=cover.out -o cover.html

.PHONY: integration-test
integration-test:
	go test -tags=integration


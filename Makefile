#!make
include .conf
include .env
export

PKG := "github.com/manabie-com/togo"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

# Self documented Makefile
# http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

help: ## Show list of make targets and their description
	@grep -E '^[/%.a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL:= help

run: 
	go run main.go

.PHONY: test
test: ## Generate mock and run all test. To run specified tests, use `./scripts/test.sh <pattern>`)
	@scripts/test.sh $*

.PHONY: lint
lint: ## Run linter
	@scripts/lint.sh

.PHONY: build
build/linux: ## Build server for linux
	env GOOS=linux GOARCH=386 go build -ldflags="-s -w" -o ./bin/server-linux ./cmd/server/

build/mac: ## Show build.sh help for building binnary package under cmd
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/server/

gen: ## Generate models using sqlboiler following pre-defined templates
	sqlboiler psql --wipe --add-soft-deletes 

db/migrate: ## Migrate database structure
	@scripts/migrate.sh up

db/up: ## Apply all the migration to the latest version to the local database
	@make db/migrate

db/down: ## Remove every in the database! (only for DEV)
	@scripts/migrate.sh down 

db/drop: ## Remove every in the database! (only for DEV)
	@scripts/migrate.sh drop -f

db/reset: ## Remove everything and recreate the database! (only for DEV)
	@echo y | make db/drop
	@make db/up

db/connect: ## Connect to the database
	pgcli ${DB_URL}

db/%: ## Run other migrate commands
	@scripts/migrate.sh $*

docker/up: ## Run docker compose
	docker-compose -f ./docker/docker-compose.yaml up -d

docker/down: ## Stop docker compose
	docker-compose -f ./docker/docker-compose.yaml down

setup: ## Perform setup script, install necessary plugins/tools
	@scripts/setup.sh

race: dep ## Run data race detector
	@go test -race -short ${PKG_LIST}

msan: dep ## Run memory sanitizer
	@go test -msan -short ${PKG_LIST}

coverage: ## Generate global code coverage report
	@scripts/coverage.sh;

coverhtml: ## Generate global code coverage report in HTML
	@scripts/coverage.sh html;

dep: ## Get the dependencies
	@go get -v -d ./...

clean: ## Remove previous build
	@rm -f $(PROJECT_NAME)

include .env
export

.PHONY: help
help: 
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL:= help

.PHONY: docker
docker/build: ## Build docker compose images first
	docker-compose -f ./docker/docker-compose.yaml build 

docker/up: ## Start docker compose
	docker-compose -f ./docker/docker-compose.yaml up -d 

docker/down:  ## Stop and remove docker compose
	docker-compose -f ./docker/docker-compose.yaml down

docker/clean: ## Clean all docker data
	@make docker/down
	rm -rf ./data

.PHONY: build
build/linux: ## Build server for linux
	env GOOS=linux GOARCH=386 go build -ldflags="-s -w" -o ./bin/server ./cmd/server/

build/prod:
	CGO_ENABLED=0 GOOS=linux go build -o ./bin/server ./cmd/server

build: ## Show build.sh help for building binnary package under cmd
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./bin/server ./cmd/server/

.PHONY: run
run:
	go run cmd/server/main.go

.PHONY: setup
setup: ## Run setup scripts to prepare development environment
	@scripts/setup.sh

.PHONY: clean 
clean: ## Clean project dir, data and docker 
	@scripts/clean.sh

.PHONY: lint
lint: ## Run linter
	@scripts/lint.sh


.PHONY: db
db/migrate: ## Migrate database structure
	@scripts/migrate.sh up

db/up: ## Apply all the migration to the latest version to the local database
	@make db/migrate

db/down: ## Remove everything the database! (only for DEV)
	@scripts/migrate.sh down 

db/drop: ## Remove everything the database! (only for DEV)
	@scripts/migrate.sh drop -f 

db/gen:
	sqlboiler psql --wipe --config ./db/sqlboiler.toml --add-soft-deletes 

db/reset:
	@make docker/down
	docker volume rm -f docker_togo-volume
	@make docker/up

# Basic commands: up/down/drop/force\ <version>
db/%:
	@scripts/migrate.sh $*

.PHONY: test
test: ## Generate mock and run all test. To run specified tests, use `./scripts/test.sh <pattern>`)
	@scripts/test.sh $*

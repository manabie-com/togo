.PHONY: prepare docker_conf docker_up docker_stop docker_down clean di gen_migration migup migup1 migdown1 migreset migstat migver migfix execdb sqlc test mock build

-include app.env
export $(shell test -f app.env && sed 's/=.*//' app.env)

prepare:
	@echo "Installing golangci-lint"
	@curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest
	@echo "Installing goose"
	@go install github.com/pressly/goose/v3/cmd/goose@latest
	@echo "Installing wire"
	@go install github.com/google/wire/cmd/wire@latest
	@echo "Installing mockgen"
	@go install github.com/golang/mock/mockgen@v1.6.0
	@echo "Copy app.env.example to app.env if app.env does not exists"
	@test -f app.env || cp -a app.env.example app.env
	@echo "Install Husky"
	cd ../ && go install github.com/go-courier/husky/cmd/husky@latest && husky init

docker_conf:
	@docker-compose -p $(APP_NAME) -f docker-compose.dev.yaml --env-file ./app.env config

docker_up:
	@docker-compose -p $(APP_NAME) -f docker-compose.dev.yaml --env-file ./app.env up -d

docker_stop:
	@docker-compose -p $(APP_NAME) -f docker-compose.dev.yaml --env-file ./app.env stop

docker_down:
	@docker-compose -p $(APP_NAME) -f docker-compose.dev.yaml --env-file ./app.env down -v

clean:
	[ -d dist ] && rm dist/*

di:
	@cd wire && wire

# Usage: make seq=${seq} migration
gen_migration:
	$(info Generating migration sequence $(seq) to migration)
	@goose -dir migration $(DB_DRIVER) "$(DB_DRIVER)://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" create $(seq) sql

migup:
	@goose -allow-missing -dir migration $(DB_DRIVER) "$(DB_DRIVER)://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" up

migup1:
	@goose -allow-missing -dir migration $(DB_DRIVER) "$(DB_DRIVER)://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" up-by-one

migdown1:
	@goose -dir migration $(DB_DRIVER) "$(DB_DRIVER)://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" down

migreset:
	@goose -dir migration $(DB_DRIVER) "$(DB_DRIVER)://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" reset

migstat:
	@goose -dir migration $(DB_DRIVER) "$(DB_DRIVER)://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" status

migver:
	@goose -dir migration $(DB_DRIVER) "$(DB_DRIVER)://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" version

migfix:
	@goose -dir migration $(DB_DRIVER) "$(DB_DRIVER)://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" fix

execdb:
	@docker-compose -p $(APP_NAME) -f docker-compose.dev.yaml --env-file ./app.env exec postgres psql -U postgres $(DB_NAME)

# Usage: make domain=${domain} mock_repo
mock_repo:
	$(info Generating mock repository for $(domain))
	@mockgen -package mock -source=./domain/$(domain)/$(domain)_repository.go -destination=./$(domain)/mock/mock_$(domain)_repository.go

# Usage: make domain=${domain} mock_service
mock_service:
	$(info Generating mock service for $(domain))
	@mockgen -package mock -source=./$(domain)/service/$(domain)_service.go -destination=./$(domain)/mock/mock_$(domain)_service.go

build: di
	@go build -o dist/app app.go

dev: docker_up di
	@export APP_ENV=local && go run app.go

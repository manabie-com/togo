BIN=$(CURDIR)/scripts
SERVICE_NAME=todo
PJT_NAME=$(basename `pwd` | sed s/-//g)
VERSION=$(shell cat ./VERSION)

.PHONY: tidy run stop create_migration migrate seed create_mock

tidy:
	$(BIN)/compose exec $(SERVICE_NAME) go mod tidy

test: 
	$(BIN)/compose exec $(SERVICE_NAME) go test ./...

run:
	$(BIN)/compose up -d

stop:
	$(BIN)/compose stop

build_prod_img:
	docker build -f build/Dockerfile.prod -t manabie-com/togo:$(VERSION) .

create_migration:
	$(BIN)/create_migration

migrate:
	$(BIN)/migrate up all

migrate_test:
	$(BIN)/migrate up all -t

seed:
	$(BIN)/compose exec $(SERVICE_NAME) go run cmd/todoseed/todoseed.go

create_mock:
	@docker run -v "$(CURDIR)":/src \
		-w /src vektra/mockery --all  --output "./internal/todo/mocks" --dir "./internal/todo/"
MIGRATIONS_DIR=infrastructure/datastore/migrations/postgres
POSTGRESQL_URL=postgres://postgres:postgres@localhost:5432/todo?sslmode=disable

dev:
	go mod tidy
	go run cmd/todo/main.go

create-migration:
	migrate create -ext=sql -dir=$(MIGRATIONS_DIR) $(FILE_NAME)

run-migration:
	migrate -database $(POSTGRESQL_URL) -path $(MIGRATIONS_DIR) up

undo-migration:
	migrate -database $(POSTGRESQL_URL) -path $(MIGRATIONS_DIR) down

# mockery --dir=./usecase/interfaces --name=UserPresenter --filename=user_formatter_mock.go --structname=UserFormatterMock

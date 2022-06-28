#!make
### Export file .env
include .env
export

### Migration
migration-up:
	@echo migrate -database ${POSTGRESQL_URL_MIGRATION} -path db/migrations up
	@migrate -database ${POSTGRESQL_URL_MIGRATION} -path db/migrations up

migration-down:
	@echo migrate -database ${POSTGRESQL_URL_MIGRATION} -path db/migrations down
	@migrate -database ${POSTGRESQL_URL_MIGRATION} -path db/migrations down

### Run project
run:
	go run cmd/main.go

### Test
test-api-common:
	go test -cover -coverprofile coverage.log ./internal/api/handlers/common/...

test-api-task:
	go test -cover -coverprofile coverage.log ./internal/api/handlers/tasks/...

test-api-user:
	go test -cover -coverprofile coverage.log ./internal/api/handlers/users/...

test-all:
	go test  -cover -coverprofile coverage.log ./internal/api/handlers/...

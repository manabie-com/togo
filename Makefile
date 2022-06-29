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

### Unit test
test-unit:
	go test  -cover -coverprofile coverage.log ./internal/...

### Integraion test
test-integration:
	go test ./integrationtest/...

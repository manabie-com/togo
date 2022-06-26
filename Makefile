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

test-migration:
	@echo migrate -database ${POSTGRESQL_URL_MIGRATION} -path db/migrations down
	@echo migrate -database ${POSTGRESQL_URL_MIGRATION} -path db/migrations up

nodemon:
	nodemon --exec go run cmd/main.go --signal SIGTERM

include .env

POSTGRES_URL=postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable
IMAGE_MIGRATE=migrate/migrate:4
IMAGE_SQLC=kjconroy/sqlc:1.13.0
PROJECT_PATH := ${CURDIR}
MIGRATION_PATH=${PROJECT_PATH}/internal/pkg/sql

infra:
	docker-compose --env-file .env up -d --build
down:
	docker-compose down
server:
	go run cmd/togo/main.go
# ex: make NAME=example migrate_init
migrate_init:
	docker run --rm --network host -v ${MIGRATION_PATH}:/migration/ ${IMAGE_MIGRATE} \
	create -ext sql -dir /migration/schema ${NAME}
db_up:
	./scripts/wait-for.sh localhost:5432 -t 60
	sleep 2
	docker run --rm --network host -v ${MIGRATION_PATH}/:/repository ${IMAGE_MIGRATE} \
	-path=/repository/schema -database ${POSTGRES_URL} up
db_down:
	docker run --rm --network host -v ${MIGRATION_PATH}/:/repository -w /repository ${IMAGE_MIGRATE} \
 	-path=/repository/schema -database ${POSTGRES_URL} down --all
check_db:
	docker exec ${POSTGRES_CONTAINER} psql -U ${POSTGRES_USER} -d ${POSTGRES_DB} -c '\dt'
sqlc:
	docker run --rm -v ${MIGRATION_PATH}:/repository -w /src ${IMAGE_SQLC} generate -f /repository/sqlc.yaml
unit_test:
	go test -v -cover -short ./...
integration_test:
	go test -v -run Integration ./...
test:
	go test -v -cover ./...
mock:
	mockgen -package mockdb -destination ./internal/pkg/sql/sqlc/mock/querier.go -source internal/pkg/sql/sqlc/query.go

up: infra db_up

.PHONY: infra down server migrate_init db_up db_down check_db sqlc unit_test integration_test test mock
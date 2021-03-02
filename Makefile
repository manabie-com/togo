PG_USERNAME=postgres
PG_PASSWORD=example
PG_HOST=localhost
PG_DATABASE=postgres
PG_SQL_INIT=./docker/postgres/data.sql

postgres-start:
	docker-compose -f ./docker/postgres/docker-compose.yml up

postgres-init:
	PGPASSWORD=$(PG_PASSWORD) psql $(PG_USERNAME) -h $(PG_HOST) -d $(PG_DATABASE) -f $(PG_SQL_INIT) -a

start:
	go run main.go

test-unit:
	go test ./internal/... -count=1 -coverprofile=coverage.out
	go tool cover -html=coverage.out

test-integrate:
	go test ./test/... -count=1 --tags=integrate
PG_USERNAME=postgres
PG_PASSWORD=example
PG_HOST=localhost
PG_DATABASE=postgres
PG_SQL_INIT=./docker/postgres/data.sql

postgres-start:
	docker-compose -f ./docker/postgres/docker-compose.yml up

postgres-init:
	PGPASSWORD=$(PG_PASSWORD) psql $(PG_USERNAME) -h $(PG_HOST) -d $(PG_DATABASE) -f $(PG_SQL_INIT) -a

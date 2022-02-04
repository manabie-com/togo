postgres:
	docker run --name todopostgres -p 5432:5432 -e POSTGRES_USER=postgres -e  POSTGRES_PASSWORD=postgres -d postgres

createdb:
	docker exec -it postgreslatest createdb --username=postgres --owner=postgres todo

migrateup:
	migrate -database ${POSTGRESQL_URL} -path db/migrations up

migratedown:
	migrate -database ${POSTGRESQL_URL} -path db/migrations down

sqlc:
	sqlc generate

PHONY: postgres createdb migrateup migratedown sqlc
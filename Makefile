db_start:
	docker run --name todopostgres -p 5432:5432 -e POSTGRES_USER=postgres -e  POSTGRES_PASSWORD=postgres -d postgres

db_stop:
  docker stop todopostgres

createdb:
	docker exec -it todopostgres createdb --username=postgres --owner=postgres todo_app

migrateup:
	migrate -database ${POSTGRESQL_URL} -path db/migrations up

migratedown:
	migrate -database ${POSTGRESQL_URL} -path db/migrations down

sqlc:
	sqlc generate

test_app:
	go test ./api/

run:
	go run main.go

PHONY: postgres createdb migrateup migratedown sqlc test_app
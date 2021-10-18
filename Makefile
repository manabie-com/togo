postgres:
	docker run --name postgres-todo -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d fintrace/postgres-uuid:latest

createdb:
	docker exec -ti postgres-todo createdb --username=root --owner=root todo

dropdb:
	docker exec -ti postgres-todo dropdb todo

migrateup:
	migrate -path ./db/migration -database "postgresql://root:secret@localhost:5432/todo?sslmode=disable" -verbose up

migratedown:
	migrate -path ./db/migration -database "postgresql://root:secret@localhost:5432/todo?sslmode=disable" -verbose down

mock:
	mockgen -package mockdb -destination internal/storages/mock/store.go github.com/jericogantuangco/togo/internal/storages/postgres Store

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown mock test server
postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root togo

dropdb:
	docker exec -it postgres12 dropdb togo

migrateup:
	migrate -path internal/storages/migration -database "postgresql://root@localhost:5432/togo?sslmode=disable" -verbose up

migratedown:
	migrate -path internal/storages/migration -database "postgresql://root@localhost:5432/togo?sslmode=disable" -verbose down

test:
	go test -v -cover -count=1 ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown test server 

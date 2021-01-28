postgresql:
	docker run --name postgresql -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -d bitnami/postgresql

createdb:
	docker exec -it postgresql createdb --username=postgres --owner=postgres togo

createdb-test:
	docker exec -it postgresql createdb --username=postgres --owner=postgres togo_test

dropdb:
	docker exec -it postgresql dropdb togo

migrateup:
	migrate -path internal/storages/migration -database "postgresql://postgres:password@localhost:5432/togo?sslmode=disable" -verbose up

migrateup-test:
	migrate -path internal/storages/migration -database "postgresql://postgres:password@localhost:5432/togo_test?sslmode=disable" -verbose up

migratedown:
	migrate -path internal/storages/migration -database "postgresql://postgres:password@localhost:5432/togo?sslmode=disable" -verbose down

migratedown-test:
	migrate -path internal/storages/migration -database "postgresql://postgres:password@localhost:5432/togo_test?sslmode=disable" -verbose down

test:
	go test -v -cover -count=1 ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown test server
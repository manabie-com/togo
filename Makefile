makeFileDir := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

postgres:
	- docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:alpine
	- docker start postgres

createmigrate:
	migrate create -ext sql -dir db/migration -seq init_schema

createdb:
	docker exec -it postgres createdb --username=root --owner=root togo

dropdb:
	docker exec -it postgres dropdb togo

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/togo?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/togo?sslmode=disable" -verbose down

sqlc:
	docker run --rm -v $(makeFileDir):/src -w /src kjconroy/sqlc generate

test:
	go test -v -count=1 -race -cover -coverprofile=profile.cov ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go togo/db/sqlc Store

clean:
	docker rm -f $(docker ps -a -q)

cover:
	go tool cover -func profile.cov

badger:
	gopherbadger -md="README.md,coverage.md"

.PHONY: postgres createmigrate createdb dropdb migrateup migratedown sqlc test server mock clean cover badger

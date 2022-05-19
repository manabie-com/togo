DB_URL=postgresql://root:root@localhost:5432/task_management?sslmode=disable

.PHONY: network postgres create-db drop-db migrate-up migrate-down sqlc test

network:
	 docker network create task-network

postgres:
	docker run --name postgres --network task-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:14-alpine

create-db:
	docker exec -it postgres createdb --username=root --owner=root task_management

drop-db:
	docker exec -it postgres dropdb task_management

migrate-up:
	migrate -path internal/db/migration -database "$(DB_URL)" -verbose up

migrate-down:
	migrate -path internal/db/migration -database "$(DB_URL)" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...
server:
	go run internal/main.go

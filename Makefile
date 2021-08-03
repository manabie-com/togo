.PHONY: migrate, rollback, migration, run, sqlc, test, docker-up, docker-build

migrate:
	go run cmd/migrate/main.go up

rollback:
	go run cmd/migrate/main.go down

migration:
	@read -p "Enter migration name: " name; \
		go run cmd/migrate/main.go create $$name sql

run:
	go run cmd/app/main.go

sqlc:
	sqlc generate

test:
	go test ./...

docker-up:
	docker-compose up

docker-build:
	docker-compose build

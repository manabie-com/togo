.PHONY: migrate, migrate-down, run

migrate:
	go run cmd/migrate/main.go up

migrate-down:
	go run cmd/migrate/main.go down

run:
	go run cmd/app/main.go

check_env:
	docker-compose config

setup:
	docker-compose up -d

migrate:
	go run cmd/migrate/main.go

seed:
	go run cmd/seed/main.go

start:
	go run cmd/server/main.go

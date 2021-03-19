check_env:
	docker-compose config

setup:
	docker-compose up -d

migrate:
	go run cmd/db/main.go

start:
	go run cmd/server/main.go

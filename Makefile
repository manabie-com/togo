check_env:
	docker-compose config

setup:
	docker-compose up -d

start:
	go run cmd/server/main.go

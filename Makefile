CMD_MAIN=cmd/main.go

run:
	go run ${CMD_MAIN} server

migrate:
	echo \# make migrate name="${name}"
	go run $(CMD_MAIN) migrate create $(name)

migrate-up:
	go run $(CMD_MAIN) migrate up

migrate-down:
	go run $(CMD_MAIN) migrate down 1

test:
	go test ./...  -count=1 -v -cover

start-db:
	docker-compose up -d
CMD_MAIN=cmd/main.go

run:
	go run ${CMD_MAIN} server

migrate:
	echo \# make migrate name="${name}"
	go run $(CMD_MAIN) migrate create $(name)

migrate-up:
	go run $(CMD_MAIN) migrate up
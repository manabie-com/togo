SHELL := /bin/bash

init-local:
	./env.sh
	docker-compose up -d postgres

run:
	go run main.go

test:
	go test ./... -v
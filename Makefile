SHELL 	 := /bin/bash

dev.up:
	@command cd deploy/dev_todo && docker-compose up

run:
	@command go run main.go ./deploy/dev_todo/


test.unit:
	@command go test ./... -v

test.integration:
	@command  go test -tags=integration ./integration -v -count=1

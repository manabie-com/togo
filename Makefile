SHELL 	 := /bin/bash

run:
	@command go run main.go ./deploy/dev_todo/

dev_up:
	@command cd deploy/dev_todo && docker-compose up

test_unit:
	@command go test github.com/manabie-com/togo/internal/services/test

test_integration:
	@command go test github.com/manabie-com/togo/integration

test_up:
	@command cd deploy/test_todo && docker-compose up

test_down:
	@command cd deploy/test_todo && docker-compose down
SHELL 	 := /bin/bash

run:
	@command go run main.go ./deploy/dev_todo/

devup:
	@command cd deploy/dev_todo && docker-compose up

test:
	@command go run main.go ./deploy/test_todo/

testup:
	@command cd deploy/test_todo && docker-compose up

testdown:
	@command cd deploy/test_todo && docker-compose down
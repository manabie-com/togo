.PHONY: install
install:
	go mod tidy

.PHONY: migrate
migrate:
	go run migrations/main.go

.PHONY: docker-compose
docker-compose:
	docker-compose up

.PHONY: run
run:
	go run main.go

.PHONY: unit-test
unit-test:


.PHONY: integration-test
integration-test:
	
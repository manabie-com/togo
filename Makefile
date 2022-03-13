#!make
BINARY=engine
ENV_LOCAL_TEST=\
	PORT=9090 \
	CACHE_HOST=localhost \
	CACHE_PORT=6379 \
	DB_SCHEME=mysql \
	DB_HOST=localhost \
	DB_PORT=3306 \
	DB_USER=triet_truong \
	DB_PASS=pw \
	DB_NAME=todo 

app.run:
	go run ./app/main.go

app.test:
	go test -v -race -covermode=atomic -cover ./...

app.unit_test:
	go test -v -race -covermode=atomic -cover ./todo/usecase/todo_usecase_test.go

app.integration_test:
	$(ENV_LOCAL_TEST) \
	go test -v -race -covermode=atomic -cover ./integration_test/e2e_test.go

docker.start:
	docker-compose up -d

docker.stop:
	docker-compose down

docker.restart: docker.stop docker.start

lint-prepare:
	@echo "Installing golangci-lint" 
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

lint:
	./bin/golangci-lint run ./...

.PHONY: clean install unittest build docker run stop vendor lint-prepare lint
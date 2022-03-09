BINARY=engine
export ENVIRONMENT=LOCAL

app.run:
	go run ./app/main.go

app.unittest:
	go test -race -covermode=atomic -cover ./...

docker.start:
		docker-compose up -d;

docker.stop:
		docker-compose down;

docker.restart: docker.stop docker.start

lint-prepare:
	@echo "Installing golangci-lint" 
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

lint:
	./bin/golangci-lint run ./...

.PHONY: clean install unittest build docker run stop vendor lint-prepare lint
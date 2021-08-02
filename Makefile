load-env:
	set -o allexport; source .env; set +o allexport
build:
	go build -o app-exe ./cmd/srv/...
start-exe:
	set -o allexport; source .env; set +o allexport && ./app-exe

start:
	set -o allexport; source .env; set +o allexport && go run ./cmd/srv/main.go
migrate:
	set -o allexport; source .env; set +o allexport && go run ./cmd/migrate/main.go

unit-test: 
	./test.sh

docker-dev:
	docker-compose -f docker-compose.dev.yml up -d
docker-start:
	docker-compose down
	docker-compose up -d --build

# .PHONY: test-all

# test-all: 
# 	go test $(go list ./... | grep -v /test)

# prerequisites: test-all

# target: prerequisites 

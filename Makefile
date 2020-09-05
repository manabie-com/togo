
start:
	docker-compose up -d

stop:
	docker-compose stop

build:
	docker-compose build

run-with-swagger:
	go build && swag init && go run main.go

build-dev:
	go build -v main.go

run-dev:
	go run main.go
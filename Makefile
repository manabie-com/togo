run-dev:
	go run cmd/app/main.go

run-test:
	go test -cover fmt ./...

docker-run:
	docker-compose down
	docker-compose up -d --build